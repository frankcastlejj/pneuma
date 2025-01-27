//+build !windows,!js,!darwin

package commands

import (
	"context"
	"os/exec"
	"syscall"
)

func getSysProcAttrs() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setsid: true,
	}
}

func getShellCommand(ctx context.Context, executor, command string) *exec.Cmd {
	var cmd *exec.Cmd
	switch executor {
	case "python":
		cmd = exec.CommandContext(ctx, "python", "-c", command)
	case "bash":
		cmd = exec.CommandContext(ctx, "bash", "-c", command)
	case "zsh":
		cmd = exec.CommandContext(ctx, "zsh", "-c", command)
	default:
		cmd = exec.CommandContext(ctx, "sh", "-c", command)
	}
	cmd.SysProcAttr = getSysProcAttrs()
	return cmd
}