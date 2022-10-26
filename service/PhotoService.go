package service

import (
	"fmt"
	"log"
	"mygram/dto"
	"mygram/entity"
	photorepository "mygram/repository/PhotoRepository"
)

type PhotoService interface {
	Create(req *dto.PhotoRequest) (dto.PhotoResponse, error)
	Update(photoID int64, req *dto.PhotoUpdateRequest) (dto.PhotoUpdateResponse, error)
	Delete(photoID int64) error
	FindOneByID(photoID int64) (dto.PhotoResponse, error)
	FindAll() ([]dto.PhotoUserResponse, error)
}

type photoService struct {
	photoRepository photorepository.PhotoRepository
}

func NewPhotoService(photoRepository photorepository.PhotoRepository) *photoService {
	return &photoService{
		photoRepository: photoRepository,
	}
}

func (s *photoService) Create(req *dto.PhotoRequest) (dto.PhotoResponse, error) {

	var photoResponse dto.PhotoResponse

	var entityPhoto entity.Photo
	entityPhoto.Title = req.Title
	entityPhoto.Caption = req.Caption
	entityPhoto.PhotoUrl = req.PhotoUrl

	_, lastInsertId, err := s.photoRepository.Insert(entityPhoto)
	if err != nil {
		log.Println(err)
		return photoResponse, err
	}

	photo, err := s.photoRepository.FindOneByID(lastInsertId)
	if err != nil {
		log.Println(err)
		return photoResponse, err
	}

	photoResponse = dto.PhotoResponse{
		ID:       int(photo.ID),
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   int(photo.UserID),
	}
	photoResponse.CreatedAt.SetValid(photo.CereatedAt.Time)

	return photoResponse, nil
}

func (s *photoService) Update(photoID int64, req *dto.PhotoUpdateRequest) (dto.PhotoUpdateResponse, error) {

	var photoUpdate dto.PhotoUpdateResponse

	photo, err := s.photoRepository.FindOneByID(photoID)
	if err != nil {
		log.Println(err)
		return photoUpdate, err
	}
	var entityPhoto entity.Photo
	entityPhoto.Title = req.Title
	entityPhoto.Caption = req.Caption
	entityPhoto.PhotoUrl = req.PhotoUrl

	_, _, err = s.photoRepository.Update(photoID, entityPhoto)
	if err != nil {
		log.Println(err)
		return photoUpdate, err
	}

	photo, err = s.photoRepository.FindOneByID(photoID)
	if err != nil {
		log.Println(err)
		return photoUpdate, err
	}

	photoUpdate = dto.PhotoUpdateResponse{
		ID:       int(photo.ID),
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   int(photo.UserID),
	}
	photoUpdate.UpdatedAt.SetValid(photo.UpdatedAt.Time)

	return photoUpdate, nil
}
func (s *photoService) Delete(photoID int64) error {

	photo, err := s.photoRepository.FindOneByID(photoID)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.photoRepository.Delete(photo.ID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *photoService) FindOneByID(photoID int64) (dto.PhotoResponse, error) {
	var photoResponse dto.PhotoResponse
	photo, err := s.photoRepository.FindOneByID(photoID)
	if err != nil {
		log.Println(err)
		return photoResponse, err
	}

	photoResponse = dto.PhotoResponse{
		ID:       int(photo.ID),
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoUrl,
		UserID:   int(photo.UserID),
	}
	photoResponse.CreatedAt.SetValid(photo.CereatedAt.Time)

	return photoResponse, nil
}

func (s *photoService) FindAll() ([]dto.PhotoUserResponse, error) {

	var photoUserResponses []dto.PhotoUserResponse

	photos, err := s.photoRepository.FindAll()
	if err != nil {
		log.Println(err)
		return photoUserResponses, err
	}
	for _, photo := range photos {
		var photoUser dto.PhotoUser
		photoUser.Email = photo.Email
		photoUser.Username = photo.Username

		var photoUserResponse dto.PhotoUserResponse
		photoUserResponse.ID = int(photo.ID)
		photoUserResponse.Title = photo.Title
		photoUserResponse.Caption = photo.Caption
		photoUserResponse.PhotoUrl = photo.PhotoUrl
		photoUserResponse.UserID = int(photo.UserID)
		photoUserResponse.CreatedAt.SetValid(photo.CereatedAt.Time)
		photoUserResponse.UpdatedAt.SetValid(photo.UpdatedAt.Time)
		photoUserResponse.PhotoUser = photoUser

		photoUserResponses = append(photoUserResponses, photoUserResponse)
	}

	return photoUserResponses, nil
}
