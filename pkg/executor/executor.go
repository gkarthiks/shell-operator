package executor

import (
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/romana/rlog"
)

var ExecutorLock = &sync.Mutex{}

func Run(cmd *exec.Cmd, debug bool) error {
	ExecutorLock.Lock()
	defer ExecutorLock.Unlock()

	if debug {
		dir := ""
		if cmd.Dir != "" {
			dir = " in '" + cmd.Dir + "'"
		}
		rlog.Debugf("Executing command%s: '%s'", dir, strings.Join(cmd.Args, " "))
	}

	return cmd.Run()
}

func Output(cmd *exec.Cmd) (output []byte, err error) {
	ExecutorLock.Lock()
	defer ExecutorLock.Unlock()

	output, err = cmd.Output()
	return
}

func MakeCommand(dir string, entrypoint string, args []string, envs []string) *exec.Cmd {
	cmd := exec.Command(entrypoint, args...)
	cmd.Env = append(cmd.Env, envs...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
