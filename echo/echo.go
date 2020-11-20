package echo

import (
	napp "github.com/ohanan/napp/core"
)

type echo struct {
}

func (e echo) Run(ctx napp.Context) error {
	ctx.StdOut().WriteString("lala")
	return nil
}

func NewEcho() napp.App {
	return &echo{}
}
