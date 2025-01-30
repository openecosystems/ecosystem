package natsnodev2

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"dario.cat/mergo"

	serverv2alphalib "libs/public/go/server/v2alpha"

	natsd "github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// NatsdServerName defines the name of the NATS server.
// NatsdServerHost specifies the host address of the NATS server.
// NatsdServerPort represents the port number for the NATS server.
// NatsdServerJetstreamStoreDir sets the directory for JetStream data storage.
const (
	NatsdServerName              = "natsv2"
	NatsdServerHost              = "localhost"
	NatsdServerPort              = 4222
	NatsdServerJetstreamStoreDir = "./nats-jetstream-data"
)

// ResolvedConfiguration holds the resolved configuration for the application.
// NatsServers defines the default NATS server addresses.
var (
	ResolvedConfiguration *Configuration
	NatsServers           = []string{"nats://127.0.0.1:4222"}
)

// Nats represents the configuration for NATS connectivity including mesh and specific connection options.
type Nats struct {
	Options nats.Options
}

// Natsd configures an embedded NATS server with customizable options for use in the application.
// Enabled specifies whether the embedded NATS server is active or not.
// Options defines the configuration settings for the embedded NATS server through natsd.Options.
type Natsd struct {
	Enabled bool
	Options natsd.Options
}

// EventStreamRegistry is a structure that holds a list of jetstream stream configurations.
type EventStreamRegistry struct {
	Streams []jetstream.StreamConfig
}

// Configuration represents the overall settings structure comprising NATS, NATS server options, and stream registry configurations.
type Configuration struct {
	Mesh                serverv2alphalib.Mesh
	Nats                Nats
	Natsd               Natsd
	EventStreamRegistry EventStreamRegistry

	err error
}

// ResolveConfiguration resolves and merges the binding's configuration with default settings, validating streams and settings.
func (b *Binding) ResolveConfiguration() {
	var c Configuration
	dc := b.GetDefaultConfiguration().(Configuration)
	sdkv2alphalib.Resolve(&c, dc)
	b.configuration = &c
	ResolvedConfiguration = &c

	dsc := jetstream.StreamConfig{
		MaxMsgs:           -1,
		MaxBytes:          -1,
		Discard:           0,
		MaxAge:            9151516080000000000, // 290 years is the max Nats supports,
		MaxMsgsPerSubject: -1,
		MaxMsgSize:        -1,
		Storage:           0,
		Replicas:          1, // TODO: Review this default
		NoAck:             false,
		Duplicates:        60 * time.Second, //"2m0s"
		DenyDelete:        true,
		DenyPurge:         true,
		AllowRollup:       false,
	}

	var mergedJsc []jetstream.StreamConfig
	var errs []error
	for _, r := range b.configuration.EventStreamRegistry.Streams {
		if err := mergo.Merge(&r, dsc); err != nil {
			fmt.Println("Error merging nats stream configuration:", err)
			errs = append(errs, err)
		}
		mergedJsc = append(mergedJsc, r)
	}
	if len(errs) > 0 {
		fmt.Println("nats configuration error: ", errs)
		b.configuration.err = errors.Join(errs...)
	}

	c.EventStreamRegistry = EventStreamRegistry{mergedJsc}
	b.configuration.EventStreamRegistry = EventStreamRegistry{mergedJsc}
	ResolvedConfiguration.EventStreamRegistry = EventStreamRegistry{mergedJsc}
}

// ValidateConfiguration checks if NATS and stream configurations are valid, returning an error if validation fails.
func (b *Binding) ValidateConfiguration() error {
	if !ResolvedConfiguration.Natsd.Enabled {
		return nil
	}

	var errs []error

	if len(b.configuration.Natsd.Options.LeafNode.Remotes) == 0 {
		errs = append(errs, errors.New(`missing leaf node remotes configuration. An example is:
natsd:
  options:
    leafNode:
      remotes:
        - urls:
            scheme: "tls"
            host:   "connect.ngs.global"
          credentials: "./example.creds"`))
	}

	for i, s := range b.configuration.EventStreamRegistry.Streams {
		if s.Name == "" {
			errs = append(errs, errors.New("missing stream name for item with index: "+strconv.Itoa(i)))
		}

		if len(s.Subjects) == 0 {
			errs = append(errs, errors.New("missing array of subjects for item with index: "+strconv.Itoa(i)))
		}
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

// GetDefaultConfiguration returns the default configuration object for the Binding, including NATS and JetStream settings.
func (b *Binding) GetDefaultConfiguration() interface{} {
	cfg := sdkv2alphalib.ResolvedConfiguration

	return Configuration{
		Mesh: serverv2alphalib.Mesh{
			Enabled: false,
		},
		Nats: Nats{
			Options: nats.Options{
				Servers: NatsServers,
				// TODO: Review how to tie Mesh with this
				//Dialer: &net.Dialer{
				//	Timeout:   0,
				//	Deadline:  time.Time{},
				//	LocalAddr: nil,
				//	KeepAliveConfig: net.KeepAliveConfig{
				//		Enable:   false,
				//		Idle:     0,
				//		Interval: 0,
				//		Count:    0,
				//	},
				//},
				//CustomDialer:                nil,
				Compression: false,
			},
		},
		Natsd: Natsd{
			Enabled: false,
			Options: natsd.Options{
				ServerName: NatsdServerName,
				Host:       NatsdServerHost,
				Port:       NatsdServerPort,
				DontListen: false,
				Trace:      cfg.App.Trace,
				Debug:      cfg.App.Debug,
				MaxConn:    -1,
				MaxSubs:    -1,
				LeafNode: natsd.LeafNodeOpts{
					Remotes: nil,
				},
				JetStream:              true,
				JetStreamMaxMemory:     -1,
				JetStreamMaxStore:      -1,
				StoreDir:               NatsdServerJetstreamStoreDir,
				DisableJetStreamBanner: true,
				// TLSConfig:              &tls.Config{},
				// AllowNonTLS:            true,
			},
		},
		EventStreamRegistry: EventStreamRegistry{
			// Streams: mergedJsc,
		},
	}
}
