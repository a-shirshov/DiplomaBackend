package server

import (
	"Diploma/internal/models"
	"Diploma/utils"
)

type Usecase interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(int) (*models.User, error)
	SignIn(*models.User) (*models.User, *utils.TokenDetails, error)
	Logout(*utils.AccessDetails) error
	Refresh(string) (*models.Tokens, error)
	UpdateUser(*utils.AccessDetails, *models.User) (*models.User, error)
}
