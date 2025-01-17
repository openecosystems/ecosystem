package infrastructurev2alphalib

import (
	"context"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"libs/public/go/sdk/v2alpha"
)

type Infrastructure struct {
	Config   *Configuration
	Bindings *sdkv2alphalib.Bindings
	Bounds   []sdkv2alphalib.Binding
}

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

func (infrastructure *Infrastructure) Run(runFunc pulumi.RunFunc, opts ...pulumi.RunOption) {
	pulumi.Run(runFunc, opts...)
}

func ShortenString(s string, limit int) string {
	if len(s) < limit {
		return s
	}

	if utf8.ValidString(s[:limit]) {
		return s[:limit]
	}
	return s[:limit+1]
}

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
