package repository

import (
	"Diploma/internal/errors"
	"Diploma/internal/models"
	"Diploma/utils"
	"log"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (uR *AuthRepository) CreateUser(user *models.User) (*models.User, error) {
	resultUser := &models.User{
		Name: user.Name,
		Surname: user.Surname,
		Email: user.Email,
	}
	_, err := uR.db.Exec(utils.CreateUserQuery , &user.Name, &user.Surname, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrPostgres
	}
	return resultUser, nil
}

func (uR *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow(utils.GetUserByEmailQuery,email).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Password, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.ErrPostgres
	}
	resultUser := models.ToUserModel(userDB)
	return resultUser, nil
}