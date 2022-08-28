package models

import (
	"mime/multipart"
)

type RegistrationUserRequest struct {
	Name string `json:"name" example:"Artyom" binding:"required"`
	Surname string `json:"surname" example:"Shirshov" binding:"required"`
	Email string `json:"email" example:"artyom@mail.ru"`
	Password string `json:"password" example:"12345678"`
}

type RegistrationUserResponse struct {
	Name string `json:"name" example:"Artyom" binding:"required"`
	Surname string `json:"surname" example:"Shirshov" binding:"required"`
	Email string `json:"email" example:"artyom@mail.ru"`
}

type LoginUserRequest struct {
	Email string `json:"email" example:"artyom@mail.ru"`
	Password string `json:"password" example:"12345678"`
}

type UserProfile struct {
	Name string `json:"name" example:"Artyom" binding:"required"`
	Surname string `json:"surname" example:"Shirshov" binding:"required"`
	Email string `json:"email" example:"artyom@mail.ru"`
	About string `json:"about,omitempty"`
	ImgUrl string `json:"imgUrl,omitempty"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token,omitempty" example:"4ffc5f18-99d8-47f6-8141-faf2c2f5a24e"`
}

type ErrorMessageBadRequest struct {
	Message string `json:"message" example:"Bad request"`
}

type ErrorMessageInternalServer struct {
	Message string `json:"message" example:"Server problems"`
}

type ErrorMessageUnauthorized struct {
	Message string `json:"message" example:"User is not authorized"`
}

type ErrorMessageUnprocessableEntity struct {
	Message string `json:"message" example:"Wrong Json Request"`
}

type UserWithTokensResponse struct {
	User UserProfile
	Tokens Tokens
}

type ImageFormRequest struct {
	Image *multipart.FileHeader `form:"image" swaggerignore:"true"`
}

type UpdateUserMpfd struct {
	Json string 
	ImageFormRequest
}