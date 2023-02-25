package auth

import (
	"Diploma/internal/models"
)

type Usecase interface {
	CreateUser(*models.User) (*models.User, *models.TokenDetails, error)
	SignIn(*models.LoginUser) (*models.User, *models.TokenDetails, error)
	Logout(au *models.AccessDetails, refreshToken string) error
	Refresh(string) (*models.Tokens, error)
	FindUserByEmail(email string) (*models.User, error)
	CreateAndSavePasswordRedeemCode(email string) (int, error)
	CheckRedeemCode(rdc *models.RedeemCodeStruct) error
	UpdatePassword(rdc *models.RedeemCodeStruct) (error)
	DeleteUser(userID int) (error)
}
