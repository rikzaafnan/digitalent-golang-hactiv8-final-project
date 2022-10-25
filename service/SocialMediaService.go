package service

import (
	"fmt"
	"log"
	"mygram/dto"
	"mygram/entity"
	socialmediarepository "mygram/repository/SocialMediaRepository"
	"time"
)

type SocialMediaService interface {
	Create(req *dto.SocialMediaRequest) (dto.SocialMediaResponse, error)
	Update(socialMediaID int64, req *dto.SocialMediaUpdateRequest) (dto.SocialMediaUpdateResponse, error)
	Delete(socialMediaID int64) error
	FindOneByID(socialMediaID int64) (dto.SocialMediaResponse, error)
	FindAll() ([]dto.SocialMediaResponse, error)
}

type socialMediaService struct {
	socialMediaRepository socialmediarepository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository socialmediarepository.SocialMediaRepository) *socialMediaService {
	return &socialMediaService{
		socialMediaRepository: socialMediaRepository,
	}
}

func (s *socialMediaService) Create(req *dto.SocialMediaRequest) (dto.SocialMediaResponse, error) {

	var socialMedia dto.SocialMediaResponse

	var entitySocialMedia entity.SocialMedia
	entitySocialMedia.Name = req.Name
	entitySocialMedia.SocialMedia.SetValid(req.SocialMediaUrl)
	_, lastInsertId, err := s.socialMediaRepository.Insert(entitySocialMedia)
	if err != nil {
		log.Println(err)
		return socialMedia, err
	}

	socialMedia, err = s.FindOneByID(lastInsertId)
	if err != nil {
		log.Println(err)
		return socialMedia, err
	}

	return socialMedia, nil
}

func (s *socialMediaService) Update(socialMediaID int64, req *dto.SocialMediaUpdateRequest) (dto.SocialMediaUpdateResponse, error) {

	var socialMediaUpdate dto.SocialMediaUpdateResponse

	socialMedia, err := s.FindOneByID(socialMediaID)
	if err != nil {
		log.Println(err)
		return socialMediaUpdate, err
	}

	var entitySocialMedia entity.SocialMedia
	entitySocialMedia.Name = req.Name
	entitySocialMedia.SocialMedia.SetValid(req.SocialMediaUrl)

	_, _, err = s.socialMediaRepository.Update(socialMediaID, entitySocialMedia)
	if err != nil {
		log.Println(err)
		return socialMediaUpdate, err
	}

	socialMediaUpdate.ID = socialMedia.ID
	socialMediaUpdate.Name = socialMedia.Name
	socialMediaUpdate.SocialMediaUrl = socialMedia.SocialMediaUrl
	socialMediaUpdate.UserID = socialMedia.UserID
	socialMediaUpdate.UpdatedAt.SetValid(time.Now())

	return socialMediaUpdate, nil
}
func (s *socialMediaService) Delete(socialMediaID int64) error {

	socialMedia, err := s.FindOneByID(socialMediaID)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.socialMediaRepository.Delete(int64(socialMedia.ID))
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (s *socialMediaService) FindOneByID(socialMediaID int64) (dto.SocialMediaResponse, error) {

	var socialMediaResponse dto.SocialMediaResponse

	socialMedia, err := s.socialMediaRepository.FindOneByID(socialMediaID)
	if err != nil {
		log.Println(err)
		return socialMediaResponse, err
	}

	socialMediaResponse.ID = socialMedia.ID
	socialMediaResponse.Name = socialMedia.Name
	socialMediaResponse.SocialMediaUrl = socialMedia.SocialMedia.String
	socialMediaResponse.UserID = socialMedia.UserID
	socialMediaResponse.CreatedAt.SetValid(socialMedia.CereatedAt.Time)

	return socialMediaResponse, nil
}

func (s *socialMediaService) FindAll() ([]dto.SocialMediaResponse, error) {

	var socialMediaResponses []dto.SocialMediaResponse

	socialMedias, err := s.socialMediaRepository.FindAll()
	if err != nil {
		log.Println(err)
		return socialMediaResponses, err
	}
	for _, socialMedia := range socialMedias {
		var socialMediaResponse dto.SocialMediaResponse

		socialMediaResponse.ID = socialMedia.ID
		socialMediaResponse.Name = socialMedia.Name
		socialMediaResponse.SocialMediaUrl = socialMedia.SocialMedia.String
		socialMediaResponse.UserID = socialMedia.UserID
		socialMediaResponse.CreatedAt.SetValid(socialMedia.CereatedAt.Time)

		socialMediaResponses = append(socialMediaResponses, socialMediaResponse)
	}

	return socialMediaResponses, nil
}
