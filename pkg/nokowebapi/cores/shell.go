package cores

import (
	"io"
	"os/exec"
)

func makeProcess(path string, args []string, envs []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (cmd *exec.Cmd, err error) {
	if path, err = exec.LookPath(path); err != nil {
		return nil, err
	}

	// create command executor
	cmd = exec.Command(path, args...)

	// binding all parameters
	cmd.Stdin = stdin
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Env = envs

	// return command executor
	return cmd, nil
}
