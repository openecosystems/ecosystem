package account

import (
	"context"
	"encoding/json"

	tea "github.com/charmbracelet/bubbletea"

	pcontext "github.com/openecosystems/ecosystem/apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	typev2pb "github.com/openecosystems/ecosystem/libs/protobuf/go/protobuf/gen/platform/type/v2"
	nebulav1ca "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/bindings/nebula/ca"
	iamv2alphapb "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta/gen/platform/iam/v2alpha"

	sdkv2betalib "github.com/openecosystems/ecosystem/libs/public/go/sdk/v2beta"
)

// LocalAccountMsg represents a command message used for communication or signaling within a program or system.
type LocalAccountMsg struct {
	EcosystemName     string
	CIDR              string
	EcosystemPeerType typev2pb.PeerType
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
		Name:       l.EcosystemName,
		PublicCert: unsignedCert,
		PeerType:   l.EcosystemPeerType,
	}

	c, err := nca.SignCert(context.Background(), &request, nebulav1ca.WithCIDR(l.CIDR), nebulav1ca.WithPeerType(l.EcosystemPeerType))
	if err != nil {
		return nil, err
	}

	c.PrivateKey = string(key.GetContent())

	provider, err := sdkv2betalib.NewCredentialProvider()
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
