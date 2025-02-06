package ecosystemv2alphapbint

import (
	"context"
	"encoding/json"
	"fmt"
	slog "log"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"

	iamv2alphapbint "apps/clients/public/cli/v2alpha/oeco/internal/iam/v2alpha"

	"connectrpc.com/connect"

	specv2pb "libs/protobuf/go/protobuf/gen/platform/spec/v2"
	ecosystemv2alphapb "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
	ecosystemv2alphapbsdk "libs/public/go/sdk/gen/ecosystem/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"

	"google.golang.org/protobuf/encoding/protojson"
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
	Args:    cobra.MaximumNArgs(1),
}

// createModel initializes a connector model and optionally sets up logging based on the provided command flags.
func createModel(cmd *cobra.Command, settings *specv2pb.SpecSettings) (*CreateEcosystemModel, *os.File) {
	var loggerFile *os.File

	d := cmd.Flag("debug").Value.String()
	debug, _ := strconv.ParseBool(d)

	if debug {
		var fileErr error
		newConfigFile, fileErr := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0o666)
		if fileErr == nil {
			log.SetOutput(newConfigFile)
			log.SetTimeFormat(time.Kitchen)
			log.SetReportCaller(true)
			log.SetLevel(log.DebugLevel)
			log.Debug("Logging to debug.log")
		} else {
			loggerFile, _ = tea.LogToFile("debug.log", "debug")
			slog.Print("Failed setting up logging", fileErr)
		}
	} else {
		log.SetOutput(os.Stderr)
		log.SetLevel(log.FatalLevel)
	}

	return NewCreateEcosystemModel(settings), loggerFile
}

// init initializes the `Cmd` execution logic by setting up the model, logging, cleanup, and running the TUI program.
func init() {
	CreateEcosystemV2AlphaCmd.Run = func(cmd *cobra.Command, _ []string) {
		settings := cmd.Root().Context().Value(sdkv2alphalib.SettingsContextKey).(*specv2pb.SpecSettings)

		model, logger := createModel(cmd, settings)

		if logger != nil {
			defer func(logger *os.File) {
				_ = logger.Close()
			}(logger)
		}

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
			fmt.Println("Form completed!")
			fmt.Println("Name:", model.Data.Name)
			fmt.Println("Email:", model.Data.EcosystemType)

			fmt.Println(model.Data)

			//_r := ecosystemv2alphapb.CreateEcosystemRequest{}
			//_r.Name = model.Data.Name

			log.Debug("Calling createEcosystem ecosystem")

			// Example JSON request argument
			// requestArg := `{"name": "oeco"}`

			// Example JSON request
			requestValue := `{"name": "123"}`

			err := iamv2alphapbint.CreateAccountV2AlphaCmd.Flags().Set("request", requestValue)
			if err != nil {
				fmt.Println("Error setting flag:", err)
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
				fmt.Println(err)
				os.Exit(1)
			}

			sdkv2alphalib.Overrides.FieldMask = createEcosystemFieldMask
			sdkv2alphalib.Overrides.ValidateOnly = createEcosystemValidateOnly

			request := connect.NewRequest[ecosystemv2alphapb.CreateEcosystemRequest](&_r)
			// Add GZIP Support: connect.WithSendGzip(),
			url := "https://" + sdkv2alphalib.Config.Platform.Mesh.Endpoint
			if sdkv2alphalib.Config.Platform.Insecure {
				url = "http://" + sdkv2alphalib.Config.Platform.Mesh.Endpoint
			}
			client := *ecosystemv2alphapbsdk.NewEcosystemServiceSpecClient(sdkv2alphalib.Config, url, connect.WithInterceptors(sdkv2alphalib.NewCLIInterceptor(sdkv2alphalib.Config, sdkv2alphalib.Overrides)))

			response, err := client.CreateEcosystem(context.Background(), request)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			val, _ := json.MarshalIndent(&response, "", "    ")
			fmt.Println(string(val))
		}
	}
}

// cleanup recovers from any panic that occurred and logs the recovery message before quitting the tea program.
func cleanup() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
	_ = tea.Quit()
}
