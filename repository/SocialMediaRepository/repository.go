package socialmediarepository

import "mygram/entity"

type SocialMediaRepository interface {
	FindAll() ([]entity.SocialMedia, error)
	FindOneByID(socialMediaID int64) (entity.SocialMedia, error)
	Insert(req entity.SocialMedia) (int64, int64, error)
	Update(socialMediaID int64, req entity.SocialMedia) (int64, int64, error)
	Delete(socialMediaID int64) error
}
