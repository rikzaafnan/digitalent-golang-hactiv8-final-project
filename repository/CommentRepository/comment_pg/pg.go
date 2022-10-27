package commentpg

import (
	"fmt"
	"log"
	"mygram/entity"
	commentrepository "mygram/repository/CommentRepository"
	"time"

	"github.com/jmoiron/sqlx"
)

type commentPG struct {
	db *sqlx.DB
}

func NewCommentPG(db *sqlx.DB) commentrepository.CommentRepository {
	return &commentPG{
		db: db,
	}
}

const (
	sqlInsertComment = `INSERT INTO comments
					(
						message, photo_id, user_id, created_at
					)
					VALUES ($1, $2, $3, $4) RETURNING id;
					`
	sqlComment = `SELECT c.id, c.message, c.photo_id, c.user_id, c.created_at, c.updated_at,
					user.email, user.username,
					p.title, p.caption, p.photo_url
					FROM comments as p
					left JOIN users as user on user.id = c.user_id
					left JOIN photos as p on p.id = c.photo_id
					`

	SqlDeleteComment = `DELETE FROM comments where id= $1`
)

func (r *commentPG) FindAll() ([]entity.Comment, error) {
	var comments []entity.Comment

	err := r.db.Select(comments, sqlComment)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return comments, nil

}
func (r *commentPG) FindOneByID(commentID int64) (entity.Comment, error) {
	var comment entity.Comment

	err := r.db.Select(comment, sqlComment+"where c.id=$1", commentID)
	if err != nil {
		log.Println(err)
		return comment, err
	}

	return comment, nil
}
func (r *commentPG) Insert(req entity.Comment) (int64, int64, error) {

	var id int
	err := r.db.QueryRowx(sqlInsertComment, req.Message, req.PhotoID, req.UserID, time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		fmt.Println("err  kesini ?")
		return 0, 0, err
	}

	return 0, int64(id), nil
}
func (r *commentPG) Update(commentID int64, req entity.Comment) (int64, int64, error) {

	result, err := r.db.Exec("UPDATE comments SET message = $2, updated_at = $3 where id = &1 ", commentID, req.Message, time.Now())
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	lastInserId, _ := result.LastInsertId()

	return rowsAffected, lastInserId, nil

}
func (r *commentPG) Delete(commentID int64) error {
	result, err := r.db.Exec(SqlDeleteComment, commentID)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
