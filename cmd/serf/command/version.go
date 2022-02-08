package command

import (
	"fmt"

	"github.com/mitchellh/cli"
	"github.com/richit-ai/serf/serf"
)

// VersionCommand is a Command implementation prints the version.
type VersionCommand struct {
	Version string
	UI      cli.Ui
}

func (c *VersionCommand) Help() string {
	return ""
}

func (c *VersionCommand) Run(_ []string) int {
	c.UI.Output(c.Version)
	c.UI.Output(fmt.Sprintf("Agent Protocol: %d (Understands back to: %d)",
		serf.ProtocolVersionMax, serf.ProtocolVersionMin))
	return 0
}

func (c *VersionCommand) Synopsis() string {
	return "Prints the Serf version"
}
