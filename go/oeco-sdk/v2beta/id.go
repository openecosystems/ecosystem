package sdkv2betalib

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

// GenerateID generates a unique entity identifier based on the provided Entity.
func GenerateID(entity Entity) string {
	id := ksuid.New().String()
	return fmt.Sprintf("%s-%s", entity.TypeName(), id)
}
