package sdkv2alphalib

// Consent-based Decryption Codec
// On the client and the multiplexer server
// New Spec for this

type Codec interface {
	Name() string
	Marshal(any) ([]byte, error)
	Unmarshal([]byte, any) error
}

type ConsentBasedDecryption struct {
	// type url
	// bytes
	// encrypted fields
}

type ConsentBasedDecryptionCodec struct{}

func (cbdc *ConsentBasedDecryptionCodec) Name() string {
	return "consent-based-decryption"
}

func (cbdc *ConsentBasedDecryptionCodec) Marshal(any) ([]byte, error) {
	return nil, nil
}

func (cbdc *ConsentBasedDecryptionCodec) Unmarshal([]byte, any) error {
	return nil
}
