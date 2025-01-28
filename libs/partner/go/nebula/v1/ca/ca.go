package nebulav1ca

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"

	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	iamv2alphapb "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
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

	_ca := "ca"
	cext := ".crt"
	kext := ".key"
	qext := ".png"
	cname := _ca + cext
	kname := _ca + kext
	qname := _ca + qext
	cpath := filepath.Join(tempDir, cname)
	kpath := filepath.Join(tempDir, kname)
	qpath := filepath.Join(tempDir, qname)

	nebula := exec.Command(nca, "ca",
		"-name", req.Name,
		"-curve", c,
		"-out-crt", cpath,
		"-out-key", kpath,
		"-out-qr", qpath,
	)

	nebula.Stdout = os.Stdout
	nebula.Stderr = os.Stderr

	if err2 := nebula.Run(); err2 != nil {
		return nil, ErrFailedToRunCommand.WithInternalErrorDetail(errors.New("failed to execute binary"), err2)
	}

	id := ksuid.New()
	now := timestamppb.Now()

	cfile, err3 := getFile(cname, cext, cpath, now)
	if err3 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err3)
	}

	kfile, err4 := getFile(kname, kext, kpath, now)
	if err4 != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(err4)
	}

	qfile, err5 := getFile(qname, qext, qpath, now)
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
