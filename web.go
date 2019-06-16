package redtea

type Web struct {
	router *router
}

func NewWeb() *Web {
	w := new(Web)
	return w
}

func (w *Web) ServeHttp() {

}

func (w *Web) Get(urlPath string, handler interface{}) {
	w.router.Get(urlPath, handler)
}

func (w *Web) Post(urlPath string, handler interface{}) {
	w.router.Post(urlPath, handler)
}

func (w *Web) Put(urlPath string, handler interface{}) {
	w.router.Put(urlPath, handler)
}

func (w *Web) Delete(urlPath string, handler interface{}) {
	w.router.Delete(urlPath, handler)
}
