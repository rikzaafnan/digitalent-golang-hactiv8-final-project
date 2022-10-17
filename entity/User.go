package entity

import "gopkg.in/guregu/null.v4"

type User struct {
	ID         int64     `db:"id"`
	Username   string    `db:"username"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Age        string    `db:"age"`
	CereatedAt null.Time `db:"created_at"`
	UpdatedAt  null.Time `db:"updated_at"`
}
