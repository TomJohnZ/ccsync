package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecCommandInDir(dir, command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("TASKRC=%v/.taskrc", dir),
		fmt.Sprintf("TASKDATA=%v", dir),
	)
	return cmd.Run()
}

func ExecCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	return cmd.Run()
}

func ExecCommandForOutputInDir(dir, command string, args ...string) ([]byte, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(),
		fmt.Sprintf("TASKRC=%v/.taskrc", dir),
		fmt.Sprintf("TASKDATA=%v", dir),
	)
	return cmd.Output()
}
