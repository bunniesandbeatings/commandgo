package commandgo

import (
	"io"
	"os/exec"
	"github.com/onsi/gomega/gexec"
	"fmt"
	"github.com/onsi/ginkgo"
)

type Runner struct {
	HandlesErrors
	Path      string
	Arguments []string
	command   *exec.Cmd
	OutWriter io.Writer
	ErrWriter io.Writer
}

func NewRunner(path string, arguments ...string) *Runner {
	return &Runner{
		HandlesErrors: NewHandlesErrors(),
		Path:          path,
		Arguments:     arguments,
		OutWriter:     ginkgo.GinkgoWriter,
		ErrWriter:     ginkgo.GinkgoWriter,
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


func (runner *Runner) Execute() *gexec.Session {
	command := runner.Command()

	session, err := gexec.Start(command, runner.OutWriter, runner.ErrWriter)
	if err != nil {
		runner.ErrorHandler(err)
	}

	return session
}


// The ExecuteWithInput function is a utility function that runs a PipeCommand and passes
// the input string to the command. It closes stdin to ensure the command completes and returns
// a gexec.Session you can use to evaluate output.
func (runner *Runner) ExecuteWithInput(input string) *gexec.Session {
	command, stdin := runner.PipeCommand()

	session, err := gexec.Start(command, runner.OutWriter, runner.ErrWriter)
	if err != nil {
		runner.ErrorHandler(err)
	}

	fmt.Fprint(stdin, input)
	stdin.Close()

	return session
}
