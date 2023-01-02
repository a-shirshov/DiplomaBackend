package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	elementsPerPage = 2
)

const (
	GetPlacesQuery = `select id, name, description, about, category, imgUrl from (
		select ROW_NUMBER() OVER() as RowNum, * from "place") as placesPaged 
		where RowNum Between 1 + $1 * ($2 - 1) and $1 * $2`
	GetPlaceQuery = `select * from place where id = $1`
)

type PlaceRepository struct {
	db *sqlx.DB
}

func NewPlaceRepository(db *sqlx.DB) *PlaceRepository {
	return &PlaceRepository{
		db: db,
	}
}

func (pR *PlaceRepository) GetPlaces(page int) ([]*models.Place, error) {
	places := []*models.Place{}
	err := pR.db.Select(&places, GetPlacesQuery, elementsPerPage, &page)
	if err != nil {
		log.Println(err)
		return places, customErrors.ErrPostgres
	}
	return places, nil
}

func (pR *PlaceRepository) GetPlace(id int) (*models.Place, error) {
	place := models.Place{}
	err := pR.db.Get(&place, GetPlaceQuery, &id)
	if err != nil {
		log.Println(err)
		return &place, customErrors.ErrPostgres
	}
	return &place, nil
}