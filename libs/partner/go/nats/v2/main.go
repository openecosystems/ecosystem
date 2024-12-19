package natsnodev2

import (
	"context"
	"crypto/tls"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"libs/public/go/sdk/v2alpha"

	natsd "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// Binding struct that holds binding specific fields
type Binding struct {
	Registry           map[string]nats.StreamConfig
	SpecEventListeners []SpecEventListener
	Listeners          map[string]*nats.Subscription
	Nats               *nats.Conn
	JetStream          *jetstream.JetStream

	server        *natsd.Server
	configuration *Configuration
}

var (
	natsOptions []nats.Option
	jsOptions   []jetstream.JetStreamOpt
	Bound       *Binding
	BindingName = "NATS_NODE_BINDING"
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	return nil
}

func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				switch ResolvedConfiguration.Natsd.Enabled {
				case true:
					options := ResolvedConfiguration.Natsd.Options

					options.TLSConfig = &tls.Config{}
					options.AllowNonTLS = true

					server, err := natsd.NewServer(&options)
					if err != nil {
						fmt.Println("natsd error: ", err)
						panic(err)
					}

					// Start NATS server
					go func() {
						if err := natsd.Run(server); err != nil {
							fmt.Println("Error running embedded NATS server: " + err.Error())
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
						server:    server,
						Nats:      _nats,
						JetStream: &js,
					}

					bindings.Registered[b.Name()] = Bound

					fmt.Println("Nats Leaf Node Server started successfully. NATS listening on port " + strconv.Itoa(options.Port))

					RegisterEventStreams()
				case false:
					servers := strings.Replace(strings.Trim(fmt.Sprint(ResolvedConfiguration.Nats.Options.Servers), "[]"), " ", ",", -1)

					_nats, err := nats.Connect(servers, natsOptions...)
					if err != nil {
						fmt.Println(err.Error())
						panic("Cannot connect to NATS")
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
						Registry:           b.Registry,
						SpecEventListeners: b.SpecEventListeners,
						Listeners:          b.Listeners,
						Nats:               _nats,
						JetStream:          &js,
					}

					bindings.Registered[b.Name()] = Bound

					for _, listener := range b.SpecEventListeners {
						configuration := listener.GetConfiguration()
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
				}
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("Nats Node already bound")
	}

	return bindings
}

func (b *Binding) GetBinding() interface{} {
	return Bound
}

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
