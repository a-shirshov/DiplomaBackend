package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"Diploma/utils/query"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	elementsPerPage = 2
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
	rows, err := pR.db.Query(query.GetPlacesQuery, elementsPerPage, page)
	if err != nil {
		log.Print(err)
		rows.Close()
		return nil, customErrors.ErrPostgres
	}
	places := []*models.Place{}

	defer rows.Close()

	for rows.Next() {
		placeDB := &models.PlaceDB{}
		err := rows.Scan(
			&placeDB.ID, 
			&placeDB.Name, 
			&placeDB.Description, 
			&placeDB.About, 
			&placeDB.Category, 
			&placeDB.ImgUrl,
		)
		if err != nil {
			return nil, customErrors.ErrPostgres
		}

		place := models.ToPlaceModel(placeDB)
		places = append(places, place)
	}
	return places, nil
}

func (pR *PlaceRepository) GetPlace(id int) (*models.Place, error) {
	placeDB := &models.PlaceDB{}
	err := pR.db.QueryRow(query.GetPlaceQuery, id).Scan(
		&placeDB.ID, 
			&placeDB.Name, 
			&placeDB.Description, 
			&placeDB.About, 
			&placeDB.Category, 
			&placeDB.ImgUrl,
	)
	if err != nil {
		return nil, customErrors.ErrPostgres
	}

	place := models.ToPlaceModel(placeDB)
	return place, nil
}