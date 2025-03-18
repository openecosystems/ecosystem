package configurationv2alphalib

import (
	"context"
	"errors"
	"fmt"
	"strings"

	configurationv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/gen/platform/configuration/v2alpha"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	"github.com/nats-io/nats.go/jetstream"
	protopb "google.golang.org/protobuf/proto"
)

const configPrefix = ".cfg."

// AdaptiveConfigurationControl provides tools for managing adaptive configurations using a KeyValue store and a filesystem.
// It allows saving and retrieving configuration data for organizations, workspaces, or configuration groups.
type AdaptiveConfigurationControl struct {
	ConfigStore *jetstream.KeyValue
	FileSystem  *sdkv2alphalib.FileSystem
}

// NewAdaptiveConfigurationControl initializes and returns a new AdaptiveConfigurationControl instance.
// It requires a KeyValue store for storing configuration and utilizes a new filesystem instance from sdkv2alphalib.
func NewAdaptiveConfigurationControl(configStore *jetstream.KeyValue) *AdaptiveConfigurationControl {
	return &AdaptiveConfigurationControl{
		ConfigStore: configStore,
		FileSystem:  sdkv2alphalib.NewFileSystem(),
	}
}

// SavePlatformConfiguration stores configuration data for a specific organization, workspace, or group in a KeyValue store.
func (acc AdaptiveConfigurationControl) SavePlatformConfiguration(ctx context.Context, orgOrWorkspaceOrConfigGroupID string, bytes []byte) error {
	cs := *acc.ConfigStore

	key, err := acc.getKey(orgOrWorkspaceOrConfigGroupID)
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

// GetPlatformConfiguration retrieves the platform configuration for the specified organization, workspace, or configuration group ID.
// Returns a Configuration object and error if any issue occurs during retrieval or unmarshalling the configuration.
func (acc AdaptiveConfigurationControl) GetPlatformConfiguration(ctx context.Context, orgOrWorkspaceOrConfigGroupID string) (*configurationv2alphapb.Configuration, error) {
	var errs []string

	if orgOrWorkspaceOrConfigGroupID == "" {
		errs = append(errs, "Please provide a valid organization, workspace or configuration group ID")
	}

	if len(errs) > 0 {
		fmt.Println("adaptive configuration control error: ", errs)
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New(strings.Join(errs, "; ")))
	}

	key, err := acc.getKey(orgOrWorkspaceOrConfigGroupID)
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

// getKey generates a key by combining the environment name, a predefined prefix, and the provided identifier.
// It returns the constructed key or an error if any issues occur during the process.
//
//nolint:unparam
func (acc AdaptiveConfigurationControl) getKey(orgOrWorkspaceOrConfigGroupID string) (string, error) {
	rootConfig := ResolvedConfiguration

	// local-1.cfg.organization123.platform.configuration.v2alpha.Configuration
	key := rootConfig.App.EnvironmentName + configPrefix + orgOrWorkspaceOrConfigGroupID

	return key, nil
}
