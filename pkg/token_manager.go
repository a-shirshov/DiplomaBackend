package pkg

import (
	"Diploma/internal/models"

	"github.com/golang-jwt/jwt/v4"
)

type TokenManager interface {
	CreateToken(userId int) (*models.TokenDetails, error)
	ExtractTokenMetadata(token string) (*models.AccessDetails, error)
	CheckTokenAndGetClaims(refreshToken string) (jwt.MapClaims, error)
}