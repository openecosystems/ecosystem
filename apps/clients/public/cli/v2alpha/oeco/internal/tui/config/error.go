package config

import (
	"fmt"
)

type configError struct {
	configDir string
	parser    Parser
	err       error
}

func (e configError) Error() string {
	return fmt.Sprintf(
		`Couldn't find a tui.yml or a tui.yaml configuration file.
Create one under: %s

Example of a tui.yml file:
%s

For more info, go to https://github.com/openecosystems
press q to exit.

Original error: %v`,
		TuiConfigurationFile,
		e.parser.getDefaultConfigYamlContents(),
		e.err,
	)
}

type parsingError struct {
	err error
}

func (e parsingError) Error() string {
	return fmt.Sprintf("failed parsing tui.yaml: %v", e.err)
}
