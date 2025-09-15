package main

import (
	pgs "github.com/lyft/protoc-gen-star/v2"
	pgsgo "github.com/lyft/protoc-gen-star/v2/lang/go"

	// Go Plugins v2beta
	clicommandsv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/cli_commands"
	climethodsv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/cli_methods"
	cliservicev2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/cli_service"
	clisystemv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/cli_system"
	clisystemsv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/cli_systems"
	entityunspecifiedv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/entity_unspecified"
	listenerv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/listener"
	multiplexerv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/multiplexer"
	sdkv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/sdk"
	sdkconnectorv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/sdk_connector"
	serverv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/server"
	specv2beta "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/go/plugins/v2beta/spec"

	// Protobuf Plugins
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/protobuf/plugins/configuration"
	"github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/protobuf/plugins/data_catalog"

	// Typescript Plugins
	clienttypescript "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/typescript/plugins/client"
	protobufindextypescript "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/typescript/plugins/protobuf_index"
	spectypescript "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/typescript/plugins/spec"
	specindextypescript "github.com/openecosystems/ecosystem/libs/plugins/protoc-gen-platform/languages/typescript/plugins/spec_index"
)

func main() {
	pgs.Init(
		pgs.DebugEnv("DEBUG"),
	).RegisterModule(

		// Go v2beta
		cliservicev2beta.GoCliServicePlugin(),
		clisystemv2beta.GoCliSystemPlugin(),
		clisystemsv2beta.GoCliSystemsPlugin(),
		clicommandsv2beta.GoCliCommandsPlugin(),
		climethodsv2beta.GoCliMethodsPlugin(),
		entityunspecifiedv2beta.GoEntityUnspecifiedPlugin(),
		listenerv2beta.GoListenerPlugin(),
		serverv2beta.GoServerPlugin(),
		multiplexerv2beta.GoMultiplexerPlugin(),
		specv2beta.GoSpecPlugin(),
		sdkv2beta.GoSdkPlugin(),
		sdkconnectorv2beta.GoSdkConnectorPlugin(),

		// Typescript
		spectypescript.TypeScriptSpecPlugin(),
		specindextypescript.TypeScriptSpecIndexPlugin(),
		protobufindextypescript.TypeScriptProtobufIndexPlugin(),
		clienttypescript.TypeScriptClientPlugin(),

		// Protobuf
		configuration.ProtobufConfigurationPlugin(),
		data_catalog.ProtobufDataCatalogPlugin(),
	).RegisterPostProcessor(
		// goSpecTypes.GoSpecTypesPlugin(),
		pgsgo.GoFmt(),
	).Render()
}
