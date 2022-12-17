package repository

import (
	"Diploma/internal/errors"
	"Diploma/internal/models"
	"Diploma/utils/query"
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
	err := uR.db.QueryRowx(query.CreateUserQuery, &user.Name, &user.Surname, &user.Email, &user.Password).Scan(&user.ID)
	if err != nil {
		log.Println(err)
		return nil, errors.ErrPostgres
	}
	user.Password = ""
	return user, nil
}

func (uR *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow(query.GetUserByEmailQuery,email).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Password, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		log.Print(err.Error())
		return nil, errors.ErrPostgres
	}
	resultUser := models.ToUserModel(userDB)
	return resultUser, nil
}