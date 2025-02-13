package ecosystemv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"connectrpc.com/connect"
	tea "github.com/charmbracelet/bubbletea"
	clog "github.com/charmbracelet/log"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/encoding/protojson"

	iamv2alphapbint "apps/clients/public/cli/v2alpha/oeco/internal/iam/v2alpha"
	cliv2alphalib "libs/public/go/cli/v2alpha"
	ecosystemv2alphapb "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
	ecosystemv2alphapbsdk "libs/public/go/sdk/gen/ecosystem/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

var (
	// createEcosystemRequest      string
	createEcosystemFieldMask    string
	createEcosystemValidateOnly bool
)

// CreateEcosystemV2AlphaCmd defines a CLI command named "connector" for interaction, supporting one optional argument and a short description.
var CreateEcosystemV2AlphaCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create an ecosystem",
	Version: "",
	// Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, _ []string) {
		settings := cmd.Context().Value(sdkv2alphalib.SettingsContextKey).(*cliv2alphalib.Configuration)
		log := cmd.Context().Value(sdkv2alphalib.LoggerContextKey).(*clog.Logger)

		model := NewCreateEcosystemModel(settings)

		defer cleanup()

		p := tea.NewProgram(
			model,
			// tea.WithAltScreen(),
			tea.WithReportFocus(),
			// tea.WithMouseCellMotion(),
		)

		if _, err := p.Run(); err != nil {
			log.Fatal("Failed starting the TUI", err)
		}

		if model.state == stateDone {
			log.Info("Form completed!")
			log.Info("Name:", model.Data.Name)
			log.Info("Email:", model.Data.EcosystemType)

			log.Info(model.Data)

			//_r := ecosystemv2alphapb.CreateEcosystemRequest{}
			//_r.Name = model.Data.Name

			log.Debug("Calling createEcosystem ecosystem")

			// Example JSON request argument
			// requestArg := `{"name": "oeco"}`

			// Example JSON request
			requestValue := `{"name": "123"}`

			err := iamv2alphapbint.CreateAccountV2AlphaCmd.Flags().Set("request", requestValue)
			if err != nil {
				log.Error("Error setting flag:", err)
				return
			}

			// iamv2alphapbint.CreateAccountV2AlphaCmd.SetArgs([]string{"request", requestArg})
			iamv2alphapbint.CreateAccountV2AlphaCmd.Run(cmd, []string{})

			//_request, err := cmd.Flags().GetString("request")
			//if err != nil {
			//	fmt.Println(err)
			//	os.Exit(1)
			//}
			//if _request == "" {
			//	_request = "{}"
			//}

			_request := `{"name": "123"}`

			_r := ecosystemv2alphapb.CreateEcosystemRequest{}
			err = protojson.Unmarshal([]byte(_request), &_r)
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}

			sdkv2alphalib.Overrides.FieldMask = createEcosystemFieldMask
			sdkv2alphalib.Overrides.ValidateOnly = createEcosystemValidateOnly

			request := connect.NewRequest[ecosystemv2alphapb.CreateEcosystemRequest](&_r)
			// Add GZIP Support: connect.WithSendGzip(),
			url := "https://" + settings.Platform.Mesh.Endpoint
			if settings.Platform.Insecure {
				url = "http://" + settings.Platform.Mesh.Endpoint
			}
			client := *ecosystemv2alphapbsdk.NewEcosystemServiceSpecClient(&settings.Platform, url, connect.WithInterceptors(cliv2alphalib.NewCLIInterceptor(settings, sdkv2alphalib.Overrides)))

			response, err := client.CreateEcosystem(context.Background(), request)
			if err != nil {
				log.Error(err)
				os.Exit(1)
			}

			val, _ := json.MarshalIndent(&response, "", "    ")
			log.Info("Response: ", string(val))
		}
	},
}

//// init initializes the `Cmd` execution logic by setting up the model, logging, cleanup, and running the TUI program.
//func init() {
//	CreateEcosystemV2AlphaCmd.Run = func(cmd *cobra.Command, _ []string) {
//	}
//}

// cleanup recovers from any panic that occurred and logs the recovery message before quitting the tea program.
func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
	_ = tea.Quit()
}
