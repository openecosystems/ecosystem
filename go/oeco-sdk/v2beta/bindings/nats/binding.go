package natsnodev1

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	natsd "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"

	apexlog "github.com/apex/log"
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
	nebulav1 "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/bindings/nebula"
)

// Binding represents a structure managing NATS connections, JetStream instances, and event stream configurations.
type Binding struct {
	Registry                map[string]nats.StreamConfig
	SpecEventListeners      []SpecEventListener
	SpecEventBatchListeners []SpecEventBatchListener
	Listeners               map[string]*nats.Subscription
	Nats                    *nats.Conn
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

// Bind initializes the Binding instance, configures NATS or JetStream connections, and registers the binding in Bindings.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2betalib.Bindings) *sdkv2betalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				switch b.configuration.Natsd.Enabled {
				case true:
					options := b.configuration.Natsd.Options

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

					// Start NATS server
					go func() {
						if err := natsd.Run(server); err != nil {
							fmt.Println("SpecError running embedded NATS server: " + err.Error())
							panic(err)
						}
					}()

					if !server.ReadyForConnections(time.Minute) {
						fmt.Println("NATS server not ready within 60 seconds")
						panic("Server not ready within 60 seconds")
					}

					_nats, err := nats.Connect(fmt.Sprintf("nats://%s:%d", options.Host, options.Port))
					if err != nil {
						fmt.Println("error connecting to NATS server at localhost port: " + strconv.Itoa(options.Port) + " " + err.Error())
					}
					b.Nats = _nats

					// Create a JetStream management interface
					js, err := jetstream.New(_nats)
					if err != nil {
						fmt.Println(err.Error())
						panic("Cannot configure Jetstream")
					}
					b.JetStream = &js

					Bound = &Binding{
						server:        server,
						Nats:          _nats,
						JetStream:     &js,
						configuration: b.configuration,
					}

					bindings.Registered[b.Name()] = Bound

					bindings = b.RegisterSpecListeners(bindings)

					fmt.Println("NATS TCP listening on " + strconv.Itoa(options.Port))

					RegisterEventStreams()
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

					_nats, js, err := connectWithRetry(servers, maxRetries, retryDelay, natsOptions...)
					if err != nil {
						panic(err)
					}

					//_nats, err := nats.Connect(servers, natsOptions...)
					//if err != nil {
					//	fmt.Println(err.Error())
					//	panic("Cannot connect to NATS")
					//}
					b.Nats = _nats

					// Create a JetStream management interface
					//js, err := jetstream.New(_nats)
					//if err != nil {
					//	fmt.Println(err.Error())
					//	panic("Cannot configure Jetstream")
					//}
					b.JetStream = &js

					Bound = &Binding{
						Registry:           b.Registry,
						SpecEventListeners: b.SpecEventListeners,
						Listeners:          b.Listeners,
						Nats:               _nats,
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

func connectWithRetry(url string, maxRetries int, retryDelay time.Duration, options ...nats.Option) (*nats.Conn, jetstream.JetStream, error) {
	var nc *nats.Conn
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

		js, err = jetstream.New(nc)
		if err != nil {
			fmt.Printf("JetStream setup failed: %v\n", err)
			nc.Close()
			time.Sleep(retryDelay)
			continue
		}

		fmt.Println("Connected to NATS and JetStream is ready.")
		return nc, js, nil
	}

	return nil, nil, fmt.Errorf("failed to connect to NATS after %d attempts", maxRetries)
}
