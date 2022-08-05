package repository

import (
	"Diploma/internal/errors"
	"Diploma/internal/models"
	"log"

	"github.com/jackc/pgx"
)

const (
	elementsPerPage = 2
)

type PlaceRepository struct {
	db *pgx.ConnPool
}

func NewPlaceRepository(db *pgx.ConnPool) *PlaceRepository {
	return &PlaceRepository{
		db: db,
	}
}

func (pR *PlaceRepository) GetPlaces(page int) ([]*models.Place, error) {
	rows, err := pR.db.Query("GetPlacesQuery", elementsPerPage, page)
	if err != nil {
		log.Print(err)
		rows.Close()
		return nil, errors.ErrPostgres
	}
	var places []*models.Place

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
			return nil, errors.ErrPostgres
		}

		place := models.ToPlaceModel(placeDB)
		places = append(places, place)
	}
	return places, nil
}

func (pR *PlaceRepository) GetPlace(id int) (*models.Place, error) {
	placeDB := &models.PlaceDB{}
	err := pR.db.QueryRow("GetPlaceQuery",id).Scan(
		&placeDB.ID, 
			&placeDB.Name, 
			&placeDB.Description, 
			&placeDB.About, 
			&placeDB.Category, 
			&placeDB.ImgUrl,
	)
	if err != nil {
		return nil, errors.ErrPostgres
	}

	place := models.ToPlaceModel(placeDB)
	return place, nil
}