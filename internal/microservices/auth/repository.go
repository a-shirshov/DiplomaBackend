package auth

import (
	"Diploma/internal/models"
)

type Repository interface {
	CreateUser(*models.User) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
}

type SessionRepository interface {
	SaveTokens(userId int, td *models.TokenDetails) error
	FetchAuth(accessUuid string) (int, error)
	DeleteAuth(accessUuid string) error
}