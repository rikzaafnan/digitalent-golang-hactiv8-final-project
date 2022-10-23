package entity

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/guregu/null.v4"
)

type User struct {
	ID              int64     `db:"id"`
	Username        string    `db:"username"`
	Email           string    `db:"email"`
	Password        string    `db:"password"`
	Age             int64     `db:"age"`
	ProfileImageUrl string    `db:"profile_image_url"`
	CereatedAt      null.Time `db:"created_at"`
	UpdatedAt       null.Time `db:"updated_at"`
}

// hashPassword
func (u *User) HashPass() error {
	salt := 8
	password := []byte(u.Password)
	hash, err := bcrypt.GenerateFromPassword(password, salt)
	if err != nil {
		log.Println("gagal generate password")
		return errors.New("something went error")
	}
	u.Password = string(hash)
	return nil
}

// compare password
func (u *User) ComparePassword(uerPassword string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(uerPassword))

	return err == nil

}
