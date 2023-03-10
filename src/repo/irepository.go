package repo

type IRepo[T any] interface {
	Save(*T)
	FindById(id int) T
	Update(entity T, updatedEntity *T)
	Delete(entity *T)
	FindOne(where T) T
}
