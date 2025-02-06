package sdkv2alphalib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"

	"google.golang.org/protobuf/encoding/protojson"
)

// CredentialExt represents the file extension used for storing credential data in JSON format.
const (
	CredentialExt = "-credential.json" //nolint:gosec
)

// SpecCredentialProvider defines methods for retrieving and saving credentials specific to ecosystems.
// GetCredential retrieves a SpecSettings instance for a given ecosystem.
// SaveCredential saves a provided Credential and returns an error if the operation fails.
type SpecCredentialProvider interface {
	GetCredential(t typev2pb.CredentialType, override string) (*typev2pb.Credential, error)
	SaveCredential(credential *typev2pb.Credential) error
}

// CLICredentialProvider manages CLI-specific credential operations using a file system abstraction.
type CLICredentialProvider struct {
	fs *FileSystem
}

// NewCLICredentialProvider initializes and returns a new instance of CLICredentialProvider with a preconfigured file system.
func NewCLICredentialProvider() (*CLICredentialProvider, error) {
	_fs := NewFileSystem()

	return &CLICredentialProvider{
		fs: _fs,
	}, nil
}

// GetCredential retrieves a credential for the specified ecosystem and returns it along with any potential error.
func (p *CLICredentialProvider) GetCredential(t typev2pb.CredentialType, override string) (*typev2pb.Credential, error) {
	_ecosystem := override
	if _ecosystem == "" {
		file, err := p.fs.ReadFile(DefaultContextFile)
		if err != nil {
			return nil, fmt.Errorf("CLICredentialProvider could not read config file: %w", err)
		}
		_ecosystem = strings.TrimSpace(string(file))
	}

	switch t {
	case typev2pb.CredentialType_CREDENTIAL_TYPE_ACCOUNT_AUTHORITY:
		_ecosystem = "aa-" + _ecosystem
	}

	_credential, err := p.fs.ReadFile(filepath.Join(p.fs.CredentialsDirectory, _ecosystem+CredentialExt))
	if err != nil {
		return nil, err
	}

	var credential typev2pb.Credential
	err = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}.Unmarshal(_credential, &credential)
	if err != nil {
		return nil, err
	}

	return &credential, nil
}

// SaveCredential saves the provided Credential object to a persistent storage medium and returns an error if it fails.
func (p *CLICredentialProvider) SaveCredential(credential *typev2pb.Credential) error {
	_credential, err := protojson.MarshalOptions{
		Multiline: true,
		Indent:    "  ",
	}.Marshal(credential)
	if err != nil {
		return err
	}

	name := credential.EcosystemSlug
	switch credential.Type {
	case typev2pb.CredentialType_CREDENTIAL_TYPE_MESH_ACCOUNT:
		name = credential.EcosystemSlug
	case typev2pb.CredentialType_CREDENTIAL_TYPE_ACCOUNT_AUTHORITY:
		name = "aa-" + credential.EcosystemSlug
	}

	err2 := p.fs.WriteFile(filepath.Join(p.fs.CredentialsDirectory, name+CredentialExt), _credential, os.ModePerm)
	if err2 != nil {
		fmt.Println(err2)
		return err2
	}

	return nil
}
