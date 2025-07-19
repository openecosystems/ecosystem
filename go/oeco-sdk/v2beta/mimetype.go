package sdkv2betalib

import (
	"fmt"
	"mime"
)

// MimeTypeRegistry is a global instance of mimeTypeRegistry that provides MIME type retrieval by file extension.
var MimeTypeRegistry = new(mimeTypeRegistry)

// mimeTypeRegistry is a structure for handling MIME type resolutions based on file extensions.
type mimeTypeRegistry struct{}

// mimeTypes is a mapping of file extensions to their corresponding MIME types for certificate and key files.
var mimeTypes = map[string]string{
	".crt": "application/x-x509-ca-cert",
	".pem": "application/x-x509-ca-cert",
	".cer": "application/x-x509-ca-cert",
	".key": "application/x-pkcs12",
	".p12": "application/x-pkcs12",
	".pfx": "application/x-pkcs12",
}

// init initializes the application by adding custom MIME types to the mime package and logs any errors encountered.
func init() {
	var errs []error
	for k, v := range mimeTypes {
		err := mime.AddExtensionType(k, v)
		if err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		fmt.Println("error adding mimetypes: ", errs)
	}
}

// GetMimeType returns the MIME type for a given file extension, including the leading period (e.g., ".png").
func (r *mimeTypeRegistry) GetMimeType(fileExtensionWithPeriod string) string {
	return mime.TypeByExtension(fileExtensionWithPeriod)
}
