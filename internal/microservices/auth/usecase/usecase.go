package usecase

import (
	"Diploma/internal/microservices/auth"
	"Diploma/internal/models"
	"Diploma/pkg"
	"Diploma/utils"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type authUsecase struct {
	authRepo auth.Repository
	sessionRepo auth.SessionRepository
	passwordHasher pkg.PasswordHasher
	tokenManager pkg.TokenManager
}

func NewAuthUsecase(userRepo auth.Repository, sessionRepo auth.SessionRepository, 
		passwordHasher pkg.PasswordHasher, tokenManager pkg.TokenManager) *authUsecase {
	return &authUsecase{
		authRepo: userRepo,
		sessionRepo: sessionRepo,
		passwordHasher: passwordHasher,
		tokenManager: tokenManager,
	}
}

func (aU *authUsecase) CreateUser(user *models.User) (*models.User, *models.TokenDetails, error) {
	hash, err := aU.passwordHasher.GenerateHashFromPassword(user.Password)
	if err != nil {
		return nil, nil, err
	}
	user.Password = hash

	resultUser, err := aU.authRepo.CreateUser(user)
	if err != nil {
		return nil, nil, err
	}

	resultUser.ImgUrl = utils.BuildImgUrl(resultUser.ImgUrl)

	td, err := aU.tokenManager.CreateToken(resultUser.ID)
	if err != nil {
		return nil, nil, err
	}

	err = aU.sessionRepo.SaveTokens(resultUser.ID, td)
	if err != nil {
		return nil, nil, err
	}

	return resultUser, td, nil
}

func (aU *authUsecase) SignIn(user *models.LoginUser) (*models.User, *models.TokenDetails, error) {
	resultUser, err := aU.authRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, nil, err
	}

	resultUser.ImgUrl = utils.BuildImgUrl(resultUser.ImgUrl)

	_, err = aU.passwordHasher.VerifyPassword(user.Password, resultUser.Password)
	if err != nil {
		return nil, nil, err
	}

	resultUser.Password = ""
	td, err := aU.tokenManager.CreateToken(resultUser.ID)
	if err != nil {
		return nil, nil, err
	}


	err = aU.sessionRepo.SaveTokens(resultUser.ID, td)
	if err != nil {
		return nil, nil, err
	}

	return resultUser, td, nil
}

func (aU *authUsecase) Logout(au *models.AccessDetails) error {
	return aU.sessionRepo.DeleteAuth(au.AccessUuid)
}

func (aU *authUsecase) Refresh(refreshToken string) (*models.Tokens, error) {
	token, err := checkToken(refreshToken)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if !(ok && token.Valid) {
		return nil, err
	}

	refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
	if !ok {
		return nil, err
	}

	userId, err := strconv.ParseInt(fmt.Sprintf("%.f",claims["user_id"]),10,0)
	if err != nil {
		return nil, err
	}

	err = aU.sessionRepo.DeleteAuth(refreshUuid)
	if err != nil  {
		return nil, errors.New("unauthorized")
	}

	ts, err := aU.tokenManager.CreateToken(int(userId))
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
}

func checkToken(refreshToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		log.Printf("Returned refresh")
		return []byte(viper.GetString("REFRESH_TOKEN")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Refresh token expired")
	}

	return token, nil
}

func (aU *authUsecase) FindUserByEmail(email string) (*models.User, error) {
	return aU.authRepo.GetUserByEmail(email)
}

func (aU *authUsecase) CreateAndSavePasswordRedeemCode(email string) (int, error) {
	redeemCode := createRedeemCode()
	err := aU.sessionRepo.SavePasswordRedeemCode(email, redeemCode)
	return redeemCode, err
}

func createRedeemCode() int {
	max := 10
	min := 0
	var resultCode int 
	multiplier := 1
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		number := rand.Intn(max - min) + min
		resultCode = resultCode + number * multiplier
		multiplier *= 10
	}
	return resultCode
}

func (aU *authUsecase) CheckRedeemCode(rdc models.RedeemCodeStruct) (error) {
	err :=  aU.sessionRepo.CheckRedeemCode(rdc.Email, rdc.RedeemCode)
	if err != nil {
		return err
	}

	hash, err := aU.passwordHasher.GenerateHashFromPassword(rdc.Password)
	if err != nil {
		return err
	}

	return aU.authRepo.UpdatePassword(hash, rdc.Email)
}