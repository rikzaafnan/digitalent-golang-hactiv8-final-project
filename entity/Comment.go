package entity

import "gopkg.in/guregu/null.v4"

type Comment struct {
	ID         int64     `db:"id"`
	UserID     int64     `db:"user_id"`
	PhotoID    int64     `db:"photo_id"`
	Message    string    `db:"message"`
	CereatedAt null.Time `db:"created_at"`
	UpdatedAt  null.Time `db:"updated_at"`
}
