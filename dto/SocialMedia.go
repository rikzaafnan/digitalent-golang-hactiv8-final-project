package dto

import "gopkg.in/guregu/null.v4"

type SocialMediaRequest struct {
	Name           string `json:"name" valid:"required~name cannot be empty"`
	SocialMediaUrl int64  `json:"socialMediaUrl" valid:"required~socialMediaUrl cannot be empty"`
}

type SocialMediaResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl int64     `json:"socialMediaUrl"`
	UserID         int64     `json:"userId"`
	CreatedAt      null.Time `json:"createdAt"`
}

type SocialMediaUser struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profileImageUrl"`
}

type SocialMediaUpdateRequest struct {
	Name           string `json:"name" valid:"required~name cannot be empty"`
	SocialMediaUrl int64  `json:"socialMediaUrl" valid:"required~socialMediaUrl cannot be empty"`
}

type SocialMediaUpdateResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl int64     `json:"socialMediaUrl"`
	UserID         int64     `json:"userId"`
	UpdatedAt      null.Time `json:"updatedAt"`
}
