package functions

import (
	"fmt"
	// "github.com/pkg/errors"
	// "os/exec"
	// "strings"
)

// DefaultFunctions returns a map with the all the default functions. Any
// new default function has to be added to function.
func DefaultFunctions() map[string]func(Args, *Output) error {
	return map[string]func(Args, *Output) error{
		"ShellCmd": ShellCmd,
	}
}

func ShellCmd(args Args, output *Output) error {
	cmd, err := args.String("cmd")
	if err != nil {
		return err
	}

	fmt.Printf("Executing shell command %s\n", cmd)
	// cmdPieces := strings.Split(cmd, "\t")

	// if len(cmdPieces) < 2 {
	// 	return exec.Command(cmdPieces[0]).Output()
	// }

	// return exec.Command(cmdPieces[0], cmdPieces[1:]...).Output()

	return nil
}
