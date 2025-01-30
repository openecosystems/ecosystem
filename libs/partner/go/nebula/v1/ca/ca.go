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

	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	iamv2alphapb "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

const (
	_ca            = "ca"
	_host          = "host"
	_signed        = "signed"
	crtext         = ".crt"
	keyext         = ".key"
	pngext         = ".png"
	cacertname     = _ca + crtext
	cakeyname      = _ca + keyext
	caqrname       = _ca + pngext
	hostcertname   = _host + crtext
	hostkeyname    = _host + keyext
	signedcertname = _signed + crtext
	signedqrname   = _signed + pngext
)

// Binding represents an entity responsible for managing Nebula certificate binary and its file path.
type Binding struct {
	NebulaCertBinaryFile *os.File
	NebulaCertBinaryPath string
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
				nebulaCertBinary, err := os.CreateTemp("", "oeco-ca-*")
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

				Bound = &Binding{
					NebulaCertBinaryFile: nebulaCertBinary,
					NebulaCertBinaryPath: nebulaCertBinaryPath,
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
	return nil
}

// GetAccountAuthority creates a new Certificate Authority using the specified request parameters.
// Returns the created Certificate Authority or an error if the operation fails.
func (b *Binding) GetAccountAuthority(_ context.Context, req *iamv2alphapb.CreateAccountAuthorityRequest) (*iamv2alphapb.AccountAuthority, error) {
	nca := b.NebulaCertBinaryPath

	// TODO: This should be done in the initial validate
	if req.Name == "" {
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("required parameter name is missing"))
	}

	var c string
	var curve iamv2alphapb.Curve
	switch req.Curve {
	case iamv2alphapb.Curve_CURVE_ECDSA:
		c = "P256"
		curve = iamv2alphapb.Curve_CURVE_ECDSA
	case iamv2alphapb.Curve_CURVE_EDDSA:
		c = "25519"
		curve = iamv2alphapb.Curve_CURVE_EDDSA
	case iamv2alphapb.Curve_CURVE_UNSPECIFIED:
		c = "25519"
		curve = iamv2alphapb.Curve_CURVE_EDDSA
	default:
		c = "25519"
		curve = iamv2alphapb.Curve_CURVE_EDDSA
	}

	// Write the binary to a temporary file
	tempDir, err := os.MkdirTemp("", "oeco-ca-*")
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to create temp file"), err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempDir)

	certpath := filepath.Join(tempDir, cacertname)
	keypath := filepath.Join(tempDir, cakeyname)
	qpath := filepath.Join(tempDir, caqrname)

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

	cfile, err3 := getFile(cacertname, crtext, certpath, now)
	if err3 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err3)
	}

	kfile, err4 := getFile(cakeyname, keyext, keypath, now)
	if err4 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err4)
	}

	qfile, err5 := getFile(caqrname, pngext, qpath, now)
	if err5 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err5)
	}

	return &iamv2alphapb.AccountAuthority{
		Id:        id.String(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      req.Name,
		Curve:     curve,
		Duration:  req.Duration,
		CaCert:    cfile,
		CaKey:     kfile,
		CaQrCode:  qfile,
	}, nil
}

// GetPKI creates a new Certificate Authority using the specified request parameters.
// Returns the created Certificate Authority or an error if the operation fails.
func (b *Binding) GetPKI(_ context.Context, req *iamv2alphapb.CreateAccountRequest) (cert *typev2pb.File, key *typev2pb.File, err error) {
	nca := b.NebulaCertBinaryPath

	var c string
	var curve iamv2alphapb.Curve
	switch req.Curve {
	case iamv2alphapb.Curve_CURVE_ECDSA:
		c = "P256"
		curve = iamv2alphapb.Curve_CURVE_ECDSA
	case iamv2alphapb.Curve_CURVE_EDDSA:
		c = "25519"
		curve = iamv2alphapb.Curve_CURVE_EDDSA
	case iamv2alphapb.Curve_CURVE_UNSPECIFIED:
		c = "25519"
		curve = iamv2alphapb.Curve_CURVE_EDDSA
	default:
		c = "25519"
		curve = iamv2alphapb.Curve_CURVE_EDDSA
	}

	_ = curve

	// Write the binary to a temporary file
	tempDir, err := os.MkdirTemp("", "oeco-ca-*")
	if err != nil {
		return nil, nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to create temp file"), err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempDir)

	hostcertpath := filepath.Join(tempDir, hostcertname)
	hostkeypath := filepath.Join(tempDir, hostkeyname)

	nebula := exec.Command(nca, "keygen",
		"-curve", c,
		"-out-pub", hostcertpath,
		"-out-key", hostkeypath,
	)

	nebula.Stdout = os.Stdout
	nebula.Stderr = os.Stderr

	if err2 := nebula.Run(); err2 != nil {
		return nil, nil, ErrFailedToRunCommand.WithInternalErrorDetail(errors.New("failed to execute binary"), err2)
	}

	now := timestamppb.Now()

	cfile, err3 := getFile(hostcertname, crtext, hostcertpath, now)
	if err3 != nil {
		return nil, nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err3)
	}

	kfile, err4 := getFile(hostkeyname, keyext, hostkeypath, now)
	if err4 != nil {
		return nil, nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err4)
	}

	return cfile, kfile, nil
}

// SignCert creates a new Certificate Authority using the specified request parameters.
// Returns the created Certificate Authority or an error if the operation fails.
func (b *Binding) SignCert(_ context.Context, req *iamv2alphapb.SignAccountRequest) (*typev2pb.Credential, error) {
	nca := b.NebulaCertBinaryPath

	// Write the binary to a temporary file
	tempDir, err := os.MkdirTemp("", "oeco-ca-*")
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to create temp file"), err)
	}
	defer func(path string) {
		_ = os.RemoveAll(path)
	}(tempDir)

	cacertpath := filepath.Join(sdkv2alphalib.CredentialDirectory, cacertname)
	cakeypath := filepath.Join(sdkv2alphalib.CredentialDirectory, cakeyname)
	signedcertpath := filepath.Join(tempDir, signedcertname)
	signedqrpath := filepath.Join(tempDir, signedqrname)

	unsignedCertPath, err := saveFileTemporarily(tempDir, req.PublicCert)
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to save unsigned cert file temporarily"), err)
	}

	ip, err := getAnAvailableIPAddress()
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to get an available IP address"), err)
	}

	hostname, err := getAvailableHostname()
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to parse a hostname"), err)
	}

	groups := getAvailableGroups()

	nebula := exec.Command(nca, "sign",
		"-ca-crt", cacertpath,
		"-ca-key", cakeypath,
		"-in-pub", unsignedCertPath,
		"-ip", ip,
		"-groups", strings.Join(groups, ","),
		"-name", hostname,
		"-out-crt", signedcertpath,
		"-out-qr", signedqrpath,
	)

	nebula.Stdout = os.Stdout
	nebula.Stderr = os.Stderr

	if err2 := nebula.Run(); err2 != nil {
		return nil, ErrFailedToRunCommand.WithInternalErrorDetail(errors.New("failed to execute binary"), err2)
	}

	now := timestamppb.Now()

	cafile, err4 := getFile(cacertname, crtext, cacertpath, now)
	if err4 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("error getting ca cert file"), err4)
	}

	cfile, err3 := getFile(signedcertname, crtext, signedcertpath, now)
	if err3 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("error getting cert file"), err3)
	}

	qrfile, err4 := getFile(signedqrname, pngext, signedqrpath, now)
	if err4 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("error getting qr code cert file"), err4)
	}

	_ = cafile
	_ = cfile
	_ = qrfile

	return &typev2pb.Credential{
		Type:           0,
		MeshAccountId:  "",
		EcosystemSlug:  "",
		MeshHostname:   hostname,
		MeshIp:         ip,
		AaCertX509:     string(cafile.GetContent()),
		CertX509:       string(cfile.GetContent()),
		CertX509QrCode: base64.StdEncoding.EncodeToString(qrfile.GetContent()),
		PrivateKey:     "",
		NKey:           "",
		Groups:         groups,
		Subnets:        nil,
	}, nil
}

func getAvailableGroups() []string {
	return []string{"connector", "user"}
}

func getAvailableHostname() (string, error) {
	return "test.oeco.mesh", nil
}

func getAnAvailableIPAddress() (string, error) {
	return "192.168.100.20/24", nil
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
