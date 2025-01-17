package main

import (
	"github.com/lyft/protoc-gen-star/v2"
	"github.com/lyft/protoc-gen-star/v2/lang/go"

	// Golang Plugins
	"libs/plugins/protoc-gen-platform/languages/go/plugins/cli_commands"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/cli_methods"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/cli_service"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/cli_system"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/client"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/entity_unspecified"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/listener"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/multiplexer"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/sdk"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/server"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/spec"
	"libs/plugins/protoc-gen-platform/languages/go/plugins/spec_entities"
	"libs/plugins/protoc-gen-platform/languages/protobuf/plugins/configuration"
	"libs/plugins/protoc-gen-platform/languages/protobuf/plugins/data_catalog"

	// Typescript Plugins
	"libs/plugins/protoc-gen-platform/languages/typescript/plugins/spec"
	"libs/plugins/protoc-gen-platform/languages/typescript/plugins/spec_index"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(

		// Go
		cli_service.GoCliServicePlugin(),
		cli_system.GoCliSystemPlugin(),
		cli_commands.GoCliCommandsPlugin(),
		cli_methods.GoCliMethodsPlugin(),
		client.GoClientPlugin(),
		entity_unspecified.GoEntityUnspecifiedPlugin(),
		listener.GoListenerPlugin(),
		server.GoServerPlugin(),
		multiplexer.GoMultiplexerPlugin(),
		spec.GoSpecPlugin(),
		spec_entities.GoSpecEntitiesPlugin(),
		sdk.GoSdkPlugin(),

		// Typescript
		spectypescript.TypeScriptSpecPlugin(),
		specindextypescript.TypeScriptSpecIndexPlugin(),

		// Protobuf
		configuration.ProtobufConfigurationPlugin(),
		data_catalog.ProtobufDataCatalogPlugin(),
	).RegisterPostProcessor(
		// goSpecTypes.GoSpecTypesPlugin(),
		pgsgo.GoFmt(),
	).Render()
}
