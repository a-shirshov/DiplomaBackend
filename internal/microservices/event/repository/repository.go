package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"database/sql"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	elementsPerPage = 2
)

const (
	GetEventsQuery = `select id, name, description, about, category, tags, specialInfo from (
		select ROW_NUMBER() OVER (ORDER BY creationDate) as RowNum, * from "event" where place_id = $1) as eventsPaged 
		where RowNum Between 1 + $2 * ($3-1) and $2 * $3;`
	GetEventQuery = `select id, name, description, about, category, tags, specialInfo from "event" 
		where id = $1;`
	
	GetPeopleCount = `select people_count from "kudago_event" where event_id = $1;`
	CreateKudaGoEvent = `insert into "kudago_event" (event_id) values ($1);`

	CreateKudaGoMeeting = `insert into "kudago_meeting" (user_id, event_id) values ($1, $2);`
	CheckKudaGoMeeting = `select id from "kudago_meeting" where user_id = $1 and event_id = $2;`
	DeleteKudaGoMeeting = `delete from "kudago_meeting" where id = $1;`

	CheckEventFavourite = `select id from "kudago_favourite" where user_id = $1 and event_id = $2;`
	AddEventToFavourite = `insert into "kudago_favourite" (user_id, event_id) values ($1, $2);`
	DeleteEventFromFavoutire = `delete from "kudago_favourite" where id = $1;`
)

type EventRepository struct {
	db *sqlx.DB
}

func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

func (eR *EventRepository) GetEvents(page int) ([]*models.Event, error) {
	events := []*models.Event{}
	err := eR.db.Select(&events,GetEventsQuery, elementsPerPage, page)
	if err != nil {
		log.Println(err)
		return events, customErrors.ErrPostgres
	}
	return events, nil
}

func (eR *EventRepository) GetEvent(eventId int) (*models.Event, error) {
	event := &models.Event{}
	err := eR.db.Get(&event, GetEventQuery, eventId)
	if err != nil {
		log.Println(err)
		return event, customErrors.ErrPostgres
	}
	return event, nil
}

func (eR *EventRepository) GetPeopleCount(eventID int) (int, error) {
	var peopleCount int
	err := eR.db.Get(&peopleCount, GetPeopleCount, &eventID)
	if err != nil {
		return 0, err
	}
	return peopleCount, nil
}

func (eR *EventRepository) CreateKudaGoEvent(eventID int) (error) {
	_, err := eR.db.Exec(CreateKudaGoEvent, &eventID)
	return err
}

func (eR *EventRepository) SwitchEventMeeting(userID int, eventID int) (error) {
	var meetingID int
	err := eR.db.Get(&meetingID, CheckKudaGoMeeting, &userID, &eventID)
	if err != nil {
		log.Println(err.Error())
		if err != sql.ErrNoRows {
			log.Println(err.Error())
			return customErrors.ErrPostgres
		}

		_, err = eR.db.Exec(CreateKudaGoMeeting, &userID, &eventID)
		if err != nil {
			log.Println(err.Error())
			return customErrors.ErrPostgres
		}
		return nil
	}

	_, err = eR.db.Exec(DeleteKudaGoMeeting, &meetingID)
	if err != nil {
		log.Println(err.Error())
		return customErrors.ErrPostgres
	}
	return nil
}

func (eR *EventRepository) CheckKudaGoMeeting(userID int, eventID int) (bool, error) {
	var meetingID int
	isGoing := false
	err := eR.db.Get(&meetingID, CheckKudaGoMeeting, &userID, &eventID)
	if err != nil {
		if err != sql.ErrNoRows {
			return isGoing, customErrors.ErrPostgres
		}
		return isGoing, nil
	}
	isGoing = true
	return isGoing, nil
}

func (eR *EventRepository) SwitchEventFavourite(userID int, eventID int) error {
	var favouriteID int
	err := eR.db.Get(&favouriteID, CheckEventFavourite, &userID, &eventID)
	if err != nil {
		if err != sql.ErrNoRows {
			return err
		}
		_, err := eR.db.Exec(AddEventToFavourite, &userID, &eventID)
		if err != nil {
			return customErrors.ErrPostgres
		}
		return nil
	}

	_, err = eR.db.Exec(DeleteEventFromFavoutire, &favouriteID)
	if err != nil {
		return customErrors.ErrPostgres
	}
	return nil
}

func (eR *EventRepository) CheckKudaGoFavourite(userID int, eventID int) (bool, error) {
	var favouriteID int
	isFavourite := false
	err := eR.db.Get(&favouriteID, CheckEventFavourite, &userID, &eventID)
	if err != nil {
		if err != sql.ErrNoRows {
			return isFavourite, customErrors.ErrPostgres
		}
		return isFavourite, nil
	}
	isFavourite = true
	return isFavourite, nil
}