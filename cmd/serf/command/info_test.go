package command

import (
	"strings"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/richit-ai/serf/testutil"
)

func TestInfoCommandRun(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &InfoCommand{Ui: ui}
	args := []string{"-rpc-addr=" + rpcAddr}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), "runtime") {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}
