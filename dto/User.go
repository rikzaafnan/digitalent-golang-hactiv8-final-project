package dto

import "gopkg.in/guregu/null.v4"

type UserRegister struct {
	Age             int    `json:"age" `
	Email           string `json:"email" valid:"required~email cannot be empty"`
	Password        string `json:"password" valid:"required~password cannot be empty"`
	Username        string `json:"username" valid:"required~username cannot be empty"`
	ProfileImageUrl string `json:"profileImageUrl"`
}

type UserResponse struct {
	ID    int64  `json:"id"`
	Age   int64  `json:"age"`
	Email string `json:"email"`
	// Password string `json:"password"`
	Username string `json:"username"`
}

type UserLogin struct {
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Password string `json:"password" valid:"required~password cannot be empty"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserUpdate struct {
	Email    string `json:"email" valid:"required~email cannot be empty"`
	Username string `json:"username" valid:"required~username cannot be empty"`
}

type UserUpdateResponse struct {
	UserResponse
	UpdatedAt null.Time `json:"updatedAt"`
}
