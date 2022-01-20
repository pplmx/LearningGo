package demo

type Repository interface {
	Save(obj interface{}) error
	FindAll() ([]interface{}, error)
	FindById(id interface{}) (interface{}, error)
	DeleteById(id interface{}) error
}
