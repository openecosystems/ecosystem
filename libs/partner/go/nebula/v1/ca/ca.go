package nebulav1ca

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"connectrpc.com/connect"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	ecosystemv2alphapb "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
	iamv2alphapb "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

const (
	_ca            = "ca"
	_host          = "host"
	_signed        = "signed"
	crText         = ".crt"
	keyExt         = ".key"
	pngExt         = ".png"
	caCertName     = _ca + crText
	caKeyName      = _ca + keyExt
	caQrName       = _ca + pngExt
	hostCertName   = _host + crText
	hostKeyName    = _host + keyExt
	signedCertName = _signed + crText
	signedQrName   = _signed + pngExt
)

// AccountAuthorityKeyType represents an enumeration for different types of account authority keys.
type AccountAuthorityKeyType int

// Cert represents a certificate-based account authority key type.
// Key represents a key-based account authority key type.
const (
	Cert AccountAuthorityKeyType = iota
	Key
)

// Binding represents an entity responsible for managing Nebula certificate binary and its file path.
type Binding struct {
	NebulaCertBinaryFile *os.File
	NebulaCertBinaryPath string

	aac *AccountAuthorityCache
	cp  *sdkv2alphalib.CredentialProvider

	// options *caOptions
}

// Bound is a global Binding instance that holds configuration and state for the Certificate Authority binding.
// BindingName is the constant string representing the name of the Certificate Authority binding.
var (
	Bound       *Binding
	BindingName = "CA_BINDING"
)

// Name returns the name of the binding as a string. It is used to identify the binding in the registered bindings map.
func (b *Binding) Name() string {
	return BindingName
}

// Validate checks the validity and correctness of the binding within the provided context and bindings.
func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {
	return nil
}

// Bind adds the current Binding instance to the provided bindings map, initializing it if not already bound.
func (b *Binding) Bind(_ context.Context, bindings *sdkv2alphalib.Bindings) *sdkv2alphalib.Bindings {
	if Bound == nil {
		var once sync.Once
		once.Do(
			func() {
				binaryData, err := embeddedFiles.ReadFile(binaryPath)
				if err != nil {
					log.Fatalf("Failed to read embedded binary: %v", err)
				}

				// Write the binary to a temporary file
				nebulaCertBinary, err := os.CreateTemp("", "oeco-bin-*")
				if err != nil {
					log.Fatalf("Failed to create temp file: %v", err)
				}

				b.NebulaCertBinaryFile = nebulaCertBinary

				// Write binary data to the temp file
				if _, err := nebulaCertBinary.Write(binaryData); err != nil {
					log.Fatalf("Failed to write to temp file: %v", err)
				}

				nebulaCertBinaryPath := nebulaCertBinary.Name()

				// Make the temp file executable
				if err := os.Chmod(nebulaCertBinary.Name(), 0o755); err != nil {
					log.Fatalf("Failed to make temp file executable: %v", err)
				}

				b.NebulaCertBinaryPath = nebulaCertBinaryPath

				provider, err := sdkv2alphalib.NewCredentialProvider()
				if err != nil {
					log.Fatalf("nebula ca: failed to load credential provider: %v", err)
				}

				Bound = &Binding{
					NebulaCertBinaryFile: nebulaCertBinary,
					NebulaCertBinaryPath: nebulaCertBinaryPath,

					aac: NewAccountAuthorityCache(),
					cp:  provider,
				}
				bindings.Registered[b.Name()] = Bound
			})
	} else {
		bindings.Registered[b.Name()] = Bound
		fmt.Println("CA already bound")
	}

	return bindings
}

// GetBinding returns the globally initialized instance of the Binding. If no instance exists, it returns nil.
func (b *Binding) GetBinding() interface{} {
	return Bound
}

// Close removes the temporary Nebula Certificate Binary file and returns any errors encountered during the removal.
func (b *Binding) Close() error {
	defer func(name string) {
		_ = os.Remove(name)
	}(b.NebulaCertBinaryFile.Name()) // Clean up after execution

	b.aac.Close()

	return nil
}

// CredStore is a structure used to store credential data and its associated file path. It manages credentials securely.
type CredStore struct {
	path string
	cred *typev2pb.Credential
}

// AccountAuthorityCache using sync.Map for concurrent access
type AccountAuthorityCache struct {
	cache sync.Map
}

// NewAccountAuthorityCache initializes and returns a new instance of AccountAuthorityCache
func NewAccountAuthorityCache() *AccountAuthorityCache {
	return &AccountAuthorityCache{}
}

// Get retrieves a value if it exists, otherwise stores a new value
func (aac *AccountAuthorityCache) Get(keyType AccountAuthorityKeyType, cp *sdkv2alphalib.CredentialProvider, ecosystem string) (*CredStore, bool, error) {
	// Check if key exists

	key := ecosystem
	switch keyType {
	case Cert:
		key = "cert-" + ecosystem
	case Key:
		key = "key-" + ecosystem
	}

	if existingPath, exists := aac.cache.Load(key); exists {
		return existingPath.(*CredStore), true, nil
	}

	tempFile, err := os.CreateTemp(sdkv2alphalib.CredentialDirectory, ".aa-"+key+"-*")
	if err != nil {
		return nil, false, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("ca: cannot create temp directory"), err)
	}

	cred, err := cp.GetCredential(typev2pb.CredentialType_CREDENTIAL_TYPE_ACCOUNT_AUTHORITY, ecosystem)
	if err != nil {
		return nil, false, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("nebula ca: failed to get credential"), err)
	}

	path := tempFile.Name()
	switch keyType {
	case Cert:
		err = os.WriteFile(path, []byte(cred.AaCertX509), 0o600)
	case Key:
		err = os.WriteFile(path, []byte(cred.AaPrivateKey), 0o600)
	}
	if err != nil {
		return nil, false, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(fmt.Errorf("ca: Error writing file %s: ", path), err)
	}

	store := &CredStore{
		path: path,
		cred: cred,
	}

	// Store the new value
	aac.cache.Store(key, store)
	return store, false, nil
}

// Close loops through the cache, deletes files from the filesystem, and clears entries
func (aac *AccountAuthorityCache) Close() {
	aac.cache.Range(func(key, value interface{}) bool {
		store := value.(*CredStore) // Ensure value is a string (file path)
		path := store.path

		// Attempt to remove the file
		err := os.Remove(path)
		if err != nil {
			fmt.Printf("Error deleting file %s: %v\n", path, err)
		} else {
			fmt.Printf("Deleted file: %s\n", path)
		}

		// Remove the entry from the cache
		aac.cache.Delete(key)

		return true
	})
}

// GetAccountAuthority creates a new Certificate Authority using the specified request parameters.
// Returns the created Certificate Authority or an error if the operation fails.
func (b *Binding) GetAccountAuthority(_ context.Context, req *iamv2alphapb.CreateAccountAuthorityRequest) (*iamv2alphapb.AccountAuthority, error) {
	nca := b.NebulaCertBinaryPath

	if req.Name == "" {
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("required parameter name is missing"))
	}

	var c string
	var curve typev2pb.Curve
	switch req.Curve {
	case typev2pb.Curve_CURVE_ECDSA:
		c = "P256"
		curve = typev2pb.Curve_CURVE_ECDSA
	case typev2pb.Curve_CURVE_EDDSA:
		c = "25519"
		curve = typev2pb.Curve_CURVE_EDDSA
	case typev2pb.Curve_CURVE_UNSPECIFIED:
		c = "25519"
		curve = typev2pb.Curve_CURVE_EDDSA
	default:
		c = "25519"
		curve = typev2pb.Curve_CURVE_EDDSA
	}

	tempDir, err := os.MkdirTemp("", "oeco-aa-*")
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to create temp file"), err)
	}

	defer os.RemoveAll(tempDir) //nolint:errcheck

	certpath := filepath.Join(tempDir, caCertName)
	keypath := filepath.Join(tempDir, caKeyName)
	qpath := filepath.Join(tempDir, caQrName)

	nebula := exec.Command(nca, "ca",
		"-name", req.Name,
		"-curve", c,
		"-out-crt", certpath,
		"-out-key", keypath,
		"-out-qr", qpath,
	)

	nebula.Stdout = os.Stdout
	nebula.Stderr = os.Stderr

	if err2 := nebula.Run(); err2 != nil {
		return nil, ErrFailedToRunCommand.WithInternalErrorDetail(errors.New("failed to execute binary"), err2)
	}

	id := ksuid.New()
	now := timestamppb.Now()

	cfile, err3 := getFile(caCertName, crText, certpath, now)
	if err3 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err3)
	}

	kfile, err4 := getFile(caKeyName, keyExt, keypath, now)
	if err4 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err4)
	}

	qfile, err5 := getFile(caQrName, pngExt, qpath, now)
	if err5 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err5)
	}

	return &iamv2alphapb.AccountAuthority{
		Id:        id.String(),
		CreatedAt: now,
		UpdatedAt: now,
		Credential: &typev2pb.Credential{
			// TODO Fix this. Ensure Validated
			EcosystemSlug: req.Name,
			Type:          typev2pb.CredentialType_CREDENTIAL_TYPE_ACCOUNT_AUTHORITY,
			Curve:         curve,
			// Duration:  req.Duration,
			AaCertX509:       string(cfile.GetContent()),
			AaCertX509QrCode: base64.StdEncoding.EncodeToString(qfile.GetContent()),
			AaPrivateKey:     string(kfile.GetContent()),
		},
	}, nil
}

// GetPKI generates a public-private key pair based on the specified curve in the request and returns the files and possible error.
func (b *Binding) GetPKI(_ context.Context, req *iamv2alphapb.CreateAccountRequest) (cert *typev2pb.File, key *typev2pb.File, err error) {
	nca := b.NebulaCertBinaryPath

	var c string
	switch req.Curve {
	case typev2pb.Curve_CURVE_ECDSA:
		c = "P256"
	case typev2pb.Curve_CURVE_EDDSA:
		c = "25519"
	case typev2pb.Curve_CURVE_UNSPECIFIED:
		c = "25519"
	default:
		c = "25519"
	}

	tempDir, err := os.MkdirTemp("", "oeco-pki-*")
	if err != nil {
		return nil, nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to create temp file"), err)
	}
	defer os.RemoveAll(tempDir) //nolint:errcheck

	hostcertpath := filepath.Join(tempDir, hostCertName)
	hostkeypath := filepath.Join(tempDir, hostKeyName)

	nebula := exec.Command(nca, "keygen",
		"-curve", c,
		"-out-pub", hostcertpath,
		"-out-key", hostkeypath,
	)

	nebula.Stdout = os.Stdout
	nebula.Stderr = os.Stderr

	if err = nebula.Run(); err != nil {
		return nil, nil, ErrFailedToRunCommand.WithInternalErrorDetail(errors.New("failed to execute binary"), err)
	}

	now := timestamppb.Now()

	cfile, err3 := getFile(hostCertName, crText, hostcertpath, now)
	if err3 != nil {
		return nil, nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err3)
	}

	kfile, err4 := getFile(hostKeyName, keyExt, hostkeypath, now)
	if err4 != nil {
		return nil, nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err4)
	}

	return cfile, kfile, nil
}

// SignCert generates and returns a signed certificate and its QR code using Nebula certificate binary and CA information.
func (b *Binding) SignCert(_ context.Context, req *iamv2alphapb.SignAccountRequest, opts ...CAOption) (*typev2pb.Credential, error) {
	options, _ := newCAOptions(opts)

	nca := b.NebulaCertBinaryPath

	tempDir, err := os.MkdirTemp("", "oeco-s-*")
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to create temp file"), err)
	}
	defer os.RemoveAll(tempDir) //nolint:errcheck

	signedcertpath := filepath.Join(tempDir, signedCertName)
	signedqrpath := filepath.Join(tempDir, signedQrName)

	caCert, _, err := b.aac.Get(Cert, b.cp, req.Name)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err)
	}
	caKey, _, err := b.aac.Get(Key, b.cp, req.Name)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err)
	}

	unsignedCertPath, err := saveFileTemporarily(tempDir, req.PublicCert)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to save unsigned cert file temporarily"), err)
	}

	ip, ipCIDR, err := b.getAnAvailableIPAddress(req, options)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to get an available IP address"), err)
	}

	hostname, err := b.getAvailableHostname(req, options)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to parse a hostname"), err)
	}

	groups := b.getAvailableGroups(req, options)

	nebula := exec.Command(nca, "sign",
		"-ca-crt", caCert.path,
		"-ca-key", caKey.path,
		"-in-pub", unsignedCertPath,
		"-ip", ipCIDR,
		"-groups", strings.Join(groups, ","),
		"-name", hostname,
		"-out-crt", signedcertpath,
		"-out-qr", signedqrpath,
	)

	nebula.Stdout = os.Stdout
	nebula.Stderr = os.Stderr

	if err = nebula.Run(); err != nil {
		return nil, ErrFailedToRunCommand.WithInternalErrorDetail(errors.New("failed to execute binary"), err)
	}

	now := timestamppb.Now()

	cfile, err := getFile(signedCertName, crText, signedcertpath, now)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("error getting cert file"), err)
	}

	qrfile, err := getFile(signedQrName, pngExt, signedqrpath, now)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("error getting qr code cert file"), err)
	}

	return &typev2pb.Credential{
		Type:           typev2pb.CredentialType_CREDENTIAL_TYPE_MESH_ACCOUNT,
		MeshAccountId:  "",
		EcosystemSlug:  req.Name,
		MeshHostname:   hostname,
		MeshIp:         ip,
		AaCertX509:     caCert.cred.AaCertX509,
		CertX509:       string(cfile.GetContent()),
		CertX509QrCode: base64.StdEncoding.EncodeToString(qrfile.GetContent()),
		Groups:         groups,
		Subnets:        nil,
	}, nil
}

func (b *Binding) getAvailableGroups(req *iamv2alphapb.SignAccountRequest, _ *caOptions) []string {
	switch req.EcosystemPeerType {
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_EDGE:
		return []string{"edge", "host"}
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_ECOSYSTEM_MULTIPLEXER:
		return []string{"multiplexer", "host"}
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_CONNECTOR:
		return []string{"connector", "host"}
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_HOST:
		return []string{"service", "host"}
	default:
		return []string{"host"}
	}
}

func (b *Binding) getAvailableHostname(req *iamv2alphapb.SignAccountRequest, _ *caOptions) (string, error) {
	switch req.EcosystemPeerType {
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_EDGE:
		return fmt.Sprintf("edge.%s.mesh", req.Name), nil
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_ECOSYSTEM_MULTIPLEXER:
		return fmt.Sprintf("api.%s.mesh", req.Name), nil
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_SERVICE_ACCOUNT:
		return fmt.Sprintf("api.%s.mesh", req.Name), nil
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_CONNECTOR:
		return fmt.Sprintf("api.%s.mesh", req.Name), nil
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_HOST:
		return fmt.Sprintf("api.%s.mesh", req.Name), nil
	default:
		return "", errors.New("unknown ecosystem peer type")
	}
}

func (b *Binding) getAnAvailableIPAddress(req *iamv2alphapb.SignAccountRequest, options *caOptions) (string, string, error) {
	switch req.EcosystemPeerType {
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_EDGE:

		if options.CIDR == nil {
			return "", "", errors.New("cidr is required; please specify it in the CA option: WithCIDR()")
		}

		// TODO: Support Multiple Edges
		ip, ipCIDR, err := options.CIDR.GetNthIP(1)
		if err != nil {
			return "", "", err
		}

		return ip, ipCIDR, nil
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_ECOSYSTEM_MULTIPLEXER:
		if options.CIDR == nil {
			return "", "", errors.New("cidr is required; please specify it in the CA option: WithCIDR()")
		}

		// TODO: Support Multiple Multiplexers
		ip, ipCIDR, err := options.CIDR.GetNthIP(5)
		if err != nil {
			return "", "", err
		}

		return ip, ipCIDR, nil
	case ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_SERVICE_ACCOUNT:
		if options.CIDR == nil {
			return "", "", errors.New("cidr is required; please specify it in the CA option: WithCIDR()")
		}

		// TODO: Support Multiple Multiplexers
		ip, ipCIDR, err := options.CIDR.GetNthIP(5)
		if err != nil {
			return "", "", err
		}

		return ip, ipCIDR, nil
	default:

		// TODO: Connect to Ecosystem and determine the next free IP
		return "192.168.100.20", "192.168.100.20/24", nil
	}
}

// getFile reads a file and constructs a File object with its metadata and content.
// Parameters:
// - name: The name of the file.
// - extensionWithPeriod: The file's extension, including the leading period (e.g., ".png").
// - path: The full path to the file.
// - time: Timestamp for creation and modification times.
// Returns:
// - A pointer to a File object containing the file details.
// - An error if the file cannot be read or MIME type cannot be determined.
func getFile(name string, extensionWithPeriod string, path string, time *timestamppb.Timestamp) (*typev2pb.File, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to read file: "+path), err)
	}

	return &typev2pb.File{
		Name:             name,
		Content:          data,
		Size:             int64(len(data)),
		Type:             sdkv2alphalib.MimeTypeRegistry.GetMimeType(extensionWithPeriod),
		CreationTime:     time.GetSeconds(),
		ModificationTime: time.GetSeconds(),
	}, nil
}

func saveFileTemporarily(tempFolder string, file *typev2pb.File) (string, error) {
	path := filepath.Join(tempFolder, file.GetName())
	// err := os.WriteFile(path, sanitizeCertificateInput(file.GetContent()), 0o600)
	err := os.WriteFile(path, file.GetContent(), 0o600)
	if err != nil {
		return "", sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to save cert temporarily in : "+tempFolder), err)
	}

	return path, nil
}

//nolint:unused
func sanitizeCertificateInput(input []byte) []byte {
	// Remove dangerous shell characters
	dangerousChars := []string{";", "&", "|", "`", "$(", ")", ">", "<"}
	for _, char := range dangerousChars {
		input = []byte(strings.ReplaceAll(string(input), char, ""))
	}
	return input
}

// A CAOption configures the certificate authority
type CAOption interface {
	apply(*caOptions)
}

type caOptions struct {
	CIDR              *sdkv2alphalib.CIDRBlock
	Spec              *specv2pb.Spec
	EcosystemPeerType ecosystemv2alphapb.EcosystemPeerType
}

type optionFunc func(*caOptions)

func (f optionFunc) apply(cfg *caOptions) { f(cfg) }

//nolint:unparam
func newCAOptions(options []CAOption) (*caOptions, *connect.Error) {
	// Defaults
	config := caOptions{}

	for _, opt := range options {
		opt.apply(&config)
	}

	if err := config.validate(); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *caOptions) validate() *connect.Error {
	return nil
}

// WithOptions composes multiple Options into one.
func WithOptions(opts ...CAOption) CAOption {
	return optionFunc(func(cfg *caOptions) {
		for _, opt := range opts {
			opt.apply(cfg)
		}
	})
}

// WithSpec sets the Spec field in the caOptions struct using the provided specv2pb.Spec instance.
func WithSpec(spec *specv2pb.Spec) CAOption {
	return optionFunc(func(cfg *caOptions) {
		cfg.Spec = spec
	})
}

// WithCIDR sets the CIDR block for the certificate authority configuration.
func WithCIDR(cidr string) CAOption {
	return optionFunc(func(cfg *caOptions) {
		c, err := sdkv2alphalib.NewCIDR(cidr)
		if err != nil {
			return
		}

		cfg.CIDR = c
	})
}

// WithEcosystemPeerType sets the EcosystemPeerType field in the caOptions configuration.
func WithEcosystemPeerType(ecosystemPeerType ecosystemv2alphapb.EcosystemPeerType) CAOption {
	return optionFunc(func(cfg *caOptions) {
		cfg.EcosystemPeerType = ecosystemPeerType
	})
}
