package natsnodev2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	"libs/public/go/sdk/v2alpha"
)

func RegisterEventStreams() {
	rootConfig := sdkv2alphalib.ResolvedConfiguration

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

			cfg.Name = GetStreamName(rootConfig.App.EnvironmentName, scope, name)
			cfg.Subjects = _subjects

			err := createOrUpdateStream(cfg)
			if err != nil {
				fmt.Println("Found error creating stream: " + err.Error())
				panic(err)
			}
		}

	}
}

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

func GetStreamName(env string, scope string, entityName string) string {
	return env + "-" + scope + "-" + entityName
}

func GetMultiplexedRequestSubjectName(scope string, subjectName string) string {
	return "req." + scope + "-" + subjectName
}

func GetSubjectName(scope string, subjectName string) string {
	return scope + "-" + subjectName
}

func GetQueueGroupName(scope string, entityName string) string {
	return "req." + scope + "-" + entityName
}
