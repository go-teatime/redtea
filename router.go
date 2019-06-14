package redtea

type Route interface {
	Get(urlPath string, handler interface{})
	Post(urlPath string, handler interface{})
	Put(urlPath string, handler interface{})
	Delete(urlPath string, handler interface{})
}

type Middleware interface {
	PreHandler(ctx *Context) error
	PostHandle(ctx *Context, result interface{}) (interface{}, error)
}

type router struct {
	middlewares []Middleware
}

func NewRouter() *router {
	r := new(router)
	r.middlewares = make([]Middleware, 0, 4)
	return r
}

func (r *router) Get(urlPath string, handler interface{}) {

}

func (r *router) Post(urlPath string, handler interface{}) {

}

func (r *router) Put(urlPath string, handler interface{}) {
}

func (r *router) Delete(urlPath string, handler interface{}) {

}

func (r *router) AddMiddleware(m Middleware) {
	r.middlewares = append(r.middlewares, m)
}

func (r *router) ClearMiddleware() {
	r.middlewares = make([]Middleware, 0, 4)
}
