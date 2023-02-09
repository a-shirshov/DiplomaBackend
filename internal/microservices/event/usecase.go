package event

type Usecase interface {
	GetPeopleCountWithEventCreatedIfNecessary(eventID int) (int, error)
	CheckKudaGoMeeting(userID, eventID int) (bool, error)
	CheckKudaGoFavourite(userID, eventID int) (bool, error)
	SwitchEventMeeting(userID int, eventID int) (error)
	SwitchEventFavourite(userID int, eventID int) error
}