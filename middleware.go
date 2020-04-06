package redtea

type Middleware interface {
	PreHandler(ctx *Context) error
	PostHandle(ctx *Context, result interface{}) (interface{}, error)
}
