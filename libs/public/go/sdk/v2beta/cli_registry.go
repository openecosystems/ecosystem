package sdkv2betalib

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CommandRegistry is a global instance of the Commands type to manage CLI commands.
// commands stores a map of FullCommandName to their respective cobra.Command.
var (
	CommandRegistry *Commands = new(Commands)
	commands                  = make(map[FullCommandName]*cobra.Command)
)

// Commands represents a registry to manage mapping of FullCommandName to cobra commands.
type Commands struct {
	systemsByName map[FullCommandName]*cobra.Command
}

// FullCommandName represents a unique identifier for a command, including its name and version.
type FullCommandName struct {
	Name    string
	Version string
}

// IsValid checks if both the Name and Version fields in FullCommandName are non-empty and returns true if valid.
func (s FullCommandName) IsValid() bool {
	if s.Name == "" || s.Version == "" {
		return false
	}
	return true
}

// RegisterCommands initializes and registers a set of commands and returns the mapped collection of commands.
func (c *Commands) RegisterCommands() map[FullCommandName]*cobra.Command {
	c.systemsByName = commands
	return commands
}

// RegisterCommand associates a FullCommandName with a cobra.Command and updates the systemsByName map.
func (c *Commands) RegisterCommand(name FullCommandName, cmd *cobra.Command) {
	commands[name] = cmd
	c.systemsByName = commands
}

// GetCommandByFullCommandName retrieves a command by its full command name from the registered commands map.
// Returns an error if the provided name is invalid or if the command does not exist.
func (c *Commands) GetCommandByFullCommandName(name FullCommandName) (*cobra.Command, error) {
	if !name.IsValid() {
		return nil, fmt.Errorf("invalid system name or version number in your configuration file: %s", name)
	}

	command, ok := c.systemsByName[name]
	if !ok {
		return nil, fmt.Errorf("cannot find the system or version number identified in your configuration file: %s", name)
	}

	return command, nil
}
