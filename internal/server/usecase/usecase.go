package Usecase

import (
	"Diploma/internal/models"
	"Diploma/internal/server"
	"Diploma/utils"
	"errors"
	"fmt"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type userUsecase struct {
	userRepo server.Repository
	sessionRepo server.SessionRepository
}

func NewUserUsecase(userRepo server.Repository, sessionRepo server.SessionRepository) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
		sessionRepo: sessionRepo,
	}
}


func (uU *userUsecase) CreateUser(user *models.User) (*models.User, error) {
	hash, err := utils.GenerateHashFromPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	return uU.userRepo.CreateUser(user)
}

func (uU *userUsecase) GetUser(userId int) (*models.User, error) {
	return uU.userRepo.GetUser(userId)
}

func (uU *userUsecase) SignIn(user *models.User) (*models.User, *utils.TokenDetails, error) {
	resultUser, err := uU.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, nil, err
	}

	_, err = utils.VerifyPassword(user.Password, resultUser.Password)
	if err != nil {
		return nil, nil, err
	}

	resultUser.Password = ""
	uId := int(resultUser.ID)
	td, err := utils.CreateToken(uId)
	if err != nil {
		return nil, nil, err
	}

	err = uU.sessionRepo.SaveTokens(uId, td)
	if err != nil {
		return nil, nil, err
	}

	return resultUser, td, nil
}

func (uU *userUsecase) Logout(au *utils.AccessDetails) error {
	return uU.sessionRepo.DeleteAuth(au.AccessUuid)
}

func (uU *userUsecase) Refresh(refreshToken string) (*models.Tokens, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("REFRESH_SECRET")), nil
	})

	if err != nil {
		return nil, errors.New("Refresh token expired")
	}

	if !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			return nil, err
		}

		userId, err := strconv.ParseInt(fmt.Sprintf("%.f",claims["user_id"]),10,0)
		if err != nil {
			return nil, err
		}

		err = uU.sessionRepo.DeleteAuth(refreshUuid)
		if err != nil  { //if any goes wrong
			return nil, errors.New("unauthorized")
		}

		ts, err := utils.CreateToken(int(userId))
		if err != nil {
		   return nil, err
		}

		err = uU.sessionRepo.SaveTokens(int(userId), ts)
		if err != nil {
			return nil, err
		}

		tokens := &models.Tokens{
			AccessToken: ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}
		return tokens, nil

	} else {
		return nil, err
	}
}

func (uU *userUsecase) UpdateUser(au *utils.AccessDetails, user *models.User) (*models.User, error) {
	userId, err := uU.sessionRepo.FetchAuth(au.AccessUuid)
	if err != nil {
		return nil, err
	}
	resultUser, err := uU.userRepo.UpdateUser(userId, user)
	if err != nil {
		return nil, err
	}
	return resultUser, nil
}