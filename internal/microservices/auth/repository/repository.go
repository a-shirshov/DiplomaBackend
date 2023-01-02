package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	log "Diploma/pkg/logger"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
)

const logMessage = "auth:repository:"

const (
	CreateUserQuery = `insert into "user" (name, surname, email, password, date_of_birth, city) values ($1, $2, $3, $4, $5, $6) returning id;`
	GetUserByEmailQuery = `select id, name, surname, email, password, date_of_birth, city, about, img_url from "user" where email = $1;`
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
	message := logMessage + "CreateUser:"
	log.Debug(message + "started")
	err := uR.db.QueryRowx(CreateUserQuery, &user.Name, &user.Surname, &user.Email, &user.Password, &user.DateOfBirth, &user.City).Scan(&user.ID)
	if err != nil {
		log.Error(err)
		if strings.Contains(err.Error(), "(SQLSTATE 23505)") {
			return &models.User{}, customErrors.ErrUserExists
		}
		return &models.User{}, customErrors.ErrPostgres
	}
	user.Password = ""
	return user, nil
}

func (uR *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	message := logMessage + "GetUserByEmail:"
	log.Debug(message + "started")
	user := models.User{}
	err := uR.db.Get(&user, GetUserByEmailQuery, email)
	if err != nil {
		if err == sql.ErrNoRows {
			return &user, customErrors.ErrWrongEmail
		}
		log.Error(err)
		return &user, customErrors.ErrPostgres
	}
	return &user, nil
}
