package pkg

import (
	"Diploma/internal/models"
)

type TokenManager interface {
	CreateToken(userId int) (*models.TokenDetails, error)
	ExtractTokenMetadata(token string) (*models.AccessDetails, error)
}