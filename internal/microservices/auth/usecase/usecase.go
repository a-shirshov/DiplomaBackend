package usecase

import (
	"Diploma/internal/microservices/auth"
	"Diploma/internal/models"
	"Diploma/pkg"
	"Diploma/utils"
	"errors"
	"math/rand"
	"time"
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
		return &models.User{}, &models.TokenDetails{}, err
	}
	user.Password = hash

	resultUser, err := aU.authRepo.CreateUser(user)
	if err != nil {
		return &models.User{}, &models.TokenDetails{}, err
	}

	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)

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

	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)

	err = aU.passwordHasher.VerifyPassword(user.Password, resultUser.Password)
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
	claims, err := aU.tokenManager.CheckTokenAndGetClaims(refreshToken)
	if err != nil {
		return nil, err
	}

	refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
	if !ok {
		return nil, errors.New("token problem")
	}

	userId, ok := claims["user_id"].(int)
	if !ok {
		return nil, errors.New("token problem")
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

func (aU *authUsecase) FindUserByEmail(email string) (*models.User, error) {
	resultUser, err := aU.authRepo.GetUserByEmail(email)
	if err != nil {
		return &models.User{}, nil
	}
	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (aU *authUsecase) CreateAndSavePasswordRedeemCode(email string) (int, error) {
	_, err := aU.authRepo.GetUserByEmail(email)
	if err != nil {
		return 0, err
	}

	redeemCode := createRedeemCode()
	err = aU.sessionRepo.SavePasswordRedeemCode(email, redeemCode)
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

func (aU *authUsecase) CheckRedeemCode(rdc *models.RedeemCodeStruct) (error) {
	return aU.sessionRepo.CheckRedeemCode(rdc.Email, rdc.RedeemCode)
}

func (aU *authUsecase) UpdatePassword(rdc *models.RedeemCodeStruct) (error) {
	allowed := aU.sessionRepo.CheckAccessToNewPassword(rdc.Email)
	if !allowed {
		return errors.New("operation is not allowed")
	}

	hash, err := aU.passwordHasher.GenerateHashFromPassword(rdc.Password)
	if err != nil {
		return err
	}

	return aU.authRepo.UpdatePassword(hash, rdc.Email)
}