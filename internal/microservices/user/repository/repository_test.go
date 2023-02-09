package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type getUserTest struct {
	name	    		string
	inputUserID 		int
	beforeTest  		func(sqlmock.Sqlmock)
	expectedUser 		*models.User
	expectedError 		error
}

var getUserTests = []getUserTest{
	{
		"Succesfully get user from database",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"id", "name", "surname", "about", "img_url"}
			rows := mockSQL.NewRows(columns).
				AddRow(1, "Artyom", "Shirshov", "author", "uuid")

			mockSQL.ExpectQuery(regexp.QuoteMeta(GetUserQuery)).
				WithArgs(1).
				RowsWillBeClosed().
				WillReturnRows(rows)
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			About: "author",
			ImgUrl: "uuid",
		},
		nil,
	},
	{
		"User not found",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(GetUserQuery)).
				WithArgs(1).
				WillReturnError(sql.ErrNoRows)
		},
		&models.User{},
		customErrors.ErrUserNotFound,
	},
	{
		"Database is not working",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(GetUserQuery)).
				WithArgs(1).
				WillReturnError(errors.New("something went wrong"))
		},
		&models.User{},
		customErrors.ErrPostgres,
	},
}

type updateUserTest struct {
	name			string
	inputUser 		*models.User
	beforeTest  	func(sqlmock.Sqlmock)
	expectedUser 	*models.User
	expectedError 	error
}

var updateUserTests = []updateUserTest {
	{
		"Successfully update user",
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			About: "About",
		},
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"id", "name", "surname", "about", "img_url"}
			rows := mockSQL.NewRows(columns).
				AddRow(1, "Artyom", "Shirshov", "About", "ImgUUID")

			mockSQL.ExpectQuery(regexp.QuoteMeta(UpdateUserQuery)).
				WithArgs("Artyom", "Shirshov", "About", 1).
				RowsWillBeClosed().
				WillReturnRows(rows)
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			About: "About",
			ImgUrl: "ImgUUID",
		},
		nil,
	},
}

type updateUserImageTest struct {
	name			string
	inputUserID		int
	inputImgUUID 	string
	beforeTest  	func(sqlmock.Sqlmock)
	expectedUser 	*models.User
	expectedError 	error
}

var updateUserImageTests = []updateUserImageTest {
	{
		"Successfully update user image",
		1,
		"ImgUUID",
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"id", "name", "surname", "about", "img_url"}
			rows := mockSQL.NewRows(columns).
				AddRow(1, "Artyom", "Shirshov", "About", "ImgUUID")

			mockSQL.ExpectQuery(regexp.QuoteMeta(UpdateUserQuery)).
				WithArgs("ImgUUID", 1).
				RowsWillBeClosed().
				WillReturnRows(rows)
		},
		&models.User{
			ID: 1,
			Name: "Artyom",
			Surname: "Shirshov",
			About: "About",
			ImgUrl: "ImgUUID",
		},
		nil,
	},
}

type getFavouriteKudagoEventsIDsTest struct {
	name				string
	inputUserID			int
	beforeTest  		func(sqlmock.Sqlmock)
	expectedEventIDs 	[]int
	expectedError 		error
}

var getFavouriteKudagoEventsIDsTests = []getFavouriteKudagoEventsIDsTest {
	{
		"Successfully get favourite kudago events ids",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			columns := []string{"event_id"}
			rows := mockSQL.NewRows(columns).
				AddRow(1).AddRow(2).AddRow(3)

			mockSQL.ExpectQuery(regexp.QuoteMeta(GetFavouriteEventsID)).
				WithArgs(1).
				RowsWillBeClosed().
				WillReturnRows(rows)
		},
		[]int{1, 2, 3},
		nil,
	},
	{
		"User doesn't have favourite kudago event ids",
		1,
		func(mockSQL sqlmock.Sqlmock) {
			mockSQL.ExpectQuery(regexp.QuoteMeta(GetFavouriteEventsID)).
				WithArgs(1).
				WillReturnError(sql.ErrNoRows)
		},
		[]int{},
		nil,
	},
}

func prepareTestEnvironment() (*UserRepository, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repositoryTest := NewUserRepository(sqlxDB)

	return repositoryTest, mock, nil
}

func TestGetUser(t *testing.T) {
	for _, test := range getUserTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualUser, actualErr := repositoryTest.GetUser(test.inputUserID)
			assert.Equal(t, test.expectedUser, actualUser)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})	
	}
}

func TestUpdateUser(t *testing.T) {
	for _, test := range updateUserTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualUser, actualErr := repositoryTest.UpdateUser(test.inputUser)
			assert.Equal(t, test.expectedUser, actualUser)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})	
	}
}

func TestUpdateUserImage(t *testing.T) {
	for _, test := range updateUserImageTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualUser, actualErr := repositoryTest.UpdateUserImage(test.inputUserID, test.inputImgUUID)
			assert.Equal(t, test.expectedUser, actualUser)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})	
	}
}

func TestGetFavouriteKudagoEventsIDs(t *testing.T) {
	for _, test := range getFavouriteKudagoEventsIDsTests {
		t.Run(test.name, func(t *testing.T) {
			repositoryTest, mock, err := prepareTestEnvironment()
			if err != nil {
				log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer repositoryTest.db.Close()

			if test.beforeTest != nil {
				test.beforeTest(mock)
			}

			actualIDs, actualErr := repositoryTest.GetFavouriteKudagoEventsIDs(test.inputUserID)
			assert.Equal(t, test.expectedEventIDs, actualIDs)
			assert.Equal(t, test.expectedError, actualErr)

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Not all expectations: %s", err)
			}
		})	
	}
}
