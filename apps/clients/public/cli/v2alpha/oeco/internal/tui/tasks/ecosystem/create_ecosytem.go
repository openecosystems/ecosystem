package ecosystem

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/segmentio/ksuid"

	pcontext "apps/clients/public/cli/v2alpha/oeco/internal/tui/context"
	tasks "apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks"
	accountauthority "apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks/account_authority"
	account "apps/clients/public/cli/v2alpha/oeco/internal/tui/tasks/local_account"
	ecosystemv2alphapb "libs/public/go/protobuf/gen/platform/ecosystem/v2alpha"
)

// CreateEcosystemMsg represents a command message used for communication or signaling within a program or system.
type CreateEcosystemMsg struct {
	ecosystemv2alphapb.CreateEcosystemRequest
}

// Execute processes the given ProgramContext and error, returning a command message encapsulated as tea.Msg.
func (l *CreateEcosystemMsg) Execute(ctx *pcontext.ProgramContext, _ error) (tea.Msg, error) {
	// nca := *nebulav1ca.Bound

	tasks.AddTask(tasks.Task{
		Ctx:          ctx,
		ID:           ksuid.New().String(),
		StartText:    "Create Account Authority",
		FinishedText: "Created Account Authority",
		State:        tasks.TaskStart,
		Error:        nil,
		StartTime:    time.Now(),
		TaskExecutor: accountauthority.LocalAccountAuthorityMsg{
			EcosystemName: l.Name,
		},
		Done: false,
	})

	tasks.AddTask(tasks.Task{
		Ctx:          ctx,
		ID:           ksuid.New().String(),
		StartText:    "Create Edge Service Account",
		FinishedText: "Created Edge Service Account",
		State:        tasks.TaskStart,
		Error:        nil,
		StartTime:    time.Now(),
		TaskExecutor: account.LocalAccountMsg{
			EcosystemName:     l.Name,
			CIDR:              l.Cidr,
			EcosystemPeerType: ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_EDGE,
		},
		Done: false,
	})

	tasks.AddTask(tasks.Task{
		Ctx:          ctx,
		ID:           ksuid.New().String(),
		StartText:    "Create Ecosystem Multiplexer Service Account",
		FinishedText: "Created Ecosystem Multiplexer Service Account",
		State:        tasks.TaskStart,
		Error:        nil,
		StartTime:    time.Now(),
		TaskExecutor: account.LocalAccountMsg{
			EcosystemName:     l.Name,
			CIDR:              l.Cidr,
			EcosystemPeerType: ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_ECOSYSTEM_MULTIPLEXER,
		},
		Done: false,
	})

	tasks.AddTask(tasks.Task{
		Ctx:          ctx,
		ID:           ksuid.New().String(),
		StartText:    "Create Local Service Account",
		FinishedText: "Created Local Service Account",
		State:        tasks.TaskStart,
		Error:        nil,
		StartTime:    time.Now(),
		TaskExecutor: account.LocalAccountMsg{
			EcosystemName:     l.Name,
			CIDR:              l.Cidr,
			EcosystemPeerType: ecosystemv2alphapb.EcosystemPeerType_ECOSYSTEM_PEER_TYPE_SERVICE_ACCOUNT,
		},
		Done: false,
	})

	return CreateEcosystemMsg{
		CreateEcosystemRequest: ecosystemv2alphapb.CreateEcosystemRequest{
			Slug:             l.Slug,
			Type:             l.Type,
			Name:             l.Name,
			ShortDescription: l.ShortDescription,
			Description:      l.Description,
			Cidr:             l.Cidr,
		},
	}, nil
}
