package account

import (
	"context"
	"encoding/json"

	tea "github.com/charmbracelet/bubbletea"

	pcontext "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	nebulav1ca "libs/partner/go/nebula/v1/ca"
	typev2pb "libs/protobuf/go/protobuf/gen/platform/type/v2"
	ecosystemv2alphapb "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
	iamv2alphapb "libs/public/go/protobuf/gen/platform/iam/v2alpha"
	sdkv2alphalib "libs/public/go/sdk/v2alpha"
)

// LocalAccountMsg represents a command message used for communication or signaling within a program or system.
type LocalAccountMsg struct {
	EcosystemName     string
	CIDR              string
	EcosystemPeerType ecosystemv2alphapb.EcosystemPeerType
}

// Execute processes the given ProgramContext and error, returning a command message encapsulated as tea.Msg.
func (l LocalAccountMsg) Execute(ctx *pcontext.ProgramContext, _ error) (tea.Msg, error) {
	nca := *nebulav1ca.Bound

	msg := LocalAccountMsg{}

	// Get Unsigned Public Key on local machine
	unsignedRequest := iamv2alphapb.CreateAccountRequest{
		Curve: typev2pb.Curve_CURVE_EDDSA,
	}
	unsignedCert, key, err := nca.GetPKI(context.Background(), &unsignedRequest)
	if err != nil {
		return nil, err
	}

	request := iamv2alphapb.SignAccountRequest{
		Name:              l.EcosystemName,
		PublicCert:        unsignedCert,
		EcosystemPeerType: l.EcosystemPeerType,
	}

	c, err := nca.SignCert(context.Background(), &request, nebulav1ca.WithCIDR(l.CIDR), nebulav1ca.WithEcosystemPeerType(l.EcosystemPeerType))
	if err != nil {
		return nil, err
	}

	c.PrivateKey = string(key.GetContent())

	provider, err := sdkv2alphalib.NewCredentialProvider()
	if err != nil {
		return nil, err
	}

	err = provider.SaveCredential(c)
	if err != nil {
		return nil, err
	}

	val, _ := json.MarshalIndent(&c, "", "    ")
	ctx.Logger.Debug(string(val))
	return msg, nil
}
