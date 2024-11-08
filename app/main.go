package app

import (
	"go.uber.org/zap"
	"net/url"
	"nokowebapi/console"
	"nokowebapi/cores"
	"nokowebapi/task"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

func Main(args []string) int {
	if len(args) == 0 {
		console.Fatal("program name it-self is not defined.", zap.Int("EXIT_CODE", cores.ExitCodeFailure))
	}

	var ok bool
	var entryPointPreload string
	cores.KeepVoid(ok, entryPointPreload)

	URL := cores.Unwrap(url.Parse("http://localhost:8080"))
	deskClientPath := "bin/electron-app/nokowebview.exe"

	cores.NoErr(os.Unsetenv("ELECTRON_RUN_AS_NODE"))
	if entryPointPreload, ok = os.LookupEnv("NOKOTAN_ENTRYPOINT_PRELOAD"); !ok {
		cores.NoErr(os.Setenv("NOKOTAN_ENTRYPOINT_PRELOAD", "YES"))
	} else {
		switch strings.ToUpper(strings.TrimSpace(entryPointPreload)) {
		case "1", "Y", "YES", "TRUE":
			Server()
			return cores.ExitCodeSuccess
		}
	}

	server := func(wg *sync.WaitGroup) {
		defer wg.Done()

		var envs []string
		cores.KeepVoid(envs)

		envs = append(envs, "NOKOTAN_ENTRYPOINT_PRELOAD=YES")

		if cores.TryFetchUrlWaitForAlive(URL, 4, time.Second) {
			console.Warn("server already running.", zap.Int("EXIT_CODE", cores.ExitCodeFailure))
			return
		}

		program := args[0]
		process := cores.Unwrap(task.MakeProcess(program, args, envs, nil, os.Stdout, os.Stderr))
		cores.NoErr(process.Run())
	}

	client := func(wg *sync.WaitGroup) {
		defer wg.Done()

		var envs []string
		cores.KeepVoid(envs)

		if !cores.TryFetchUrlWaitForAlive(URL, 12, time.Second) {
			console.Fatal("server not running.", zap.Int("EXIT_CODE", cores.ExitCodeFailure))
		}

		program := path.Join(cores.Unwrap(os.Getwd()), deskClientPath)
		process := cores.Unwrap(task.MakeProcess(program, args, envs, nil, os.Stdout, os.Stderr))
		cores.NoErr(process.Run())
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	start := func(wg *sync.WaitGroup) {
		defer wg.Wait()

		go server(wg)
		go client(wg)
	}

	start(wg)

	return cores.ExitCodeSuccess
}
