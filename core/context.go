package napp

import "context"

type Context interface {
	context.Context
	Env
	Flag() Flag
}

type defaultContext struct {
	context.Context
	Env
	flag Flag
}

func (d defaultContext) Flag() Flag {
	return d.flag
}

func DefaultContext(env Env) Context {
	return &defaultContext{
		Env:     env,
		Context: context.Background(),
	}
}
