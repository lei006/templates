package exec_fix

import (
	"context"
	"os/exec"
	"strings"
	"syscall"
)

func CommandEx(name, args string) *exec.Cmd {

	cmd := exec.Command(name, args)
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: name + " " + args}
	return cmd
}

func CommandContextEx(ctx context.Context, name string, arg ...string) *exec.Cmd {
	if ctx == nil {
		panic("nil Context")
	}
	cmd := exec.CommandContext(ctx, name, arg...)
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: name + " " + strings.Join(arg, " ")}

	return cmd
}
