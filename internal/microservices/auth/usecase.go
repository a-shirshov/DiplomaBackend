package auth

import (
	"Diploma/internal/models"
	"Diploma/utils"
)

type Usecase interface {
	CreateUser(*models.User) (*models.User, error)
	SignIn(*models.LoginUser) (*models.User, *utils.TokenDetails, error)
	Logout(*utils.AccessDetails) error
	Refresh(string) (*models.Tokens, error)
}
