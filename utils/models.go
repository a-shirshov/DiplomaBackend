package utils

import "Diploma/internal/models"

func ToMyEvent(result *models.KudaGoResult) models.MyEvent {
	event := models.MyEvent{}
	event.KudaGoID = result.ID
	event.Title = result.Title
	event.Start = result.Dates[0].Start
	event.End = result.Dates[0].End
	event.Image = result.Images[0].Image
	event.Place = result.Place.ID
	event.Location = result.Location.Slug
	event.Description = result.Description
	event.Price = result.Price
	return event
}

func ToMyEventSearch(result *models.KudaGoSearchResult) models.MyEvent {
	event := models.MyEvent{}
	event.KudaGoID = result.ID
	event.Title = result.Title
	event.Start = result.Daterange.Start
	event.End = result.Daterange.End
	event.Image = result.FirstImage.Image
	event.Place = result.Place.ID
	event.Description = result.Description
	return event
}