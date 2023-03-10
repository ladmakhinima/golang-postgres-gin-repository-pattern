package repo

import (
	"gorm.io/gorm"
)

type IRepository any

type GormImplement[T any] struct {
	Model *gorm.DB
}

func InitialRepo[T any](db *gorm.DB) IRepo[T] {
	return &GormImplement[T]{Model: db}
}

func (db GormImplement[T]) Save(entity *T) {
	db.Model.Create(&entity)
}

func (db GormImplement[T]) Delete(entity *T) {
	db.Model.Delete(&entity)
}

func (db GormImplement[T]) Update(entity T, updatedEntity *T) {
	db.Model.Where(entity).Updates(&updatedEntity)
}

func (db GormImplement[T]) FindById(id int) T {
	var result T
	db.Model.Where("id = ?", id).First(&result)
	return result
}

func (db GormImplement[T]) FindOne(where T) T {
	var result T
	db.Model.Where(where).First(&result)
	return result
}
