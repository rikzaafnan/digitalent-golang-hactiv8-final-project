package photopg

import (
	"fmt"
	"log"
	"mygram/entity"
	photorepository "mygram/repository/PhotoRepository"
	"time"

	"github.com/jmoiron/sqlx"
)

type photoPG struct {
	db *sqlx.DB
}

func NewPhotoPG(db *sqlx.DB) photorepository.PhotoRepository {
	return &photoPG{
		db: db,
	}
}

const (
	sqlInsertPhoto = `INSERT INTO photos
					(
						title, caption, photo_url, user_id, created_at
					)
					VALUES ($1, $2, $3, $4,$5) RETURNING id;
					`
	sqlPhoto = `SELECT p.id, p.title, p.caption, p.photo_url, p.user_id, p.created_at, p.updated_at
					 	user.email as user_email, user.username as user_username
					FROM photos as p
					left JOIN users as user on user.id = p.user_id`

	SqlDeletePhoto = `DELETE FROM photos where id= $1`
)

func (r *photoPG) FindAll() ([]entity.Photo, error) {
	var photos []entity.Photo

	err := r.db.Select(photos, sqlPhoto)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return photos, nil

}
func (r *photoPG) FindOneByID(photoID int64) (entity.Photo, error) {
	var photo entity.Photo

	err := r.db.Select(photo, sqlPhoto+"where p.id=$1", photoID)
	if err != nil {
		log.Println(err)
		return photo, err
	}

	return photo, nil
}
func (r *photoPG) Insert(req entity.Photo) (int64, int64, error) {

	var id int
	err := r.db.QueryRowx(sqlInsertPhoto, req.Title, req.Caption, req.PhotoUrl, req.UserID, time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		fmt.Println("err  kesini ?")
		return 0, 0, err
	}

	return 0, int64(id), nil

}
func (r *photoPG) Update(photoID int64, req entity.Photo) (int64, int64, error) {

	result, err := r.db.Exec("UPDATE photos SET title = $2, caption= $3,photo_url= $3, updated_at = $4 where id = &1 ", photoID, req.Title, req.Caption, req.PhotoUrl, time.Now())
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	lastInserId, _ := result.LastInsertId()

	return rowsAffected, lastInserId, nil

}
func (r *photoPG) Delete(photoID int64) error {
	result, err := r.db.Exec(SqlDeletePhoto, photoID)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
