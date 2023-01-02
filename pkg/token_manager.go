package pkg

import (
	"Diploma/internal/models"
	"net/http"
)

type TokenManager interface {
	CreateToken(userId int) (*models.TokenDetails, error)
	ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error)
}