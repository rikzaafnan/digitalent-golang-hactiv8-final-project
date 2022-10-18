package dto

import "gopkg.in/guregu/null.v4"

type UserRegister struct {
	Age             int    `json:"age"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Username        string `json:"username"`
	ProfileImageUrl string `json:"profileImageUrl"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	// Password string `json:"password"`
	Username string `json:"username"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdate struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserUpdateResponse struct {
	UserResponse
	UpdatedAt null.Time `json:"updatedAt"`
}
