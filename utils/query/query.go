package query

const (
	GetPlacesQuery = `select id, name, description, about, category, imgUrl from (
		select ROW_NUMBER() OVER() as RowNum, * from "place") as placesPaged 
		where RowNum Between 1 + $1 * ($2-1) and $1 * $2`
	GetPlaceQuery = `select * from place where id = $1`

	GetEventsQuery = `select id, name, description, about, category, tags, specialInfo from (
		select ROW_NUMBER() OVER (ORDER BY creationDate) as RowNum, * from "event" where place_id = $1) as eventsPaged 
		where RowNum Between 1 + $2 * ($3-1) and $2 * $3`
	GetEventQuery = `select id, name, description, about, category, tags, specialInfo from "event" 
		where id = $1`
)