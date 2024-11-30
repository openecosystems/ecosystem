package nebulav1ca

import (
	"context"
	"errors"
	"fmt"
	"github.com/segmentio/ksuid"
	"google.golang.org/protobuf/types/known/timestamppb"
	"libs/protobuf/go/protobuf/gen/platform/type/v2"
	"libs/public/go/protobuf/gen/platform/cryptography/v2alpha"
	"libs/public/go/sdk/v2alpha"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

type Binding struct {
	NebulaCertBinaryFile *os.File
	NebulaCertBinaryPath string
}

var (
	Bound       *Binding
	BindingName = "CA_BINDING"
)

func (b *Binding) Name() string {
	return BindingName
}

func (b *Binding) Validate(_ context.Context, _ *sdkv2alphalib.Bindings) error {

	return nil
}

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
				if err := os.Chmod(nebulaCertBinary.Name(), 0755); err != nil {
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

func (b *Binding) GetBinding() interface{} {
	return Bound
}

func (b *Binding) Close() error {
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(b.NebulaCertBinaryFile.Name()) // Clean up after execution
	return nil
}

func (b *Binding) GetCertificateAuthority(_ context.Context, req *cryptographyv2alphapb.CreateCertificateAuthorityRequest) (*cryptographyv2alphapb.CertificateAuthority, error) {

	nca := b.NebulaCertBinaryPath

	// TODO: This should be done in the initial validate
	if req.Name == "" {
		return nil, sdkv2alphalib.ErrServerPreconditionFailed.WithInternalErrorDetail(errors.New("required parameter name is missing"))
	}

	c := "25519"
	curve := cryptographyv2alphapb.Curve_CURVE_EDDSA
	switch req.Curve {
	case cryptographyv2alphapb.Curve_CURVE_ECDSA:
		c = "P256"
		curve = cryptographyv2alphapb.Curve_CURVE_ECDSA
	case cryptographyv2alphapb.Curve_CURVE_EDDSA:
		c = "25519"
		curve = cryptographyv2alphapb.Curve_CURVE_EDDSA
	case cryptographyv2alphapb.Curve_CURVE_UNSPECIFIED:
		c = "25519"
		curve = cryptographyv2alphapb.Curve_CURVE_EDDSA
	default:
		c = "25519"
		curve = cryptographyv2alphapb.Curve_CURVE_EDDSA
	}

	// Write the binary to a temporary file
	tempDir, err := os.MkdirTemp("", "oeco-ca-*")
	if err != nil {
		return nil, sdkv2alphalib.ErrServerInternal.WithInternalErrorDetail(errors.New("failed to create temp file"), err)
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {

		}
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

	return &cryptographyv2alphapb.CertificateAuthority{
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
