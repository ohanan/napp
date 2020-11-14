package napp

import "strconv"

type ExitError interface {
	error
	ExitCode() int
}

type exitError struct {
	code int
}

func (e exitError) Error() string {
	return "exit with code as " + strconv.Itoa(e.code)
}

func (e exitError) ExitCode() int {
	return e.code
}

func ExitWith(code int) ExitError {
	return &exitError{code: code}
}

func GetExitCode(err error) int {
	if err == nil {
		return 0
	}
	if ee, ok := err.(ExitError); ok {
		return ee.ExitCode()
	}
	return 1
}
