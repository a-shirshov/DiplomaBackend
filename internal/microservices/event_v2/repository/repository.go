package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"

	"github.com/jmoiron/sqlx"
)

const (
	GetExternalEvents = `select kudago_id, title, start, end, location, image, place, description, price from recomendation_events limit 10;`
)

type EventRepositoryV2 struct {
	db *sqlx.DB
}

func NewEventRepositoryV2(db *sqlx.DB) *EventRepositoryV2 {
	return &EventRepositoryV2{
		db: db,
	}
}

func (eR *EventRepositoryV2) GetExternalEvents(userID int) (*[]models.MyEvent, error) {
	var events []models.MyEvent

	err := eR.db.Select(&events, GetExternalEvents)
	if err != nil {
		return nil, customErrors.ErrPostgres
	}
	return &events, nil
}