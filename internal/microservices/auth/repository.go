package auth

import (
	"Diploma/internal/models"
	"Diploma/utils"
)

type Repository interface {
	CreateUser(*models.User) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
}

type SessionRepository interface {
	SaveTokens(userId int, td *utils.TokenDetails) error
	FetchAuth(accessToken string) (int, error)
	DeleteAuth(accessUuid string) error
}
