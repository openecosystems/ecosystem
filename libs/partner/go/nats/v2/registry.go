package natsnodev2

import (
	"context"
	"errors"
	"fmt"
	"time"

	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"github.com/nats-io/nats.go/jetstream"
)

// RegisterEventStreams initializes and registers event streams with specified scopes and subjects in the configuration.
// It creates or updates streams by adding appropriate prefixes to subjects and setting up stream configurations.
func RegisterEventStreams() {
	var scopes [3]string
	scopes[0] = NewInternalStream().StreamPrefix()
	scopes[1] = NewInboundStream().StreamPrefix()
	scopes[2] = NewOutboundStream().StreamPrefix()

	conf := ResolvedConfiguration
	for _, cfg := range conf.EventStreamRegistry.Streams {
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

			cfg.Name = GetStreamName(conf.App.EnvironmentName, scope, name)
			cfg.Subjects = _subjects

			err := createOrUpdateStream(cfg)
			if err != nil {
				fmt.Println("Found error creating stream: " + err.Error())
				panic(err)
			}
		}
	}
}

// createOrUpdateStream creates a new stream or updates an existing one using the provided StreamConfig.
func createOrUpdateStream(cfg jetstream.StreamConfig) error {
	// Check if stream exists
	js := *Bound.JetStream
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	_, err := js.Stream(ctx, cfg.Name)
	if err != nil {
		if errors.Is(err, jetstream.ErrStreamNotFound) {
			// if stream does not exist, create it
			fmt.Println("Creating stream " + cfg.Name)
			_, e := js.CreateOrUpdateStream(ctx, cfg)
			if e != nil {
				fmt.Println(e.Error())
				fmt.Println("issue creating stream when determined it does not exist: " + cfg.Name)
				return sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("issue creating stream when determined it does not exist: " + cfg.Name))
			}
		} else {
			fmt.Println("Other error creating or updating stream besides not found")
			fmt.Println(err.Error())
		}
	} else {
		_, err3 := js.CreateOrUpdateStream(ctx, cfg)
		if err3 != nil {
			return err3
		}

		return nil
	}

	return nil
}

// GetStreamName generates a stream name by concatenating the environment, scope, and entity name with hyphens.
func GetStreamName(env string, scope string, entityName string) string {
	return env + "-" + scope + "-" + entityName
}

// GetMultiplexedRequestSubjectName generates a subject name by combining the provided scope and subject name with a "req." prefix.
func GetMultiplexedRequestSubjectName(scope string, subjectName string) string {
	return "req." + scope + "-" + subjectName
}

// GetSubjectName generates a subject name by combining the provided scope and subject name with a hyphen separator.
func GetSubjectName(scope string, subjectName string) string {
	return scope + "-" + subjectName
}

// GetQueueGroupName generates a queue group name by combining the given scope and entityName with a predefined prefix "req.".
func GetQueueGroupName(scope string, entityName string) string {
	return "req." + scope + "-" + entityName
}
