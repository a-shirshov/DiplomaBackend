package repository

import (
	"Diploma/internal/models"
	"Diploma/internal/errors"
	"log"

	"github.com/jackc/pgx"
)

type AuthRepository struct {
	db *pgx.ConnPool
}

func NewAuthRepository(db *pgx.ConnPool) *AuthRepository {
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
	_, err := uR.db.Exec("CreateUserQuery", &user.Name, &user.Surname, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.ErrPostgres
	}
	return resultUser, nil
}

func (uR *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow("GetUserByEmailQuery",email).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Password, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.ErrPostgres
	}
	resultUser := models.ToUserModel(userDB)
	return resultUser, nil
}