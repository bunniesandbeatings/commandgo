package commandgo

import (
	"io"
	"os/exec"
)

type Runner struct {
	HandlesErrors
	Path      string
	Arguments []string
	command   *exec.Cmd
}

func NewRunner(path string, arguments ...string) *Runner {
	return &Runner{
		HandlesErrors: NewHandlesErrors(),
		Path: path,
		Arguments: arguments,
	}
}

func (runner *Runner) AddArguments(arguments ...string) {
	runner.Arguments = append(runner.Arguments, arguments...)
}

func (runner *Runner) Command(additonalArguments ...string) *exec.Cmd {
	arguments := append(runner.Arguments, additonalArguments...)
	runner.command = exec.Command(runner.Path, arguments...)

	return runner.command
}

func (runner *Runner) PipeCommand(additonalArguments ...string) (*exec.Cmd, io.WriteCloser) {
	runner.Command(additonalArguments...)
	stdin, err := runner.command.StdinPipe()

	if (err != nil) {
		runner.ErrorHandler(err)
	}

	return runner.command, stdin
}
