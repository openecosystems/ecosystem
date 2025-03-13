package accountauthority

import (
	"context"
	"encoding/json"

	tea "github.com/charmbracelet/bubbletea"

	pcontext "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	nebulav1ca "libs/partner/go/nebula/v1/ca"
	iamv2alphapb "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// LocalAccountAuthorityMsg represents a command message used for communication or signaling within a program or system.
type LocalAccountAuthorityMsg struct {
	EcosystemName string
}

// Execute processes the given ProgramContext and error, returning a command message encapsulated as tea.Msg.
func (l LocalAccountAuthorityMsg) Execute(ctx *pcontext.ProgramContext, _ error) (tea.Msg, error) {
	nca := *nebulav1ca.Bound

	msg := LocalAccountAuthorityMsg{}

	request := iamv2alphapb.CreateAccountAuthorityRequest{
		Name:  l.EcosystemName,
		Curve: 0,
	}

	ca, err := nca.GetAccountAuthority(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	response := iamv2alphapb.CreateAccountAuthorityResponse{
		AccountAuthority: ca,
	}

	provider, err := sdkv2alphalib.NewCredentialProvider()
	if err != nil {
		return nil, err
	}

	err = provider.SaveCredential(ca.Credential)
	if err != nil {
		return nil, err
	}

	val, _ := json.MarshalIndent(&response, "", "    ")
	ctx.Logger.Debug(string(val))
	return msg, nil
}
