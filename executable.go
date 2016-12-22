package commandgo

import (
	"io"
	"os/exec"
	"github.com/onsi/gomega/gexec"
	"fmt"
	"github.com/onsi/ginkgo"
)

type ExecutableContext struct {
	HandlesErrors
	Path        string
	Arguments   []string
	command     *exec.Cmd
	OutWriter   io.Writer
	ErrWriter   io.Writer
}

func NewExecutableContext(path string, arguments ...string) *ExecutableContext {
	return &ExecutableContext{
		HandlesErrors: NewHandlesErrors(),
		Path:          path,
		Arguments:     arguments,
		OutWriter:     ginkgo.GinkgoWriter,
		ErrWriter:     ginkgo.GinkgoWriter,
	}
}

func (executable *ExecutableContext) AddArguments(arguments ...string) {
	executable.Arguments = append(executable.Arguments, arguments...)
}

func (executable *ExecutableContext) Command(additonalArguments ...string) *exec.Cmd {
	arguments := append(executable.Arguments, additonalArguments...)
	executable.command = exec.Command(executable.Path, arguments...)

	return executable.command
}

func (executable *ExecutableContext) PipeCommand(additonalArguments ...string) (*exec.Cmd, io.WriteCloser) {
	executable.Command(additonalArguments...)
	stdin, err := executable.command.StdinPipe()

	if err != nil {
		executable.ErrorHandler(err)
	}

	return executable.command, stdin
}

func (executable *ExecutableContext) Execute() *gexec.Session {
	command := executable.Command()

	session, err := gexec.Start(command, executable.OutWriter, executable.ErrWriter)
	if err != nil {
		executable.ErrorHandler(err)
	}

	return session
}

// The ExecuteWithInput function is a utility function that runs a PipeCommand and passes
// the input string to the command. It closes stdin to ensure the command completes and returns
// a gexec.Session you can use to evaluate output.
func (executable *ExecutableContext) ExecuteWithInput(input string) *gexec.Session {
	command, stdin := executable.PipeCommand()

	session, err := gexec.Start(command, executable.OutWriter, executable.ErrWriter)
	if err != nil {
		executable.ErrorHandler(err)
	}

	fmt.Fprint(stdin, input)
	stdin.Close()

	return session
}
