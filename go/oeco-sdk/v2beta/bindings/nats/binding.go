package natsnodev1

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	apexlog "github.com/apex/log"
	natsd "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	nebulav1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/nebula"
	zaploggerv1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/zap"
)

const (
	DEFAULT_CONNECTOR_ACCOUNT = "CONNECTOR"
)

// Binding represents a structure managing NATS connections, JetStream instances, and event stream configurations.
type Binding struct {
	Registry                map[string]nats.StreamConfig
	SpecEventListeners      []SpecEventListener
	SpecEventBatchListeners []SpecEventBatchListener
	Listeners               map[string]*nats.Subscription
	Nats                    *nats.Conn
	JetStreamContext        *nats.JetStreamContext
	JetStream               *jetstream.JetStream

	server        *natsd.Server
	configuration *Configuration
}

// natsOptions defines configuration options for NATS connection.
// jsOptions defines options specific to JetStream interactions.
// Bound refers to the binding instance used for NATS and JetStream integrations.
// BindingName is the constant name for the NATS node binding.
var (
	natsOptions []nats.Option
	// jsOptions   []jetstream.JetStreamOpt

	Bound       *Binding
	BindingName = "NATS_NODE_BINDING"
)

// Name returns the name of the binding, which is a constant string "NATS_NODE_BINDING".
func (b *Binding) Name() string {
	return BindingName
}

// Validate checks the validity of the provided Bindings. Returns an error if validation fails, otherwise nil.
func (b *Binding) Validate(_ context.Context, _ *sdkv2betalib.Bindings) error {
	return nil
}

func NewNatsBinding(configuration *Configuration) *Binding {
	return &Binding{
		configuration: configuration,
	}
}

// Bind initializes the Binding instance, configures NATS or JetStream connections, and registers the binding in Bindings.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2betalib.Bindings) *sdkv2betalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				switch b.configuration.Natsd.Enabled {
				case true:

					if b.configuration.Natsd.Username == "" && b.configuration.Natsd.Password == "" {
						panic("Please provide Natsd username and password")
					}

					if b.configuration.Nats.Username == "" && b.configuration.Nats.Password == "" {
						panic("Please provide Nats username and password; This allows for connecting to the NATS server from the NATS client")
					}

					// r := rand.New(rand.NewSource(time.Now().UnixNano()))

					systemAccount := natsd.NewAccount(natsd.DEFAULT_SYSTEM_ACCOUNT)
					connectorAccount := natsd.NewAccount(DEFAULT_CONNECTOR_ACCOUNT)
					storeDir := NatsdServerJetstreamStoreDir
					if b.configuration.Natsd.StoreDir != "" {
						storeDir = b.configuration.Natsd.StoreDir
					}

					options := natsd.Options{
						// Standard client options
						ServerName:    b.configuration.Natsd.ServerName,
						Host:          b.configuration.Natsd.Host,
						Port:          b.configuration.Natsd.Port,
						HTTPPort:      b.configuration.Natsd.HTTPPort,
						SystemAccount: natsd.DEFAULT_SYSTEM_ACCOUNT,
						Accounts: []*natsd.Account{
							systemAccount,
							connectorAccount,
						},
						Users: []*natsd.User{
							{
								Account: &natsd.Account{
									Name: natsd.DEFAULT_SYSTEM_ACCOUNT,
								},
								Username: b.configuration.Natsd.Username,
								Password: b.configuration.Natsd.Password,
							},
							{
								Account: &natsd.Account{
									Name: DEFAULT_CONNECTOR_ACCOUNT,
								},
								Username: b.configuration.Nats.Username,
								Password: b.configuration.Nats.Password,
							},
						},

						Debug: b.configuration.App.Debug,
						Trace: b.configuration.App.Verbose,

						DontListen:             false,
						MaxConn:                natsd.DEFAULT_MAX_CONNECTIONS,
						MaxSubs:                natsd.DEFAULT_MAX_CONNECTIONS,
						JetStream:              true,
						JetStreamMaxMemory:     -1,
						JetStreamMaxStore:      -1,
						StoreDir:               storeDir,
						DisableJetStreamBanner: true,
					}

					if b.configuration.Natsd.Clustered {
						fmt.Println("Clustering enabled")
						options.Cluster = natsd.ClusterOpts{
							Host:           b.configuration.Natsd.Cluster.Host,
							Port:           b.configuration.Natsd.Cluster.Port,
							Name:           b.configuration.Natsd.Cluster.Name,
							ConnectRetries: 3,
							ConnectBackoff: true,
							PoolSize:       b.configuration.Natsd.Replicas,
							Compression: natsd.CompressionOpts{
								Mode:          "s2_auto",
								RTTThresholds: []time.Duration{100 * time.Millisecond, 300 * time.Millisecond, 500 * time.Millisecond},
							},
						}

						routes := RoutesFromStr(b.configuration.Natsd.RoutesStr)

						options.Routes = routes
					}

					// Check if we are running inside the mesh, if so, use the CustomDialer option
					if b.configuration.Platform.Mesh.Enabled {
						if !nebulav1.IsBound {
							fmt.Println("You have enabled the mesh network for Nats traffic, however, you haven't bound Nebula. Please add the Nebula binding.")
							panic("Missing Nebula binding")
						}
					}

					server, err := natsd.NewServer(&options)
					if err != nil {
						fmt.Println("natsd error: ", err)
						panic(err)
					}

					logger := zaploggerv1.Bound.SugaredLoggerWrapper
					server.SetLoggerV2(logger, b.configuration.App.Debug, b.configuration.App.Verbose, false)

					// Start the NATS server
					go server.Start()

					if !server.ReadyForConnections(10 * time.Second) {
						fmt.Println("NATS server not ready within 10 seconds")
						panic("Server not ready within 10 seconds")
					}

					acc, err := server.LookupAccount(DEFAULT_CONNECTOR_ACCOUNT)
					if err != nil {
						fmt.Println("lookup account failed: " + err.Error())
					}

					// Keep limits empty for now, Nats will create them dynamically
					limits := map[string]natsd.JetStreamAccountLimits{}
					err = acc.EnableJetStream(limits)
					if err != nil {
						panic("Cannot enable JetStream for connector account: " + err.Error())
					}

					_nats, err := nats.Connect(fmt.Sprintf("nats://%s:%s@%s:%d", b.configuration.Nats.Username, b.configuration.Nats.Password, options.Host, options.Port))
					if err != nil {
						fmt.Println("error connecting to NATS server at localhost port: " + strconv.Itoa(options.Port) + " " + err.Error())
						panic(err)
					}
					b.Nats = _nats

					jsc, err := _nats.JetStream()
					if err != nil {
						fmt.Println(err.Error())
						panic("Cannot get Jetstream from Nats Context")
					}
					b.JetStreamContext = &jsc

					// Create a JetStream management interface
					js, err := jetstream.New(_nats)
					if err != nil {
						fmt.Println(err.Error())
						panic("Cannot configure Jetstream")
					}
					b.JetStream = &js

					Bound = &Binding{
						server:           server,
						Nats:             _nats,
						JetStreamContext: &jsc,
						JetStream:        &js,
						configuration:    b.configuration,
					}

					bindings.Registered[b.Name()] = Bound

					bindings = b.RegisterSpecListeners(bindings)

					fmt.Println("NATS TCP listening on " + strconv.Itoa(options.Port))

				case false:
					servers := strings.Replace(strings.Trim(fmt.Sprint(b.configuration.Nats.Options.Servers), "[]"), " ", ",", -1)

					// Check if we are running inside the mesh, if so, use the CustomDialer option
					if b.configuration.Platform.Mesh.Enabled {
						if !nebulav1.IsBound {
							fmt.Println("You have enabled the mesh network for Nats traffic, however, you haven't bound Nebula. Please add the Nebula binding.")
							panic("Missing Nebula binding")
						}

						natsOptions = append(natsOptions, nats.SetCustomDialer(nebulav1.Bound.MeshSocket))
					}

					natsOptions = append(natsOptions, nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
						apexlog.Info("Disconnected due to: " + err.Error())
					}))
					natsOptions = append(natsOptions, nats.ReconnectHandler(func(nc *nats.Conn) {
						apexlog.Info("Reconnected to " + nc.ConnectedUrl())
					}))
					natsOptions = append(natsOptions, nats.ClosedHandler(func(_ *nats.Conn) {
						apexlog.Info("Connection closed.")
					}))

					maxRetries := 180              // -1 for infinite retries
					retryDelay := 10 * time.Second // wait 3 seconds between attempts

					_nats, jsc, js, err := connectWithRetry(servers, maxRetries, retryDelay, natsOptions...)
					if err != nil {
						panic(err)
					}

					b.Nats = _nats
					b.JetStreamContext = &jsc

					Bound = &Binding{
						Registry:           b.Registry,
						SpecEventListeners: b.SpecEventListeners,
						Listeners:          b.Listeners,
						Nats:               _nats,
						JetStreamContext:   &jsc,
						JetStream:          &js,
						configuration:      b.configuration,
					}

					bindings.Registered[b.Name()] = Bound

					bindings = b.RegisterSpecListeners(bindings)
					bindings = b.RegisterSpecBatchListeners(bindings)
				}
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Nats Node already bound")
	}

	return bindings
}

// GetBinding returns the currently bound instance of the Binding struct.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close shuts down the NATS node connection and drains/unsubscribes all active listeners.
func (b *Binding) Close() error {
	fmt.Println("Shutting down Nats Node connection for: ")

	for _, listener := range b.Listeners {
		err := listener.Drain()
		if err != nil {
			fmt.Println(err)
		}

		err = listener.Unsubscribe()
		if err != nil {
			fmt.Println(err)
		}
	}

	if b.Nats != nil {
		err := b.Nats.Drain()
		if err != nil {
			return err
		}
	}

	if b.server != nil {
		b.server.Shutdown()
	}

	return nil
}

// RegisterSpecListeners registers specification event listeners by configuring them and associating them with bindings.
func (b *Binding) RegisterSpecListeners(bindings *sdkv2betalib.Bindings) *sdkv2betalib.Bindings {
	for _, listener := range b.SpecEventListeners {
		configuration := listener.Configure()
		if configuration == nil {
			fmt.Println("Please configure the Listener")
			panic("Misconfigured")
		}

		name := ""
		if configuration.JetstreamConfiguration.Name == "" && configuration.JetstreamConfiguration.Durable == "" {
			fmt.Println("Either the Name or the Durable name is required")
			panic("Misconfigured")
		}

		if configuration.JetstreamConfiguration.Name != "" {
			name = configuration.JetstreamConfiguration.Durable
		}

		// Use the durable name if set
		if configuration.JetstreamConfiguration.Durable != "" {
			name = configuration.JetstreamConfiguration.Durable
		}

		bindings.RegisteredListenableChannels[name] = listener
	}

	return bindings
}

// RegisterSpecBatchListeners registers specification event listeners by configuring them and associating them with bindings.
func (b *Binding) RegisterSpecBatchListeners(bindings *sdkv2betalib.Bindings) *sdkv2betalib.Bindings {
	for _, listener := range b.SpecEventBatchListeners {
		configuration := listener.Configure()
		if configuration == nil {
			fmt.Println("Please configure the Listener")
			panic("Misconfigured")
		}

		name := ""
		if configuration.JetstreamConfiguration.Name == "" && configuration.JetstreamConfiguration.Durable == "" {
			fmt.Println("Either the Name or the Durable name is required")
			panic("Misconfigured")
		}

		if configuration.JetstreamConfiguration.Name != "" {
			name = configuration.JetstreamConfiguration.Durable
		}

		// Use the durable name if set
		if configuration.JetstreamConfiguration.Durable != "" {
			name = configuration.JetstreamConfiguration.Durable
		}

		bindings.RegisteredListenableChannels[name] = listener
	}

	return bindings
}

func connectWithRetry(url string, maxRetries int, retryDelay time.Duration, options ...nats.Option) (*nats.Conn, nats.JetStreamContext, jetstream.JetStream, error) {
	var nc *nats.Conn
	var jsc nats.JetStreamContext
	var js jetstream.JetStream
	var err error

	for attempt := 1; maxRetries == -1 || attempt <= maxRetries; attempt++ {
		fmt.Printf("Connecting to NATS (attempt %d)...\n", attempt)

		nc, err = nats.Connect(url)
		if err != nil {
			fmt.Printf("NATS connection failed: %v\n", err)
			time.Sleep(retryDelay)
			continue
		}

		jsc, err = nc.JetStream()
		if err != nil {
			fmt.Printf("JetStream setup failed: %v\n", err)
			nc.Close()
			time.Sleep(retryDelay)
			continue
		}
		js, err = jetstream.New(nc)
		if err != nil {
			fmt.Printf("JetStream setup failed: %v\n", err)
			nc.Close()
			time.Sleep(retryDelay)
			continue
		}

		fmt.Println("Connected to NATS and JetStream is ready.")
		return nc, jsc, js, nil
	}

	return nil, nil, nil, fmt.Errorf("failed to connect to NATS after %d attempts", maxRetries)
}

// RoutesFromStr parses route URLs from a string
func RoutesFromStr(routesStr string) []*url.URL {
	routes := strings.Split(routesStr, ",")
	if len(routes) == 0 {
		return nil
	}
	routeUrls := []*url.URL{}
	for _, r := range routes {
		r = strings.TrimSpace(r)
		u, _ := url.Parse(r)
		routeUrls = append(routeUrls, u)
	}
	return routeUrls
}
