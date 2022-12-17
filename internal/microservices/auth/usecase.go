package auth

import (
	"Diploma/internal/models"
	"Diploma/utils"
)

type Usecase interface {
	CreateUser(*models.User) (*models.User, *utils.TokenDetails, error)
	SignIn(*models.LoginUser) (*models.User, *utils.TokenDetails, error)
	Logout(*utils.AccessDetails) error
	Refresh(string) (*models.Tokens, error)
}
