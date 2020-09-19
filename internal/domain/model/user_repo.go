package model

type UserRepo interface {
	GetByID(id string) (User, error)
}
