package napp

type App interface {
	Run(ctx Context) error
}
