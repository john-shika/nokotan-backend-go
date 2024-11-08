package cores

import "os"

const (
	EXIT_SUCCESS int = 0
	EXIT_FAILURE int = 1
)

type MainFunc func([]string) int

func (m MainFunc) Call(args []string) int {
	return m(args)
}

func ApplyMainFunc(mainFunc MainFunc) {
	exitCode := mainFunc.Call(os.Args)
	os.Exit(exitCode)
}
