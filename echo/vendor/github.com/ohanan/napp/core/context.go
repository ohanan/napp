package napp

import "context"

type Context interface {
	context.Context
	Env
}

type defaultContext struct {
	Env
	context.Context
}

func DefaultContext(env Env) Context {
	return &defaultContext{
		Env:     env,
		Context: context.Background(),
	}
}
