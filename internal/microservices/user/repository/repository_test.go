package repository

import (
	"Diploma/internal/models"
	"Diploma/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type getUserTest struct {
	userId int
	outputUser *models.User
	outputErr error
}

var getUserTests = []getUserTest{
	{
		1, &models.User{
			ID: 1,
			Name: "Name_1",
			Surname: "Surname_1",
			Email: "Email_1",
			About: "About_1",
			ImgUrl: "ImgUrl_1",
		}, nil,
	},
}

type getUserByEmailTest struct {
	userEmail string
	outputUser *models.User
	outputErr error
}

var getUserByEmailTests = []getUserByEmailTest {
	{
		"Email_1", &models.User{
			ID: 1,
			Name: "Name_1",
			Surname: "Surname_1",
			Email: "Email_1",
			Password: "Password_1",
			About: "About_1",
			ImgUrl: "ImgUrl_1",
		}, nil,
	},
}

type updateUserTest struct {
	userId int
	inputUser *models.User
	outputUser *models.User
	outputErr error
}

var updateUserTests = []updateUserTest{
	{
		1, &models.User{
			ID: 1,
			Name: "Name_1",
			Surname: "Surname_1",
			About: "About_1",
			ImgUrl: "ImgUrl_1",
		},
		&models.User{
			ID: 1,
			Name: "Name_1",
			Surname: "Surname_1",
			Email: "Email_1",
			About: "About_1",
			ImgUrl: "ImgUrl_1",
		}, nil,
	},
}

func TestGetUser(t *testing.T){
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db,"sqlmock")
	repositoryTest := NewUserRepository(sqlxDB)

	columns := []string{"id","name","surname","email", "about","imgurl"}
	
	for _, test := range getUserTests {
		rows := mock.NewRows(columns).
			AddRow(
				test.outputUser.ID,
				test.outputUser.Name,
				test.outputUser.Surname,
				test.outputUser.Email,
				test.outputUser.About,
				test.outputUser.ImgUrl)
		mock.ExpectQuery(utils.GetUserQuery).
		WithArgs(test.userId).
		RowsWillBeClosed().WillReturnRows(rows)

		out, dbErr := repositoryTest.GetUser(test.userId)
		assert.Equal(t, test.outputUser, out)
		assert.Nil(t,dbErr)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations: %s", err)
		}
	}
}

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db,"sqlmock")
	repositoryTest := NewUserRepository(sqlxDB)

	columns := []string{"id", "name", "surname", "email", "password", "about", "imgurl"}

	for _, test := range getUserByEmailTests {
		rows := mock.NewRows(columns).
			AddRow(
				test.outputUser.ID,
				test.outputUser.Name,
				test.outputUser.Surname,
				test.outputUser.Email,
				test.outputUser.Password,
				test.outputUser.About,
				test.outputUser.ImgUrl,
			)
		
		mock.ExpectQuery(utils.GetUserByEmailQuery).
			WithArgs(test.userEmail).
			RowsWillBeClosed().WillReturnRows(rows)
			
		out, dbErr := repositoryTest.GetUserByEmail(test.userEmail)
		assert.Equal(t, test.outputUser, out)
		assert.Nil(t, dbErr)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations: %s", err)
		}
	}
}

func TestUpdateUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db,"sqlmock")
	repositoryTest := NewUserRepository(sqlxDB)
	columns := []string{"id", "name", "surname", "email", "about", "imgurl"}
	
	for _, test := range updateUserTests {
		rows := mock.NewRows(columns).
			AddRow(test.outputUser.ID,
				test.outputUser.Name,
				test.outputUser.Surname,
				test.outputUser.Email,
				test.outputUser.About,
				test.outputUser.ImgUrl)
		
		mock.ExpectQuery(utils.UpdateUserQuery).
			WithArgs(
				test.inputUser.Name, 
				test.inputUser.Surname, 
				test.inputUser.About,
				test.inputUser.ImgUrl,
				test.userId).
			RowsWillBeClosed().WillReturnRows(rows)
		
		out, dbErr := repositoryTest.UpdateUser(test.userId, test.inputUser)
		assert.Equal(t, test.outputUser, out)
		assert.Nil(t, dbErr)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations: %s", err)
		}
	}
}