package service

import (
	"log"
	"mygram/dto"
	"mygram/entity"
	commentrepository "mygram/repository/CommentRepository"
)

type CommentService interface {
	Create(req *dto.CommentRequest) (dto.CommentResponse, error)
	Update(commentID int64, req *dto.CommentUpdateRequest) (dto.CommentUpdateResponse, error)
	Delete(commentID int64) error
	FindOneByID(commentID int64) (dto.CommentResponse, error)
	FindAll() ([]dto.CommentUserPhotoResponse, error)
}

type commentService struct {
	commentRepository commentrepository.CommentRepository
}

func NewCommentService(commentRepository commentrepository.CommentRepository) *commentService {
	return &commentService{
		commentRepository: commentRepository,
	}
}

func (s *commentService) Create(req *dto.CommentRequest) (dto.CommentResponse, error) {

	var commentResponse dto.CommentResponse

	var entityComment entity.Comment
	entityComment.Message = req.Message
	entityComment.PhotoID = req.PhotoID

	_, lastInsertId, err := s.commentRepository.Insert(entityComment)
	if err != nil {
		log.Println(err)
		return commentResponse, err
	}

	comment, err := s.commentRepository.FindOneByID(lastInsertId)
	if err != nil {
		log.Println(err)
		return commentResponse, err
	}

	commentResponse = dto.CommentResponse{
		ID:      int(comment.ID),
		Message: comment.Message,
		PhotoID: comment.PhotoID,
	}
	commentResponse.CreatedAt.SetValid(comment.CereatedAt.Time)

	return commentResponse, nil
}

func (s *commentService) Update(commentID int64, req *dto.CommentUpdateRequest) (dto.CommentUpdateResponse, error) {

	var commentUpdate dto.CommentUpdateResponse

	comment, err := s.commentRepository.FindOneByID(commentID)
	if err != nil {
		log.Println(err)
		return commentUpdate, err
	}
	var entityComment entity.Comment
	entityComment.Message = entityComment.Message

	_, _, err = s.commentRepository.Update(comment.ID, entityComment)
	if err != nil {
		log.Println(err)
		return commentUpdate, err
	}

	comment, err = s.commentRepository.FindOneByID(commentID)
	if err != nil {
		log.Println(err)
		return commentUpdate, err
	}

	commentUpdate.CommentPhoto.ID = int(commentID)
	commentUpdate.CommentPhoto.Title = comment.Title.String
	commentUpdate.CommentPhoto.Caption = comment.Caption.String
	commentUpdate.CommentPhoto.PhotoUrl = comment.PhotoURL.String
	commentUpdate.CommentPhoto.UserID = int(comment.UserID)
	commentUpdate.UpdatedAt.SetValid(comment.UpdatedAt.Time)

	return commentUpdate, nil
}
func (s *commentService) Delete(commentID int64) error {

	comment, err := s.commentRepository.FindOneByID(commentID)
	if err != nil {
		log.Println(err)
		return err
	}

	err = s.commentRepository.Delete(comment.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *commentService) FindOneByID(commentID int64) (dto.CommentResponse, error) {
	var commentResponse dto.CommentResponse
	comment, err := s.commentRepository.FindOneByID(commentID)
	if err != nil {
		log.Println(err)
		return commentResponse, err
	}

	commentResponse = dto.CommentResponse{
		ID:      int(comment.ID),
		Message: comment.Message,
		PhotoID: comment.PhotoID,
	}
	commentResponse.CreatedAt.SetValid(comment.CereatedAt.Time)

	return commentResponse, nil
}

func (s *commentService) FindAll() ([]dto.CommentUserPhotoResponse, error) {

	var commentResponses []dto.CommentUserPhotoResponse

	comments, err := s.commentRepository.FindAll()
	if err != nil {
		log.Println(err)
		return commentResponses, err
	}
	for _, comment := range comments {
		commentResponse := dto.CommentResponse{
			ID:      int(comment.ID),
			Message: comment.Message,
			PhotoID: comment.PhotoID,
			UserID:  comment.UserID,
		}
		commentResponse.CreatedAt.SetValid(comment.CereatedAt.Time)

		commentUser := dto.CommentUser{
			ID:       int(comment.UserID),
			Email:    comment.Email.String,
			Username: comment.Username.String,
		}

		commentPhoto := dto.CommentPhoto{
			ID:       int(comment.PhotoID),
			Title:    comment.Title.String,
			Caption:  comment.Caption.String,
			PhotoUrl: comment.PhotoURL.String,
			UserID:   int(comment.UserID),
		}

		var commentUserPhotoResponse dto.CommentUserPhotoResponse
		commentUserPhotoResponse.CommentResponse = commentResponse
		commentUserPhotoResponse.CommentUser = commentUser
		commentUserPhotoResponse.CommentPhoto = commentPhoto
		commentUserPhotoResponse.UpdatedAt.SetValid(comment.UpdatedAt.Time)

		commentResponses = append(commentResponses, commentUserPhotoResponse)
	}

	return commentResponses, nil
}
