package functions

import (
	"github.com/pkg/errors"
)

func MissingArgErr(arg string) error {
	return errors.Errorf("missing function arg: %s", arg)
}

func WrongArgTypeErr(arg, argType string) error {
	return errors.Errorf("function expected the %s arg to be %s type", arg, argType)
}
