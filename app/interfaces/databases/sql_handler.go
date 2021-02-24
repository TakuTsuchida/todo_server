package databases

type SqlHandler interface {
	FindAll(object interface{})
}
