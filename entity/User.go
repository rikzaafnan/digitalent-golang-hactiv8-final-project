package entity

import "gopkg.in/guregu/null.v4"

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
