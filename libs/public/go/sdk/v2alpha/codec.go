package sdkv2alphalib

// Consent-based Decryption Codec
// On the client and the multiplexer server
// New Spec for this

// Codec defines an interface for encoding and decoding data with specific serialization implementations.
// Name returns the name identifier for the codec implementation.
// Marshal serializes the given object into a byte slice.
// Unmarshal deserializes the provided byte slice into the specified object.
type Codec interface {
	Name() string
	Marshal(any) ([]byte, error)
	Unmarshal([]byte, any) error
}

// ConsentBasedDecryption represents a structure for handling decryption requiring user consent.
// It encapsulates necessary fields like the type URL, encrypted data, and relevant metadata.
type ConsentBasedDecryption struct {
	// type url
	// bytes
	// encrypted fields
}

// ConsentBasedDecryptionCodec is a struct for handling consent-based decryption encoding and decoding functionality.
type ConsentBasedDecryptionCodec struct{}

// Name returns the name of the codec as a string.
func (cbdc *ConsentBasedDecryptionCodec) Name() string {
	return "consent-based-decryption"
}

// Marshal serializes the provided data into a byte array using the ConsentBasedDecryptionCodec logic.
func (cbdc *ConsentBasedDecryptionCodec) Marshal(any) ([]byte, error) {
	return nil, nil
}

// Unmarshal decodes encrypted data bytes into the specified target structure, using consent-based decryption logic.
func (cbdc *ConsentBasedDecryptionCodec) Unmarshal([]byte, any) error {
	return nil
}
