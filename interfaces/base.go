package interfaces

import "go-pattern/models"

type IRepository interface {
	FetchAll(receiver interface{}, cons string, args ...interface{}) error
	FindById(receiver interface{}, id uint) error
	Store(objects models.IModel) error
	Update(objects interface{}) error
	Delete(obj interface{}, id uint) error
}
