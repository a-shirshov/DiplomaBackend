package server

import (
	"Diploma/internal/models"
	"Diploma/utils"
)

type Repository interface {
	CreateUser(*models.User) (*models.User, error)
	GetUser(int) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	UpdateUser(userId int, user *models.User) (*models.User, error)
}

type SessionRepository interface {
	SaveTokens(userId int, td *utils.TokenDetails) error
	FetchAuth(accessToken string) (int, error)
	DeleteAuth(accessUuid string) error
}
