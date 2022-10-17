package entity

import "gopkg.in/guregu/null.v4"

type SocialMedia struct {
	ID          int64       `db:"id"`
	Name        string      `db:"name"`
	SocialMedia null.String `db:"social_media"`
	UserID      int64       `db:"user_id"`
	CereatedAt  null.Time   `db:"created_at"`
	UpdatedAt   null.Time   `db:"updated_at"`
}
