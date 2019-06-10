package redtea

type Route interface {
	Get()
	Post()
	Put()
	Delete()
}

type router struct {

}
