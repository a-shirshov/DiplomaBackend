package usecase

import (
	"Diploma/internal/microservices/auth"
	"Diploma/internal/models"
	"Diploma/utils"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type authUsecase struct {
	authRepo auth.Repository
	sessionRepo auth.SessionRepository
}

func NewAuthUsecase(userRepo auth.Repository, sessionRepo auth.SessionRepository) *authUsecase {
	return &authUsecase{
		authRepo: userRepo,
		sessionRepo: sessionRepo,
	}
}

func (uU *authUsecase) CreateUser(user *models.User) (*models.User, error) {
	hash, err := utils.GenerateHashFromPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hash
	return uU.authRepo.CreateUser(user)
}

func (uU *authUsecase) SignIn(user *models.LoginUser) (*models.User, *utils.TokenDetails, error) {
	resultUser, err := uU.authRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, nil, err
	}

	_, err = utils.VerifyPassword(user.Password, resultUser.Password)
	if err != nil {
		return nil, nil, err
	}

	resultUser.Password = ""
	td, err := utils.CreateToken(resultUser.ID)
	if err != nil {
		return nil, nil, err
	}


	err = uU.sessionRepo.SaveTokens(resultUser.ID, td)
	if err != nil {
		return nil, nil, err
	}

	return resultUser, td, nil
}

func (aU *authUsecase) Logout(au *utils.AccessDetails) error {
	return aU.sessionRepo.DeleteAuth(au.AccessUuid)
}

func (aU *authUsecase) Refresh(refreshToken string) (*models.Tokens, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		log.Printf("Returned refresh")
		return []byte(viper.GetString("REFRESH_TOKEN")), nil
	})

	if err != nil {
		log.Println(err.Error())
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

		err = aU.sessionRepo.DeleteAuth(refreshUuid)
		if err != nil  { //if any goes wrong
			return nil, errors.New("unauthorized")
		}

		ts, err := utils.CreateToken(int(userId))
		if err != nil {
		   return nil, err
		}

		err = aU.sessionRepo.SaveTokens(int(userId), ts)
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