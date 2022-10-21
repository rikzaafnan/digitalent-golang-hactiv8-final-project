package userrepository

import "mygram/entity"

type UserRepository interface {
	Insert(req entity.User) (int64, int64, error)
	Update(userID int64, req entity.User) (int64, int64, error)
	Delete(userID int64) error
	FindByID(userID int64) (entity.User, error)
	FindByUsername(username string) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}
