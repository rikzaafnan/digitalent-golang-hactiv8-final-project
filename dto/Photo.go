package dto

import "gopkg.in/guregu/null.v4"

type PhotoRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
}

type PhotoResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl"`
	UserID    int       `json:"userId"`
	CreatedAt null.Time `json:"createdAt"`
}

type PhotoUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoUserResponse struct {
	PhotoResponse
	PhotoUser
	UpdatedAt null.Time `json:"updatedAt"`
}

type PhotoUpdateRequest struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
}

type PhotoUpdateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl"`
	UserID    int       `json:"userId"`
	UpdatedAt null.Time `json:"updatedAt"`
}
