package napp

import "flag"

type FlagSet flag.FlagSet

type Flag interface {
	Args() []string
	Flag(key string) FlagValue
	UnmarshalFlag(data interface{}) error
}

type FlagValue interface {
	String() string
	Int() int
	Int64() int64
	Bool() bool
}

func init() {
	flag.Parse()
}
