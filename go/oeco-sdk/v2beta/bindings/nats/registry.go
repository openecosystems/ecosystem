package natsnodev1

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	sdkv2betalib "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta"
)

// RegisterEventStreams initializes and registers event streams with specified scopes and subjects in the configuration.
// It creates or updates streams by adding appropriate prefixes to subjects and setting up stream configurations.
func RegisterEventStreams(environmentName string, streams []jetstream.StreamConfig) {
	if environmentName == "" {
		panic("environment name cannot be empty")
	}

	var scopes [3]string
	scopes[0] = NewInternalStream().StreamPrefix()
	scopes[1] = NewInboundStream().StreamPrefix()
	scopes[2] = NewOutboundStream().StreamPrefix()

	for _, cfg := range streams {
		name := cfg.Name
		subjects := cfg.Subjects

		for _, scope := range scopes {
			var _subjects []string
			cfg.Subjects = nil

			// Prepend scope to subjects
			for _, subject := range subjects {
				s := scope + "-" + subject
				_subjects = append(_subjects, s)
			}

			cfg.Name = GetStreamName(environmentName, scope, name)
			cfg.Subjects = _subjects

			err := createOrUpdateStream(cfg)
			if err != nil {
				fmt.Println("Found error creating stream: " + err.Error())
				panic(err)
			}
		}
	}
}

func createOrUpdateStream2(cfg jetstream.StreamConfig) error {
	js := Bound.JetStream
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := js.Stream(ctx, cfg.Name)
	if err != nil {
		// Stream not found — create new
		if errors.Is(err, jetstream.ErrStreamNotFound) {
			fmt.Println("Creating stream:", cfg.Name)
			if _, err := js.CreateOrUpdateStream(ctx, cfg); err != nil {
				return fmt.Errorf("failed to create stream %s: %w", cfg.Name, err)
			}
			return nil
		}
		// Any other error
		return fmt.Errorf("error fetching stream info for %s: %w", cfg.Name, err)
	}

	// Compare existing config with desired config
	_cfg := info.CachedInfo().Config
	if !streamConfigsEqual(&_cfg, &cfg) {
		fmt.Printf("Updating stream %s (config has changed)\n", cfg.Name)
		if _, err := js.CreateOrUpdateStream(ctx, cfg); err != nil {
			return fmt.Errorf("failed to update stream %s: %w", cfg.Name, err)
		}
	} else {
		fmt.Printf("Stream %s is up to date, no changes needed\n", cfg.Name)
	}

	return nil
}

func streamConfigsEqual(a, b *jetstream.StreamConfig) bool {
	// ignore zero values (JetStream may fill defaults)
	cleanA := normalizeStreamConfig(a)
	cleanB := normalizeStreamConfig(b)

	if cleanA.Replicas != cleanB.Replicas {
		return false
	}

	if cleanA.MaxMsgs != cleanB.MaxMsgs {
		return false
	}

	if cleanA.Name != cleanB.Name {
		return false
	}

	if cleanA.Name != cleanB.Name {
		return false
	}

	if !reflect.DeepEqual(cleanA.Subjects, cleanB.Subjects) {
		return false
	}

	return true
	// return reflect.DeepEqual(cleanA, cleanB)
}

func normalizeStreamConfig(cfg *jetstream.StreamConfig) *jetstream.StreamConfig {
	c := *cfg
	if c.MaxAge == 0 {
		c.MaxAge = 0
	}
	if c.MaxMsgs == 0 {
		c.MaxMsgs = -1
	}
	return &c
}

// createOrUpdateStream creates a new stream or updates an existing one using the provided StreamConfig.
func createOrUpdateStream(cfg jetstream.StreamConfig) error {
	// Check if stream exists
	js := Bound.JetStream
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := js.Stream(ctx, cfg.Name)
	if err != nil {
		if errors.Is(err, jetstream.ErrStreamNotFound) {
			// if stream does not exist, create it
			fmt.Println("Creating stream " + cfg.Name)
			_, e := js.CreateOrUpdateStream(ctx, cfg)
			if e != nil {
				fmt.Println(e.Error())
				fmt.Println("issue creating stream when determined it does not exist: " + cfg.Name)
				return sdkv2betalib.ErrServerInternal.WithInternalErrorDetail(errors.New("issue creating stream when determined it does not exist: " + cfg.Name))
			}
		} else {
			var apiErr *nats.APIError
			var jsApiErr *jetstream.APIError
			if errors.Is(err, nats.ErrJetStreamNotEnabled) || errors.Is(err, ErrTimeout) || errors.Is(err, context.DeadlineExceeded) {
				// if creating consumer failed, retry
				fmt.Println("oeco: if creating consumer failed, retry")
			} else if errors.As(err, &apiErr) {
				if apiErr.ErrorCode == nats.JSErrCodeInsufficientResourcesErr {
					// retry for insufficient resources, as it may mean that client is connected to a running
					// server in cluster while the server hosting R1 JetStream resources is restarting
					fmt.Println("oeco: retry for insufficient resources, as it may mean that client is connected to a running server in cluster while the server hosting R1 JetStream resources is restarting")
				} else if apiErr.ErrorCode == nats.JSErrCodeJetStreamNotAvailable {
					// retry if JetStream meta leader is temporarily unavailable
					fmt.Println("oeco: retry if JetStream meta leader is temporarily unavailable")
				} else {
					fmt.Println("oeco: other nats api issue creating or updating stream")
					fmt.Println(err.Error())
				}
			} else if errors.As(err, &jsApiErr) {
				if jsApiErr.ErrorCode == 10008 {
					// retry if JetStream meta leader is temporarily unavailable
					fmt.Println("oeco: waiting for JetStream meta leader to be available")
				} else {
					fmt.Println("oeco: other Jetstream API issue received")
					fmt.Println(jsApiErr.Error())
				}
			} else {
				fmt.Println("oeco: Other issue creating or updating stream besides not found")
				fmt.Println(err.Error())
			}

		}
		// Any other error
		fmt.Printf("error fetching stream info for %s", cfg.Name)
		return nil
	}

	// Compare existing config with desired config
	_cfg := info.CachedInfo().Config
	if !streamConfigsEqual(&_cfg, &cfg) {
		fmt.Printf("Updating stream %s (config has changed)\n", cfg.Name)
		if _, err := js.CreateOrUpdateStream(ctx, cfg); err != nil {
			return fmt.Errorf("failed to update stream %s: %w", cfg.Name, err)
		}
	} else {
		// fmt.Printf("Stream %s is up to date, no changes needed\n", cfg.Name)
	}

	return nil
}

// GetStreamName generates a stream name by concatenating the environment, scope, and entity name with hyphens.
func GetStreamName(env string, scope string, entityName string) string {
	return env + "-" + scope + "-" + entityName
}

// GetMultiplexedRequestSubjectName generates a subject name by combining the provided scope and subject name with a "req." prefix.
func GetMultiplexedRequestSubjectName(scope string, subjectName string, procedure string) string {
	return "req." + scope + "-" + subjectName + "." + procedure
}

// GetSubjectName generates a subject name by combining the provided scope and subject name with a hyphen separator.
func GetSubjectName(scope string, subjectName string, procedure string) string {
	return scope + "-" + subjectName + "." + procedure
}

// GetStreamResponseSubjectName generates a subject name by combining the provided scope and subject name with a hyphen separator.
func GetStreamResponseSubjectName(scope string, subjectName string, procedure string, messageID string) string {
	return scope + "-" + subjectName + "." + procedure + "." + messageID
}

// GetQueueGroupName generates a queue group name by combining the given scope and entityName with a predefined prefix "req.".
func GetQueueGroupName(scope string, entityName string, procedure string) string {
	return "req." + scope + "-" + entityName + "." + procedure
}
