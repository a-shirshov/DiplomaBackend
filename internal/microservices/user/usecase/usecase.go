package usecase

import (
	"Diploma/internal/microservices/user"
	"Diploma/internal/models"
	"Diploma/pkg"
	"Diploma/utils"
)

type userUsecase struct {
	userRepo user.Repository
	passwordHasher pkg.PasswordHasher
}

func NewUserUsecase(userRepo user.Repository, passwordHasher pkg.PasswordHasher) *userUsecase {
	return &userUsecase{
		userRepo: userRepo,
		passwordHasher: passwordHasher,
	}
}

func (uU *userUsecase) GetUser(userID int) (*models.User, error) {
	resultUser, err := uU.userRepo.GetUser(userID)
	if err != nil {
		return &models.User{}, err
	}
	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (uU *userUsecase) UpdateUser(user *models.User) (*models.User, error) {
	resultUser, err := uU.userRepo.UpdateUser(user)
	if err != nil {
		return &models.User{}, err
	}

	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (uU *userUsecase) UpdateUserImage(userID int, imgUUID string) (*models.User, error) {
	resultUser, err := uU.userRepo.UpdateUserImage(userID, imgUUID)
	if err != nil {
		return &models.User{}, err
	}

	resultUser.ImgUrl = utils.TryBuildImgUrl(resultUser.ImgUrl)
	return resultUser, nil
}

func (uU *userUsecase) ChangePassword(userID int, password string) (error) {
	hash, err := uU.passwordHasher.GenerateHashFromPassword(password)
	if err != nil {
		return err
	}
	
	err = uU.userRepo.ChangePassword(userID, hash)
	if err != nil {
		return err
	}
	
	return nil
}