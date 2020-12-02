package repositories

import (
	"go-pattern/interfaces"
	"go-pattern/models"
	"gorm.io/gorm"
)

type BaseRepository struct {
	Db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{db}
}

func (b BaseRepository) FetchAll(receiver interface{}, cons string, args ...interface{}) error {
	return b.Db.Where(cons, args...).Find(receiver).Error
}

func (b BaseRepository) FindById(receiver interface{}, id uint) error {
	return b.Db.First(receiver, id).Error
}

func (b BaseRepository) Store(object models.IModel) error {
	object.Setter(b.NewKey(object))
	return b.Db.Create(object).Error
}

func (b BaseRepository) Update(objects interface{}) error {
	return b.Db.Updates(objects).Error
}

func (b BaseRepository) Delete(obj interface{}, id uint) error {
	return b.Db.Delete(obj, id).Error
}

func (b BaseRepository) NewKey(obj models.IModel) uint {
	var newKey uint
	b.Db.Table(obj.TableName()).Select("count(*)").Scan(&newKey)
	return newKey + 1
}

var _ interfaces.IRepository = &BaseRepository{}
