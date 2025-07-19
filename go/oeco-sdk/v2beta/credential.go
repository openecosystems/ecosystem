package sdkv2betalib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	typev2pb "github.com/openecosystems/ecosystem/go/oeco-sdk/v2beta/gen/platform/type/v2"

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

// CredentialProvider manages CLI-specific credential operations using a file system abstraction.
type CredentialProvider struct {
	fs *FileSystem
}

// NewCredentialProvider initializes and returns a new instance of CredentialProvider with a preconfigured file system.
func NewCredentialProvider() (*CredentialProvider, error) {
	_fs := NewFileSystem()

	return &CredentialProvider{
		fs: _fs,
	}, nil
}

// GetCredential retrieves a credential for the specified ecosystem and returns it along with any potential error.
func (p *CredentialProvider) GetCredential(t typev2pb.CredentialType, override string) (*typev2pb.Credential, error) {
	_ecosystem := override
	if _ecosystem == "" {
		file, err := p.fs.ReadFile(DefaultContextFile)
		if err != nil {
			return nil, fmt.Errorf("CredentialProvider could not read config file: %w", err)
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
func (p *CredentialProvider) SaveCredential(credential *typev2pb.Credential) error {
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
