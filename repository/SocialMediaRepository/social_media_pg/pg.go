package socialmediapg

import (
	"log"
	"mygram/entity"
	socialmediarepository "mygram/repository/SocialMediaRepository"
	"time"

	"github.com/jmoiron/sqlx"
)

// FindAll() ([]entity.SocialMedia, error)
// 	FindOneByID(socialMediaID int64) (entity.SocialMedia, error)
// 	Insert(req entity.SocialMedia) (int64, int64, error)
// 	Update(socialMediaID int64, req entity.SocialMedia) (int64, int64, error)
// 	Delete(socialMediaID int64) error

type socialMediaPG struct {
	db *sqlx.DB
}

func NewSocialMediaPG(db *sqlx.DB) socialmediarepository.SocialMediaRepository {
	return &socialMediaPG{
		db: db,
	}
}

const (
	sqlInsertSocialMedia = `INSERT INTO social_medias
					(
						name, social_media_url, user_id, created_at
					)
					VALUES ($1, $2, $3, $4)
					`
	sqlSocialMedia = `SELECT sc.id, sc.name, sc.social_media_url, sc.user_id, sc.created_at, sc.updated_at

					user.id as id_user,	user.username as user_username,user.profile_image_url
					FROM social_medias as sc
					left JOIN users as user on user.id = c.user_id
					`

	SqlDeleteSocialMedia = `DELETE social_medias where id= $1`
)

func (r *socialMediaPG) FindAll() ([]entity.SocialMedia, error) {
	var socialMedias []entity.SocialMedia

	err := r.db.Select(socialMedias, sqlSocialMedia)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return socialMedias, nil

}
func (r *socialMediaPG) FindOneByID(socialMediaID int64) (entity.SocialMedia, error) {
	var socialMedia entity.SocialMedia

	err := r.db.Select(socialMedia, sqlSocialMedia+"where sc.id=$1", socialMediaID)
	if err != nil {
		log.Println(err)
		return socialMedia, err
	}

	return socialMedia, nil
}
func (r *socialMediaPG) Insert(req entity.SocialMedia) (int64, int64, error) {

	result, err := r.db.Exec(sqlInsertSocialMedia, req.Name, req.SocialMediaURl, req.UserID, time.Now())
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	lastInserId, _ := result.LastInsertId()

	return rowsAffected, lastInserId, nil
}
func (r *socialMediaPG) Update(socialMediaID int64, req entity.SocialMedia) (int64, int64, error) {

	result, err := r.db.Exec("UPDATE social_medias SET name = $2,social_media_url= $3, updated_at = $4 where id = &1 ", socialMediaID, req.Name, req.SocialMediaURl, time.Now())
	if err != nil {
		log.Println(err)
		return 0, 0, err
	}

	rowsAffected, _ := result.RowsAffected()
	lastInserId, _ := result.LastInsertId()

	return rowsAffected, lastInserId, nil

}
func (r *socialMediaPG) Delete(socialMediaID int64) error {
	result, err := r.db.Exec(SqlDeleteSocialMedia, socialMediaID)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
