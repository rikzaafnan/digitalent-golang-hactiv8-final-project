package dto

import "gopkg.in/guregu/null.v4"

type PhotoRequest struct {
	Title    string `json:"title" valid:"required~title cannot be empty"`
	Caption  string `json:"caption" valid:"required~caption cannot be empty"`
	PhotoUrl string `json:"photoUrl" valid:"required~photoUrl cannot be empty"`
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
	Title    string `json:"title" valid:"required~title cannot be empty"`
	Caption  string `json:"caption" valid:"required~caption cannot be empty"`
	PhotoUrl string `json:"photoUrl" valid:"required~photoUrl cannot be empty"`
}

type PhotoUpdateResponse struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photoUrl"`
	UserID    int       `json:"userId"`
	UpdatedAt null.Time `json:"updatedAt"`
}
