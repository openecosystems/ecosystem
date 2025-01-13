package natsnodev2

import (
  "errors"
  "fmt"
  "strconv"
  "time"

  "libs/public/go/sdk/v2alpha"

  "dario.cat/mergo"

  natsd "github.com/nats-io/nats-server/v2/server"
  "github.com/nats-io/nats.go"
  "github.com/nats-io/nats.go/jetstream"
)

const (
  NatsdServerName              = "natsv2"
  NatsdServerHost              = "localhost"
  NatsdServerPort              = 4222
  NatsdServerJetstreamStoreDir = "./nats-jetstream-data"
)

var (
  ResolvedConfiguration *Configuration
  NatsServers           = []string{"nats://127.0.0.1:4222"}
)

type Nats struct {
  Mesh    bool
  Options nats.Options
}

// Natsd represents the configuration for an embedded NATS server.
type Natsd struct {
  Enabled bool
  Options natsd.Options
}

// EventStreamRegistry holds the configuration details for a set of event streams.
type EventStreamRegistry struct {
  Streams []jetstream.StreamConfig
}

type Configuration struct {
  Nats                Nats
  Natsd               Natsd
  EventStreamRegistry EventStreamRegistry

  err error
}

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

func (b *Binding) ValidateConfiguration() error {
  if !ResolvedConfiguration.Natsd.Enabled {
    return nil
  }

  var errs []error

  if b.configuration.Natsd.Options.LeafNode.Remotes == nil || len(b.configuration.Natsd.Options.LeafNode.Remotes) == 0 {
    errs = append(errs, errors.New(`
missing leaf node remotes configuration. An example is:
natsd:
  options:
    leafNode:
      remotes:
        - urls:
            scheme: "tls"
            host:   "connect.ngs.global"
          credentials: "./example.creds"
`))
  }

  for i, s := range b.configuration.EventStreamRegistry.Streams {
    if s.Name == "" {
      errs = append(errs, errors.New("missing stream name for item with index: "+strconv.Itoa(i)))
    }

    if s.Subjects == nil || len(s.Subjects) == 0 {
      errs = append(errs, errors.New("missing array of subjects for item with index: "+strconv.Itoa(i)))
    }
  }

  if len(errs) > 0 {
    return errors.Join(errs...)
  }

  return nil
}

func (b *Binding) GetDefaultConfiguration() interface{} {
  cfg := sdkv2alphalib.ResolvedConfiguration

  return Configuration{
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
        //TLSConfig:              &tls.Config{},
        //AllowNonTLS:            true,
      },
    },
    EventStreamRegistry: EventStreamRegistry{
      // Streams: mergedJsc,
    },
  }
}
