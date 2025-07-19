package sdkv2betalib

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	specv2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/spec/v2"
)

// GithubRegistryName defines the base name of the GitHub registry for the company.
// GithubRegistryBasePath specifies the base path for the GitHub ecosystem registry.
// PathRegistryName defines the local path name used for registry operations.
// DependencyFileExt specifies the file extension for dependency binary protocol buffers.
const (
	GithubRegistryName     = "github.com/jeannotcompany"
	GithubRegistryBasePath = GithubRegistryName + "/ecosystem"
	PathRegistryName       = "path/local"
	DependencyFileExt      = ".binpb"
)

// Dependency represents a structure containing a reference to a registry provider and associated dependency data.
type Dependency struct {
	registry DependencyRegistryProvider
	data     []byte
}

// DependencyRegistryProvider is an interface to provide access to dependency registry details and metadata.
// GetDependency retrieves the associated Dependency instance or returns an error if unavailable.
// Name returns the name or identifier of the dependency registry.
type DependencyRegistryProvider interface {
	GetDependency() (*Dependency, error)
	Name() string
}

// NewDynamicDependencyProvider creates a Dependency based on the specified SpecSystem configuration.
// It determines the appropriate DependencyRegistryProvider to use (e.g., GitHub, Path) or logs an error
// if a custom registry type is unsupported.
func NewDynamicDependencyProvider(s *specv2pb.SpecSystem) *Dependency {
	dependency := Dependency{}
	if s.Registry == nil {
		dependency = Dependency{
			registry: newGitHubDependencyProvider(s),
		}
	} else if s.Registry.Git != nil {
		dependency = Dependency{
			registry: newGitHubDependencyProvider(s),
		}
	} else if s.Registry.Path != "" {
		dependency = Dependency{
			registry: newPathDependencyProvider(s),
		}
	} else if s.Registry.Registry != "" {
		fmt.Println("ERROR: Custom registries are not yet supported")
	}

	return &dependency
}

// GitHubDependencyProvider provides dependencies from GitHub using a specified system configuration and file system.
type GitHubDependencyProvider struct {
	settings   *specv2pb.SpecSystem
	fileSystem *FileSystem
}

// newGitHubDependencyProvider creates a new GitHubDependencyProvider instance with the given settings and initializes a file system.
func newGitHubDependencyProvider(settings *specv2pb.SpecSystem) *GitHubDependencyProvider {
	return &GitHubDependencyProvider{
		settings:   settings,
		fileSystem: NewFileSystem(),
	}
}

// Name returns the constant name of the GitHub dependency provider, used to identify the provider in the registry system.
func (provider *GitHubDependencyProvider) Name() string { return GithubRegistryName }

// GetDependency retrieves a dependency data from the cache or downloads it, storing it in the file system if necessary.
func (provider *GitHubDependencyProvider) GetDependency() (*Dependency, error) {
	baseDest, dest := generatePath(provider.settings.Name, provider)
	err := provider.fileSystem.UnderlyingFileSystem.MkdirAll(baseDest, os.ModePerm)
	if err != nil {
		fmt.Println("github dependency error: ", err)
	}

	exists, _ := provider.fileSystem.Exists(dest)
	if exists {
		fmt.Println("Dependency already in cache:", dest)
		data, err := provider.fileSystem.ReadFile(dest)
		if err != nil {
			return nil, err
		}

		return &Dependency{
			registry: provider,
			data:     data,
		}, nil
	}

	// Download the file
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	url := "https://" + GithubRegistryBasePath + "/" + provider.settings.Name + DependencyFileExt
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	// Use the default HTTP client to execute the request
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to download: %w", err)
	}
	defer resp.Body.Close() //nolint:errcheck

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("github dependency read error: ", err)
	}

	// Save to destination
	err = provider.fileSystem.WriteFile(dest, data, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	return &Dependency{
		registry: provider,
		data:     data,
	}, nil
}

// PathDependencyProvider manages dependencies by integrating specification settings with file system operations.
type PathDependencyProvider struct {
	settings   *specv2pb.SpecSystem
	fileSystem *FileSystem
}

// newPathDependencyProvider creates and initializes a PathDependencyProvider using the given SpecSystem settings.
func newPathDependencyProvider(settings *specv2pb.SpecSystem) *PathDependencyProvider {
	return &PathDependencyProvider{
		settings:   settings,
		fileSystem: NewFileSystem(),
	}
}

// Name returns the name of the registry as a string, which is a constant value defined as PathRegistryName.
func (provider *PathDependencyProvider) Name() string { return PathRegistryName }

// GetDependency retrieves a dependency from the cache or creates it if not present and returns it as a Dependency instance.
func (provider *PathDependencyProvider) GetDependency() (*Dependency, error) {
	src := provider.settings.Registry.Path
	baseDest, dest := generatePath(provider.settings.Name, provider)
	err := provider.fileSystem.UnderlyingFileSystem.MkdirAll(baseDest, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	exists, _ := provider.fileSystem.Exists(dest)
	if exists {
		fmt.Println("Dependency already in cache:", dest)
		data, err := provider.fileSystem.ReadFile(dest)
		if err != nil {
			return nil, err
		}

		return &Dependency{
			registry: provider,
			data:     data,
		}, nil
	}

	err = provider.fileSystem.CopyFile(filepath.Join(src, provider.settings.Name+DependencyFileExt), dest)
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}

	data, err := provider.fileSystem.ReadFile(dest)
	if err != nil {
		return nil, err
	}

	return &Dependency{
		registry: provider,
		data:     data,
	}, nil
}

// generatePath generates a unique cache path and file path for a dependency based on the registry provider and system.
func generatePath(system string, provider DependencyRegistryProvider) (string, string) {
	// Generate SHA-256 hash of the input and encode the hash to a hex string and truncate to 16 characters
	_url := provider.Name()
	u := strings.Split(_url, "/")
	url := u[0]
	hash := sha256.Sum256([]byte(url))
	h := hex.EncodeToString(hash[:])[:16]
	basePath := filepath.Join(RegistryCacheDirectory, url+"-"+h, u[1])
	path := filepath.Join(basePath, system+DependencyFileExt)

	return basePath, path
}
