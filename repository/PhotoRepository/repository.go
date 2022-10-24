package photorepository

import "mygram/entity"

type PhotoRepository interface {
	FindAll() ([]entity.Photo, error)
	FindOneByID(photoID int64) (entity.Photo, error)
	Insert(req entity.Photo) (int64, int64, error)
	Update(photoID int64, req entity.Photo) (int64, int64, error)
	Delete(photoID int64) error
}
