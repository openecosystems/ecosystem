package configurationv2alphalib

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/nats-io/nats.go/jetstream"
	protopb "google.golang.org/protobuf/proto"
	"libs/public/go/connector/v2alpha"
	"libs/public/go/protobuf/gen/platform/configuration/v2alpha"
	"libs/public/go/sdk/v2alpha"
)

const configPrefix = ".cfg."

type AdaptiveConfigurationControl struct {
	ConfigStore *jetstream.KeyValue
	FileSystem  *sdkv2alphalib.FileSystem
}

func NewAdaptiveConfigurationControl(configStore *jetstream.KeyValue) *AdaptiveConfigurationControl {
	return &AdaptiveConfigurationControl{
		ConfigStore: configStore,
		FileSystem:  sdkv2alphalib.NewFileSystem(),
	}
}

func (acc AdaptiveConfigurationControl) SavePlatformConfiguration(ctx context.Context, orgOrWorkspaceOrConfigGroupId string, bytes []byte) error {
	cs := *acc.ConfigStore

	key, err := acc.getKey(orgOrWorkspaceOrConfigGroupId)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	_, err2 := cs.Put(ctx, key, bytes)
	if err2 != nil {
		fmt.Println(err2)
		return nil
	}

	return nil
}

func (acc AdaptiveConfigurationControl) GetPlatformConfiguration(ctx context.Context, orgOrWorkspaceOrConfigGroupId string) (*configurationv2alphapb.Configuration, error) {
	var errs []string

	if orgOrWorkspaceOrConfigGroupId == "" {
		errs = append(errs, "Please provide a valid organization, workspace or configuration group ID")
	}

	if len(errs) > 0 {
		fmt.Println("adaptive configuration control error: ", errs)
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New(strings.Join(errs, "; ")))
	}

	key, err := acc.getKey(orgOrWorkspaceOrConfigGroupId)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	cs := *acc.ConfigStore
	entry, err := cs.Get(ctx, key)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	var c configurationv2alphapb.Configuration
	err2 := protopb.Unmarshal(entry.Value(), &c)
	if err2 != nil {
		fmt.Println(err2.Error())
		return nil, err2
	}

	return &c, nil
}

func (acc AdaptiveConfigurationControl) getKey(orgOrWorkspaceOrConfigGroupId string) (string, error) {
	rootConfig := connectorv2alphalib.ResolvedConfiguration

	// local-1.cfg.organization123.platform.configuration.v2alpha.Configuration
	key := rootConfig.App.EnvironmentName + configPrefix + orgOrWorkspaceOrConfigGroupId

	return key, nil
}
