package sdkv2alphalib

import (
	"fmt"
	"mime"
)

var MimeTypeRegistry = new(mimeTypeRegistry)

type mimeTypeRegistry struct{}

var mimeTypes = map[string]string{
	".crt": "application/x-x509-ca-cert",
	".pem": "application/x-x509-ca-cert",
	".cer": "application/x-x509-ca-cert",
	".key": "application/x-pkcs12",
	".p12": "application/x-pkcs12",
	".pfx": "application/x-pkcs12",
}

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

func (r *mimeTypeRegistry) GetMimeType(fileExtensionWithPeriod string) string {
	return mime.TypeByExtension(fileExtensionWithPeriod)
}
