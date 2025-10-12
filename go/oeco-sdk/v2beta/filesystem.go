package sdkv2betalib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
)

// FileSystem represents a file system abstraction for managing directories and files.
// It includes paths for home, log, temporary, context, credentials, registry, and configuration directories.
type FileSystem struct {
	UnderlyingFileSystem afero.Fs
	HomeDirectory        string
	// DefaultConfigFile    string
	LogDirectory           string
	TmpDirectory           string
	ContextDirectory       string
	CredentialsDirectory   string
	RegistryDirectory      string
	RegistryCacheDirectory string
	ConfigurationDirectory string
	ConnectorDirectory     string
}

// HomeDirectoryName defines the name of the home configuration directory.
// ConfigurationName specifies the name of the main configuration file.
// ConfigurationExtension defines the file extension for configuration files.
// LogDirectoryName specifies the directory name for log files.
// TmpDirectoryName defines the temporary directory name.
// ContextDirectoryName specifies the directory name for context-related files.
// DefaultContextFileName specifies the default context file name.
// CredentialDirectoryName defines the directory name for storing credentials.
// RegistryDirectoryName specifies the directory name for registry-related data.
// RegistryCacheDirectoryName defines the directory name for registry cache data.
// ConfigurationDirectoryName specifies the directory name for configuration storage.
// ConnectorDirectoryName specifies the directory name for connector storage.
const (
	HomeDirectoryName          = ".config/adino"
	ConfigurationName          = "config"
	ConfigurationExtension     = "yaml"
	LogDirectoryName           = "logs"
	LogExtension               = "log"
	TmpDirectoryName           = "tmp"
	ContextDirectoryName       = "context"
	AdinoContextFileName       = "adino"
	DefaultContextFileName     = "default"
	CredentialDirectoryName    = "credentials"
	RegistryDirectoryName      = "registry"
	RegistryCacheDirectoryName = "cache"
	ConfigurationDirectoryName = "configuration"
	ConnectorDirectoryName     = "connector"
)

// HomeDirectory defines the default home directory path using getHomeDirectory function.
// UserHomeDirectory defines the user's home directory path using getUserHomeDirectory function.
// ConfigFile defines the full path to the configuration file in the home directory.
// LogDirectory defines the path to the log directory in the home directory.
// TmpDirectory defines the path to the temporary directory in the home directory.
// ContextDirectory defines the path to the context directory in the home directory.
// CredentialDirectory defines the path to the credential directory in the home directory.
// ConnectorDirectory defines the path to the connector directory in the home directory.
// RegistryDirectory defines the path to the registry directory in the home directory.
// ConfigurationDirectory defines the path to the configuration directory in the home directory.
// RegistryCacheDirectory defines the path to the registry cache directory in the registry directory.
// DefaultContextFile defines the full path to the default context file in the context directory.
// Filesystem is a pointer to the FileSystem struct managing file operations.
var (
	HomeDirectory          = getHomeDirectory()
	UserHomeDirectory      = getUserHomeDirectory()
	ConfigFile             = filepath.Join(HomeDirectory, ConfigurationName+"."+ConfigurationExtension)
	LogDirectory           = filepath.Join(HomeDirectory, LogDirectoryName)
	TmpDirectory           = filepath.Join(HomeDirectory, TmpDirectoryName)
	ContextDirectory       = filepath.Join(HomeDirectory, ContextDirectoryName)
	CredentialDirectory    = filepath.Join(HomeDirectory, CredentialDirectoryName)
	RegistryDirectory      = filepath.Join(HomeDirectory, RegistryDirectoryName)
	ConfigurationDirectory = filepath.Join(HomeDirectory, ConfigurationDirectoryName)
	ConnectorDirectory     = filepath.Join(HomeDirectory, ConnectorDirectoryName)
	RegistryCacheDirectory = filepath.Join(RegistryDirectory, RegistryCacheDirectoryName)
	AdinoContextFile       = filepath.Join(ContextDirectory, AdinoContextFileName)
	DefaultContextFile     = filepath.Join(ContextDirectory, DefaultContextFileName)
	OecoLogFile            = filepath.Join(LogDirectory, AdinoContextFileName+"."+LogExtension)
	Filesystem             *FileSystem
)

// NewFileSystem initializes and returns a new instance of FileSystem with preconfigured directories and files.
func NewFileSystem() *FileSystem {
	fs := afero.NewOsFs()

	filesystem := &FileSystem{
		UnderlyingFileSystem: fs,
		HomeDirectory:        HomeDirectory,
		// DefaultConfigFile:    ConfigFile,
		LogDirectory:           LogDirectory,
		TmpDirectory:           TmpDirectory,
		ContextDirectory:       ContextDirectory,
		CredentialsDirectory:   CredentialDirectory,
		RegistryDirectory:      RegistryDirectory,
		RegistryCacheDirectory: RegistryCacheDirectory,
		ConfigurationDirectory: ConfigurationDirectory,
		ConnectorDirectory:     ConnectorDirectory,
	}

	if err := filesystem.CreateDirectory(HomeDirectory); err != nil {
		fmt.Println("create home directory error: ", err)
	}

	if err := filesystem.CreateDirectory(LogDirectory); err != nil {
		fmt.Println("log directory error: ", err)
	}

	if err := filesystem.CreateDirectory(TmpDirectory); err != nil {
		fmt.Println("tmp directory error: ", err)
	}

	if err := filesystem.CreateDirectory(ContextDirectory); err != nil {
		fmt.Println("context directory error: ", err)
	}

	if err := filesystem.CreateDirectory(CredentialDirectory); err != nil {
		fmt.Println("credential directory error: ", err)
	}

	if err := filesystem.CreateDirectory(RegistryDirectory); err != nil {
		fmt.Println("registry directory error", err)
	}

	if err := filesystem.CreateDirectory(RegistryCacheDirectory); err != nil {
		fmt.Println("registry cache directory error: ", err)
	}

	if err := filesystem.CreateDirectory(ConfigurationDirectory); err != nil {
		fmt.Println("configuration directory error: ", err)
	}

	if err := filesystem.CreateDirectory(ConnectorDirectory); err != nil {
		fmt.Println("connector directory error: ", err)
	}

	if err := filesystem.CreateFile(DefaultContextFile); err != nil {
		fmt.Println("default context file error: ", err)
	}

	Filesystem = filesystem

	return filesystem
}

// getUserHomeDirectory returns the current user's home directory as a string or exits the program on error.
func getUserHomeDirectory() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		// Panic here... should not happen
	}

	return home
}

// getHomeDirectory returns the path to the user's home configuration directory, constructing it with a predefined name.
// Exits the program with an error message if the home directory cannot be determined.
func getHomeDirectory() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		// Panic here... should not happen
	}

	return filepath.Join(home, HomeDirectoryName)
}

// CreateDirectory ensures the specified directory exists, creating it with permissions 0755 if it does not exist.
func (filesystem *FileSystem) CreateDirectory(directory string) error {
	fs := filesystem.UnderlyingFileSystem

	exists, err := afero.Exists(fs, directory)
	if err != nil {
		return err
	}

	if !exists {
		err := fs.MkdirAll(directory, 0o755)
		if err != nil {
			return err
		}
	}

	return nil
}

// CreateFile creates a new file with the specified name if it does not already exist in the filesystem.
// Returns an error if the file cannot be created or if an underlying issue occurs.
func (filesystem *FileSystem) CreateFile(file string) error {
	fs := filesystem.UnderlyingFileSystem
	exists, err := afero.Exists(fs, file)
	if err != nil {
		return err
	}

	if !exists {
		_, err = fs.Create(file)
		if err != nil {
			return err
		}
	}

	return nil
}

// Exists checks if a file exists at the specified path using the underlying file system and returns a boolean and an error.
func (filesystem *FileSystem) Exists(file string) (bool, error) {
	fs := filesystem.UnderlyingFileSystem
	return afero.Exists(fs, file)
}

// DirExists checks if the specified directory exists in the underlying file system and returns a boolean and an error.
func (filesystem *FileSystem) DirExists(directory string) (bool, error) {
	fs := filesystem.UnderlyingFileSystem
	return afero.DirExists(fs, directory)
}

// WriteFile writes the provided data to the specified file with the given permissions. Creates the file if it does not exist.
func (filesystem *FileSystem) WriteFile(file string, data []byte, perm os.FileMode) error {
	fs := filesystem.UnderlyingFileSystem

	err := filesystem.CreateFile(file)
	if err != nil {
		return err
	}

	err = afero.WriteFile(fs, file, data, perm)
	if err != nil {
		return err
	}

	return nil
}

// ReadFile reads the contents of the specified file and returns its content as a byte slice. Returns an error if reading fails.
func (filesystem *FileSystem) ReadFile(file string) ([]byte, error) {
	fs := filesystem.UnderlyingFileSystem
	bytes, err := afero.ReadFile(fs, file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

// DeleteFile removes the specified file if it exists in the underlying filesystem and returns an error if any issues occur.
func (filesystem *FileSystem) DeleteFile(file string) error {
	fs := filesystem.UnderlyingFileSystem
	exists, err := afero.Exists(fs, file)
	if err != nil {
		return err
	}

	if exists {
		err = fs.Remove(file)
		if err != nil {
			return err
		}
	}

	return nil
}

// CopyFile copies a file from the specified source path to the destination path within the file system.
func (filesystem *FileSystem) CopyFile(file string, destination string) error {
	fs := filesystem.UnderlyingFileSystem
	exists, err := afero.Exists(fs, file)
	if err != nil {
		return err
	}

	if exists {
		bytes, err := filesystem.ReadFile(file)
		if err != nil {
			return err
		}

		err = filesystem.WriteFile(destination, bytes, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

// ReadDir reads the contents of the specified directory and returns a slice of os.FileInfo or an error if it fails.
func (filesystem *FileSystem) ReadDir(dir string) ([]os.FileInfo, error) {
	fs := filesystem.UnderlyingFileSystem

	fileInfos, err := afero.ReadDir(fs, dir)
	if err != nil {
		return nil, err
	}

	return fileInfos, nil
}

// DeleteDirectory removes the specified directory and all its contents if it exists in the underlying file system.
func (filesystem *FileSystem) DeleteDirectory(directory string) error {
	fs := filesystem.UnderlyingFileSystem

	log.Debug(directory)

	exists, err := afero.Exists(fs, directory)
	if err != nil {
		return err
	}

	if exists {
		err := fs.RemoveAll(directory)
		if err != nil {
			return err
		}
	}

	return nil
}
