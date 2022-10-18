package dto

import "gopkg.in/guregu/null.v4"

type CommentRequest struct {
	Message string `json:"message"`
	PhotoID int64  `json:"photoId"`
}

type CommentResponse struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoID   int64     `json:"photoId"`
	UserID    int64     `json:"userId"`
	CreatedAt null.Time `json:"createdAt"`
}

type CommentUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CommentPhoto struct {
	ID       int    `json:"-"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photoUrl"`
	UserID   int    `json:"userId"`
}

type CommentUserPhotoResponse struct {
	CommentResponse
	CommentUser
	CommentPhoto
	UpdatedAt null.Time `json:"updatedAt"`
}

type CommentUpdateRequest struct {
	Message string `json:"message"`
}

type CommentUpdateResponse struct {
	CommentPhoto
	UpdatedAt null.Time `json:"updatedAt"`
}
