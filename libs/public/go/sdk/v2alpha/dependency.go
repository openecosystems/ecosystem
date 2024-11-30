package sdkv2alphalib

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	GithubRegistryName     = "github.com/jeannotcompany"
	GithubRegistryBasePath = GithubRegistryName + "/ecosystem"
	PathRegistryName       = "path/local"
	DependencyFileExt      = ".binpb"
)

type Dependency struct {
	registry DependencyRegistryProvider
	data     []byte
}

type DependencyRegistryProvider interface {
	GetDependency() (*Dependency, error)
	Name() string
}

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

type GitHubDependencyProvider struct {
	settings   *specv2pb.SpecSystem
	fileSystem *FileSystem
}

func newGitHubDependencyProvider(settings *specv2pb.SpecSystem) *GitHubDependencyProvider {
	return &GitHubDependencyProvider{
		settings:   settings,
		fileSystem: NewFileSystem(),
	}
}

func (provider *GitHubDependencyProvider) Name() string { return GithubRegistryName }

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
	resp, err := http.Get("https://" + GithubRegistryBasePath + "/" + provider.settings.Name + DependencyFileExt)
	if err != nil {
		return nil, fmt.Errorf("failed to download: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

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

type PathDependencyProvider struct {
	settings   *specv2pb.SpecSystem
	fileSystem *FileSystem
}

func newPathDependencyProvider(settings *specv2pb.SpecSystem) *PathDependencyProvider {
	return &PathDependencyProvider{
		settings:   settings,
		fileSystem: NewFileSystem(),
	}
}

func (provider *PathDependencyProvider) Name() string { return PathRegistryName }

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

// generateIdentifier creates a unique identifier based on the given input.
// The identifier is an SHA-256 hash truncated to 16 characters.
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
