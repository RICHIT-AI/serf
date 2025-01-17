package command

import (
	"strings"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/richit-ai/serf/testutil"
)

func TestMembersCommandRun(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{"-rpc-addr=" + rpcAddr}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_statusFilter(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + rpcAddr,
		"-status=alive",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_statusFilter_failed(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + rpcAddr,
		"-status=(failed|left)",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_roleFilter(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + rpcAddr,
		"-role=test",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_roleFilter_failed(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + rpcAddr,
		"-role=primary",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_tagFilter(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + rpcAddr,
		"-tag=tag1=foo",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_tagFilter_failed(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + rpcAddr,
		"-tag=tag1=nomatch",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}
func TestMembersCommandRun_mutliTagFilter(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + rpcAddr,
		"-tag=tag1=foo",
		"-tag=tag2=bar",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}

func TestMembersCommandRun_multiTagFilter_failed(t *testing.T) {
	ip1, returnFn1 := testutil.TakeIP()
	defer returnFn1()

	ip2, returnFn2 := testutil.TakeIP()
	defer returnFn2()

	a1 := testAgent(t, ip1)
	defer a1.Shutdown()

	rpcAddr, ipc := testIPC(t, ip2, a1)
	defer ipc.Shutdown()

	ui := new(cli.MockUi)
	c := &MembersCommand{Ui: ui}
	args := []string{
		"-rpc-addr=" + rpcAddr,
		"-tag=tag1=foo",
		"-tag=tag2=nomatch",
	}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if strings.Contains(ui.OutputWriter.String(), a1.SerfConfig().NodeName) {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}
