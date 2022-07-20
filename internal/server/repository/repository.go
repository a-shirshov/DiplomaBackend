package repository

import (
	"Diploma/internal/models"

	"github.com/jackc/pgx"
)

type UserRepository struct {
	db *pgx.ConnPool
}

func NewUserRepository(db *pgx.ConnPool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func toUserModel(userDB *models.UserDB) (*models.User) {
	user := &models.User{
		ID: userDB.ID,
		Name: userDB.Name,
		Surname: userDB.Surname,
		Email: userDB.Email,
		Password: userDB.Password,
	}

	if userDB.About.Valid {
		user.About = userDB.About.String
	}

	if userDB.ImgUrl.Valid {
		user.ImgUrl = userDB.ImgUrl.String
	}

	return user
}

func (uR *UserRepository) CreateUser(user *models.User) (*models.User, error) {
	resultUser := &models.User{
		Name: user.Name,
		Surname: user.Surname,
		Email: user.Email,
	}
	_, err := uR.db.Exec("CreateUserQuery", &user.Name, &user.Surname, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return resultUser, nil
}

func (uR *UserRepository) GetUser(Id int) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow("GetUserQuery",Id).Scan(&userDB.Name, &userDB.Surname, &userDB.Email, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		return nil, err
	}
	resultUser := toUserModel(userDB)
	return resultUser, nil
}

func (uR *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow("GetUserByEmailQuery",email).Scan(&userDB.ID, &userDB.Name, &userDB.Surname, &userDB.Email, &userDB.Password, &userDB.About, &userDB.ImgUrl)
	if err != nil {
		return nil, err
	}
	resultUser := toUserModel(userDB)
	return resultUser, nil
}

func (uR *UserRepository) UpdateUser(userId int, user *models.User) (*models.User, error) {
	userDB := &models.UserDB{}
	err := uR.db.QueryRow("UpdateUserQuery", &user.Name, &user.Surname, &user.About, &userId).Scan(
		&userDB.Name, &userDB.Surname, &userDB.Email, &userDB.About, &userDB.ImgUrl,
	)
	if err != nil {
		return nil, err
	}

	resultUser := toUserModel(userDB)
	return resultUser, nil
}