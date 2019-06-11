package redtea

type Route interface {
	Get(urlPath string, handler interface{})
	Post(urlPath string, handler interface{})
	Put(urlPath string, handler interface{})
	Delete(urlPath string, handler interface{})
}

type router struct {
}

func NewRouter() *router {
	return nil
}

func (r *router) Get(urlPath string, handler interface{}) {

}

func (r *router) Post(urlPath string, handler interface{}) {

}
func (r *router) Put(urlPath string, handler interface{}) {

}

func (r *router) Delete(urlPath string, handler interface{}) {

}
