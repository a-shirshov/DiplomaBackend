package repository

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/models"
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

const elementsPerPage = 10

const radiusInKilometers = 25

const (
	GetExternalEvents = `SELECT kudago_events_paged.kudago_id, kudago_events_paged.place_id, kudago_events_paged.title, kudago_events_paged.start_time,
    kudago_events_paged.end_time, kudago_events_paged.location, kudago_events_paged.image, kudago_events_paged.description,
    kudago_events_paged.price, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked
	FROM (select ROW_NUMBER() OVER() as RowNum, * from kudago_event) as kudago_events_paged
    LEFT JOIN kudago_favourite ON kudago_events_paged.kudago_id = kudago_favourite.event_id AND kudago_favourite.user_id = $1
		where RowNum Between 1 + $2 * ($3 - 1) and $2 * $3;`

	GetTodayEvents = `SELECT kudago_events_paged.kudago_id, kudago_events_paged.place_id, kudago_events_paged.title, kudago_events_paged.start_time,
    kudago_events_paged.end_time, kudago_events_paged.location, kudago_events_paged.image, kudago_events_paged.description,
    kudago_events_paged.price, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked
	FROM (select ROW_NUMBER() OVER() as RowNum, * from kudago_event where kudago_event.start_time > $1 and kudago_event.end_time < $2) as kudago_events_paged
    LEFT JOIN kudago_favourite ON kudago_events_paged.kudago_id = kudago_favourite.event_id AND kudago_favourite.user_id = $3
		where RowNum Between 1 + $4 * ($5 - 1) and $4 * $5;`
	
	GetCloseEvents = `SELECT events_paged.kudago_id, events_paged.place_id, events_paged.title, events_paged.start_time,
	events_paged.end_time, events_paged.location, events_paged.image, events_paged.description,
	events_paged.price, events_paged.is_liked  from (
		SELECT ROW_NUMBER() OVER() as RowNum, * from (
			SELECT kudago_event.kudago_id, kudago_event.place_id, kudago_event.title, kudago_event.start_time,
			kudago_event.end_time, kudago_event.location, kudago_event.image, kudago_event.description,
			kudago_event.price, 
			CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked, 
			(acos(sin(radians($1)) * sin(radians(lat)) + cos(radians($1)) * cos(radians(lat)) * cos(radians($2 - lon))) * 6371) as distance 
			from kudago_event
			left join kudago_favourite on kudago_event.kudago_id = kudago_favourite.event_id and kudago_favourite.user_id = $3
			join kudago_place on kudago_place.kudago_id = kudago_event.place_id) as close_events 
		where distance <= $4) as events_paged
	where RowNum Between 1 + $5 * ($6 - 1) and $5 * $6;`
	
	GetFullEvent = `select kudago_event.kudago_id, kudago_event.place_id, kudago_event.title, 
	kudago_event.start_time, kudago_event.end_time, kudago_event.location, kudago_event.image,
	kudago_event.description, kudago_event.price, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked, 
	kudago_place.kudago_id, kudago_place.title, kudago_place.address, kudago_place.lat, kudago_place.lon,
	kudago_place.timetable, kudago_place.phone, kudago_place.site_url, kudago_place.foreign_url from kudago_event 
	join kudago_place on kudago_event.place_id = kudago_place.kudago_id
	left join kudago_favourite on kudago_favourite.event_id = kudago_event.kudago_id and kudago_favourite.user_id = $1
	where kudago_event.kudago_id = $2;`

	GetEvent = `select kudago_event.kudago_id, kudago_event.place_id, kudago_event.title, 
	kudago_event.start_time, kudago_event.end_time, kudago_event.location, kudago_event.image,
	kudago_event.description, kudago_event.price, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked
	from kudago_event
	left join kudago_favourite on kudago_favourite.event_id = kudago_event.kudago_id and kudago_favourite.user_id = $1
	where kudago_event.kudago_id = $2;`

	GetRandomEvents = `select kudago_event.kudago_id,kudago_event.place_id, kudago_event.title, 
	kudago_event.start_time, kudago_event.end_time, kudago_event.location, kudago_event.image,
	kudago_event.description, kudago_event.price, kudago_event.vector, kudago_event.vector_title, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked
	from kudago_event
	left join kudago_favourite on kudago_event.kudago_id = kudago_favourite.event_id and kudago_favourite.user_id = $1;`

	GetVector = `select vector from kudago_event where kudago_id = $1;`
	GetVectorTitle = `select vector_title from kudago_event where kudago_id = $1;`

	LikeEvent = `insert into kudago_favourite (user_id, event_id) values ($1, $2);`
	DislikeEvent = `delete from kudago_favourite where user_id = $1 and event_id = $2;`
	CheckLike = `select id from kudago_favourite where user_id = $1 and event_id = $2;`

	GetFavourites = `select liked_by_another_user.kudago_id, liked_by_another_user.place_id, liked_by_another_user.title, liked_by_another_user.start_time, 
    liked_by_another_user.end_time, liked_by_another_user.location, liked_by_another_user.image, liked_by_another_user.description, liked_by_another_user.price, 
    case when kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked from (
		select ROW_NUMBER() OVER() as RowNum, ke.kudago_id, ke.place_id, ke.title, ke.start_time, 
		ke.end_time, ke.location, ke.image, ke.description, ke.price
		from kudago_event as ke
		JOIN kudago_favourite on ke.kudago_id = kudago_favourite.event_id
		and kudago_favourite.user_id = $1
	) as liked_by_another_user
	LEFT JOIN kudago_favourite on kudago_favourite.event_id = liked_by_another_user.kudago_id
	and kudago_favourite.user_id = $2
	where liked_by_another_user.RowNum Between 1 + $3 * ($4 - 1) and $3 * $4;`

	SearchEvents = `select search_result_paged.kudago_id, search_result_paged.place_id, search_result_paged.title, 
		search_result_paged.start_time, search_result_paged.end_time, search_result_paged.location, search_result_paged.image,
		search_result_paged.description, search_result_paged.price, search_result_paged.is_liked from (
		select DISTINCT ROW_NUMBER() OVER() as RowNum, * from (
		  select kudago_event.kudago_id, kudago_event.place_id, kudago_event.title, 
		  kudago_event.start_time, kudago_event.end_time, kudago_event.location, kudago_event.image,
		  kudago_event.description, kudago_event.price, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked
		  from kudago_event
		  left join kudago_favourite on kudago_event.kudago_id = kudago_favourite.event_id and kudago_favourite.user_id = $1 
		  where title ~* $2
	
		  UNION
	
		  select kudago_event.kudago_id, kudago_event.place_id, kudago_event.title, 
		  kudago_event.start_time, kudago_event.end_time, kudago_event.location, kudago_event.image,
		  kudago_event.description, kudago_event.price, CASE WHEN kudago_favourite.event_id IS NULL THEN FALSE ELSE TRUE END AS is_liked
		  from kudago_event
		  left join kudago_favourite on kudago_event.kudago_id = kudago_favourite.event_id and kudago_favourite.user_id = $1 
		  where make_tsvector(title, description) @@ to_tsquery($2)
		) as search_result
	  ) as search_result_paged
	where RowNum Between 1 + $3 * ($4 - 1) and $3 * $4;`
)

type EventRepositoryV2 struct {
	db *sqlx.DB
}

func NewEventRepositoryV2(db *sqlx.DB) *EventRepositoryV2 {
	return &EventRepositoryV2{
		db: db,
	}
}

func (eR *EventRepositoryV2) GetExternalEvents(userID int, page int) (*[]models.MyEvent, error) {
	events := []models.MyEvent{}

	err := eR.db.Select(&events, GetExternalEvents, userID, elementsPerPage, page)
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}
	return &events, nil
}

func (eR *EventRepositoryV2) GetTodayEvents(startTime int64, endTime int64, userID int, page int) (*[]models.MyEvent, error) {
	events := []models.MyEvent{}

	err := eR.db.Select(&events, GetTodayEvents, startTime, endTime, userID, elementsPerPage, page)
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}
	return &events, nil
}

func (eR *EventRepositoryV2) GetCloseEvents(lat string, lon string, userID int, page int) (*[]models.MyEvent, error) {
	events := []models.MyEvent{}

	err := eR.db.Select(&events, GetCloseEvents, lat, lon, userID, radiusInKilometers, elementsPerPage, page)
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}
	return &events, nil
}

func (eR *EventRepositoryV2) GetExternalEvent(userID int, eventID int) (*models.MyFullEvent, error) {
	var event models.MyFullEvent

	err := eR.db.QueryRowx(GetFullEvent, userID, eventID).Scan(
		&event.Event.KudaGoID, &event.Event.Place, &event.Event.Title, &event.Event.Start,
		&event.Event.End, &event.Event.Location, &event.Event.Image, &event.Event.Description,
		&event.Event.Price, &event.Event.IsLiked, &event.Place.KudaGoID, &event.Place.Title,
		&event.Place.Address, &event.Place.Coords.Lat, &event.Place.Coords.Lon, &event.Place.Timetable,
		&event.Place.Phone,&event.Place.SiteURL,&event.Place.ForeignURL,
	)
	if err != nil {
		fmt.Println(err.Error())
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}
	return &event, nil
}

func (eR *EventRepositoryV2) GetEvent(userID int, eventID int) (*models.MyEvent, error) {
	var event *models.MyEvent

	err := eR.db.Get(&event, GetEvent)
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}
	return event, nil
}

func (eU *EventRepositoryV2) GetRandomEvents(userID int) (*[]models.MyEvent, error) {
	events := []models.MyEvent{}
	
	rows, err := eU.db.Queryx(GetRandomEvents, userID)
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}
	defer rows.Close()

	for rows.Next() {
		
		event := models.MyEvent{}
		err := rows.Scan(&event.KudaGoID, &event.Place, &event.Title, &event.Start,
			&event.End, &event.Location, &event.Image, &event.Description,
			&event.Price, pq.Array(&event.Vector), pq.Array(&event.VectorTitle), &event.IsLiked)
		if err != nil {
			log.Println(err.Error())
			return nil, customErrors.ErrPostgres
		}
		events = append(events, event)
	}
	return &events, nil
}

func (eU *EventRepositoryV2) GetVector(eventID int) (*[]float64, error) {
	var vector []float64
	err := eU.db.QueryRowx(GetVector, eventID).Scan(pq.Array(&vector))
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}

	return &vector, nil
}

func (eU *EventRepositoryV2) GetVectorTitle(eventID int) (*[]float64, error) {
	var vector []float64
	err := eU.db.QueryRowx(GetVectorTitle, eventID).Scan(pq.Array(&vector))
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}

	return &vector, nil
}

func (eU *EventRepositoryV2) SwitchLikeEvent(userID, eventID int) (error) {
	var id int
	err := eU.db.Get(&id, CheckLike, userID, eventID)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Println(err.Error())
			return customErrors.ErrPostgres
		}

		_, err := eU.db.Exec(LikeEvent, userID, eventID)
		if err != nil {
			log.Println(err.Error())
			return customErrors.ErrPostgres
		}
		
		return nil
	}

	_, err = eU.db.Exec(DislikeEvent, userID, eventID)
	if err != nil {
		log.Println(err.Error())
		return customErrors.ErrPostgres
	}

	return nil
}

func (eU *EventRepositoryV2) GetFavourites(userID, checkedUserID, page int) (*[]models.MyEvent, error) {
	events := []models.MyEvent{}

	err := eU.db.Select(&events, GetFavourites, checkedUserID, userID, elementsPerPage, page)
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}
	return &events, nil	
}

func (eU *EventRepositoryV2) SearchEvents(userID int, searchingEvent string, page int) (*[]models.MyEvent, error) {
	events := []models.MyEvent{}
	err := eU.db.Select(&events, SearchEvents, userID, searchingEvent, elementsPerPage, page)
	if err != nil {
		log.Println(err.Error())
		return nil, customErrors.ErrPostgres
	}
	return &events, nil	
}