package entity

import (
	"mygram/dto"

	"gopkg.in/guregu/null.v4"
)

type Photo struct {
	ID            int64  `db:"id"`
	Title         string `db:"title"`
	Caption       string `db:"caption"`
	PhotoUrl      string `db:"photo_url"`
	UserID        int64  `db:"user_id"`
	dto.PhotoUser `db:"user_id"`
	CereatedAt    null.Time `db:"created_at"`
	UpdatedAt     null.Time `db:"updated_at"`
}
