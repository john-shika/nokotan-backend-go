package cores

import (
	"errors"
	"io"
	"os/exec"
)

type ProcessStateOverlayImpl interface {
	Exited() bool
	ExitCode() int
	Pid() int
}

type ProcessStateOverlay struct {
	exited   bool
	exitCode int
	pid      int
}

func NewProcessStateOverlay(exited bool, exitCode int, pid int) ProcessStateOverlayImpl {
	return &ProcessStateOverlay{
		exited:   exited,
		exitCode: exitCode,
		pid:      pid,
	}
}

func (s *ProcessStateOverlay) Exited() bool {
	return s.exited
}

func (s *ProcessStateOverlay) ExitCode() int {
	return s.exitCode
}

func (s *ProcessStateOverlay) Pid() int {
	return s.pid
}

func MakeProcess(path string, args []string, envs []string, stdin io.Reader, stdout io.Writer, stderr io.Writer) (cmd *exec.Cmd, err error) {
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

func GetProcessState(cmd *exec.Cmd) (ProcessStateOverlayImpl, error) {
	if cmd.ProcessState != nil {
		return cmd.ProcessState, nil
	}
	if cmd.Process != nil {
		return NewProcessStateOverlay(false, 0, cmd.Process.Pid), nil
	}
	return nil, errors.New("invalid process state")
}
