package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"

	// Golang Plugins
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/cli_commands"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/cli_methods"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/cli_service"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/cli_system"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/cli_systems"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/client"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/entity_unspecified"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/listener"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/multiplexer"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/sdk"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/server"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/spec"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/spec_entities"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/protobuf/plugins/configuration"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/protobuf/plugins/data_catalog"

	// Typescript Plugins
	protobufindextypescript "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/typescript/plugins/protobuf_index"
	spectypescript "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/typescript/plugins/spec"
	specindextypescript "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/typescript/plugins/spec_index"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(

		// Go
		cli_service.GoCliServicePlugin(),
		cli_system.GoCliSystemPlugin(),
		cli_systems.GoCliSystemsPlugin(),
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
		protobufindextypescript.TypeScriptProtobufIndexPlugin(),

		// Protobuf
		configuration.ProtobufConfigurationPlugin(),
		data_catalog.ProtobufDataCatalogPlugin(),
	).RegisterPostProcessor(
		// goSpecTypes.GoSpecTypesPlugin(),
		pgsgo.GoFmt(),
	).Render()
}
