package commentrepository

import "mygram/entity"

type CommentRepository interface {
	FindAll() ([]entity.Comment, error)
	FindOneByID(commentID int64) (entity.Comment, error)
	Insert(req entity.Comment) (int64, int64, error)
	Update(commentID int64, req entity.Comment) (int64, int64, error)
	Delete(commentID int64) error
}
