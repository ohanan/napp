package napp

import "os"

type Env interface {
	StdIn() *os.File
	StdOut() *os.File
	StdErr() *os.File
	Flags() []string
}

func NewDefaultEnv(flags []string) Env { return &defaultEnv{flags: flags} }

type defaultEnv struct {
	flags []string
}

func (e defaultEnv) StdIn() *os.File { return os.Stdin }

func (e defaultEnv) StdOut() *os.File { return os.Stdout }

func (e defaultEnv) StdErr() *os.File { return os.Stderr }

func (e defaultEnv) Flags() []string { return e.flags }

var _ Env = (*defaultEnv)(nil)
