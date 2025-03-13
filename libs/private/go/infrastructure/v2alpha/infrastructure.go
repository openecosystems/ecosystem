package infrastructurev2alphalib

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	sdkv2alphalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2alpha"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Infrastructure represents the core structure for managing configuration, bindings, and their associated instances.
type Infrastructure struct {
	Config   *Configuration
	Bindings *sdkv2alphalib.Bindings
	Bounds   []sdkv2alphalib.Binding
}

// NewInfrastructure initializes a new Infrastructure instance with specified bindings and resolved configuration.
func NewInfrastructure(bounds []sdkv2alphalib.Binding) *Infrastructure {
	ctx := context.Background()

	c := Configuration{}
	c.ResolveConfiguration()
	cfg := ResolvedConfiguration
	ResolvedConfiguration = cfg

	bindings := sdkv2alphalib.RegisterBindings(ctx, bounds)

	return &Infrastructure{
		Config:   cfg,
		Bindings: bindings,
		Bounds:   bounds,
	}
}

// Run executes a Pulumi program using the provided RunFunc and optional runtime configuration options.
func (infrastructure *Infrastructure) Run(runFunc pulumi.RunFunc, opts ...pulumi.RunOption) {
	pulumi.Run(runFunc, opts...)
}

// ShortenString truncates the input string `s` to the specified `limit` while ensuring it remains a valid UTF-8 string.
func ShortenString(s string, limit int) string {
	if len(s) < limit {
		return s
	}

	if utf8.ValidString(s[:limit]) {
		return s[:limit]
	}
	return s[:limit+1]
}

// WriteIndentedMultilineText takes a multiline string and returns it with each line prefixed by an 8-space indentation.
func WriteIndentedMultilineText(text string) string {
	indent := "        "
	lines := strings.Split(text, "\n")

	var builder strings.Builder

	for _, line := range lines {
		_, err := builder.WriteString(indent + line + "\n")
		if err != nil {
			fmt.Println(err)
		}
	}
	return builder.String()
}
