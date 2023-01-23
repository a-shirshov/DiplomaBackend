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