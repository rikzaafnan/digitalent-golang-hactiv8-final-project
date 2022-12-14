package userpg

import (
	"fmt"
	"log"
	"mygram/entity"
	userrepository "mygram/repository/UserRepository"
	"time"

	"github.com/jmoiron/sqlx"
)

const (
	sqlInsertUser = `INSERT INTO users
					(
						username, email, password, age, profile_image_url, created_at
					)
					VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;
					`
	sqlUser       = `SELECT id, username, email, password, age, profile_image_url, created_at, updated_at FROM users`
	SqlDeleteUser = `DELETE FROM users where id= $1`
)

type userPG struct {
	db *sqlx.DB
}

func NewUserPG(db *sqlx.DB) userrepository.UserRepository {
	return &userPG{
		db: db,
	}
}

func (upg *userPG) Insert(req entity.User) (int64, int64, error) {

	var id int
	err := upg.db.QueryRowx(sqlInsertUser, req.Username, req.Email, req.Password, req.Age, req.ProfileImageUrl, time.Now()).Scan(&id)
	if err != nil {
		log.Println(err)
		fmt.Println("err  kesini ?")
		return 0, 0, err
	}

	return 0, int64(id), nil
}
func (upg *userPG) Update(userID int64, req entity.User) (int64, int64, error) {
	sqlResult, err := upg.db.Exec("UPDATE users SET username = $2, email= $3, updated_at = $3 where id = &1 ", userID, req.Username, req.Email, time.Now().UTC())
	if err != nil {
		log.Println(err)
		fmt.Println("err  kesini ?")
		return 0, 0, err
	}
	rowAffected, err := sqlResult.RowsAffected()
	if err != nil {
		return 0, 0, err
	}

	lastInsertID, err := sqlResult.LastInsertId()
	if err != nil {
		return 0, 0, err
	}

	return rowAffected, lastInsertID, nil
}
func (upg *userPG) Delete(userID int64) error {
	result, err := upg.db.Exec(SqlDeleteUser, userID)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
func (upg *userPG) FindByID(userID int64) (entity.User, error) {
	var user entity.User

	err := upg.db.Get(&user, sqlUser+" where id = $1 limit 1 ", userID)
	if err != nil {

		log.Println(sqlUser+" where id = $1  limit 1 ", userID)

		return user, err
	}

	return user, nil
}
func (upg *userPG) FindByUsername(username string) (entity.User, error) {
	var user entity.User

	log.Println(username)

	err := upg.db.Get(&user, sqlUser+" where username = $1 limit 1 ", username)
	if err != nil {

		log.Println(sqlUser+" where username = $1 limit 1 ", username)

		return user, err
	}

	return user, nil
}
func (upg *userPG) FindByEmail(email string) (entity.User, error) {
	var user entity.User

	err := upg.db.Get(&user, sqlUser+" where email = $1 limit 1 ", email)
	if err != nil {

		log.Println(sqlUser+" where email = $1 limit 1 ", email)

		return user, err
	}

	return user, nil
}
