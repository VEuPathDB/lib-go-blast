package cli

import "os/exec"

type CommandSerializable interface {
	ToCLI() *exec.Cmd
}
