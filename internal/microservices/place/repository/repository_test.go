package repository

import (
	"Diploma/internal/models"
	"Diploma/utils"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type getPlacesTest struct {
	page         int
	outputPlaces []*models.Place
	outputErr    error
}

var getPlacesTests = []getPlacesTest{
	{
		1, []*models.Place{
			{
				ID:          1,
				Name:        "Place_1",
				Description: "Description_1",
				About:       "About_1",
				Category:    "Category_1",
				ImgUrl:      "ImgUrl_1",
			},
			{
				ID:          2,
				Name:        "Place_2",
				Description: "Description_2",
				About:       "About_2",
				Category:    "Category_2",
				ImgUrl:      "ImgUrl_2",
			},
		}, nil,
	},
}

type getPlaceTest struct {
	id int
	outputPlace *models.Place
	outputErr    error
}

var getPlaceTests = []getPlaceTest{
	{
		1, &models.Place{
			ID:          1,
			Name:        "Place_1",
			Description: "Description_1",
			About:       "About_1",
			Category:    "Category_1",
			ImgUrl:      "ImgUrl_1",
		}, nil,
	},
}

func TestGetPlaces(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db,"sqlmock")
	repositoryTest := NewPlaceRepository(sqlxDB)

	columns := []string{"id", "name", "description", "about", "category", "imgUrl"}

	for _, test := range getPlacesTests {
		rows := mock.NewRows(columns).
			AddRow(
				test.outputPlaces[0].ID,
				test.outputPlaces[0].Name,
				test.outputPlaces[0].Description,
				test.outputPlaces[0].About,
				test.outputPlaces[0].Category,
				test.outputPlaces[0].ImgUrl).
			AddRow(
				test.outputPlaces[1].ID,
				test.outputPlaces[1].Name,
				test.outputPlaces[1].Description,
				test.outputPlaces[1].About,
				test.outputPlaces[1].Category,
				test.outputPlaces[1].ImgUrl)

		mock.ExpectQuery(utils.GetPlacesQuery).
			WithArgs(elementsPerPage, test.page).
			RowsWillBeClosed().WillReturnRows(rows)

		out, dbErr := repositoryTest.GetPlaces(test.page)
		assert.Equal(t, test.outputPlaces, out)
		assert.Nil(t, dbErr)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations: %s", err)
		}
	}

}

func TestGetPlace(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db,"sqlmock")
	repositoryTest := NewPlaceRepository(sqlxDB)

	columns := []string{"id", "name", "description", "about", "category", "imgUrl"}

	for _, test := range getPlaceTests {
		rows := mock.NewRows(columns).
			AddRow(
				test.outputPlace.ID,
				test.outputPlace.Name,
				test.outputPlace.Description,
				test.outputPlace.About,
				test.outputPlace.Category,
				test.outputPlace.ImgUrl)

		mock.ExpectQuery(utils.GetPlaceQuery).
			WithArgs(test.id).
			WillReturnRows(rows)

		out, dbErr := repositoryTest.GetPlace(test.id)
		assert.Equal(t, test.outputPlace, out)
		assert.Equal(t, dbErr, nil)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Not all expectations: %s", err)
		}
	}

}