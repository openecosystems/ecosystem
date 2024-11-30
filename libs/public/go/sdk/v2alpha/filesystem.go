package sdkv2alphalib

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/afero"
)

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
}

const (
	HomeDirectoryName          = ".config/oeco"
	ConfigurationName          = "config"
	ConfigurationExtension     = "yaml"
	LogDirectoryName           = "logs"
	TmpDirectoryName           = "tmp"
	ContextDirectoryName       = "context"
	DefaultContextFileName     = "default"
	CredentialDirectoryName    = "credentials"
	RegistryDirectoryName      = "registry"
	RegistryCacheDirectoryName = "cache"
	ConfigurationDirectoryName = "configuration"
)

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
	RegistryCacheDirectory = filepath.Join(RegistryDirectory, RegistryCacheDirectoryName)
	DefaultContextFile     = filepath.Join(ContextDirectory, DefaultContextFileName)
	Filesystem             *FileSystem
)

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
		fmt.Println("config directory error: ", err)
	}

	if err := filesystem.CreateFile(DefaultContextFile); err != nil {
		fmt.Println("default context file error: ", err)
	}

	Filesystem = filesystem

	return filesystem
}

func getUserHomeDirectory() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		// Panic here... should not happen
	}

	return home
}

func getHomeDirectory() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		// Panic here... should not happen
	}

	return filepath.Join(home, HomeDirectoryName)
}

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

func (filesystem *FileSystem) Exists(file string) (bool, error) {
	fs := filesystem.UnderlyingFileSystem
	return afero.Exists(fs, file)
}

func (filesystem *FileSystem) DirExists(directory string) (bool, error) {
	fs := filesystem.UnderlyingFileSystem
	return afero.DirExists(fs, directory)
}

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

func (filesystem *FileSystem) ReadFile(file string) ([]byte, error) {
	fs := filesystem.UnderlyingFileSystem
	bytes, err := afero.ReadFile(fs, file)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

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

func (filesystem *FileSystem) ReadDir(dir string) ([]os.FileInfo, error) {
	fs := filesystem.UnderlyingFileSystem

	fileInfos, err := afero.ReadDir(fs, dir)
	if err != nil {
		return nil, err
	}

	return fileInfos, nil
}

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
