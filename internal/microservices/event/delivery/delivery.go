package delivery

import (
	"Diploma/internal/customErrors"
	"Diploma/internal/microservices/event"
	"Diploma/internal/models"
	"Diploma/pkg/kudagoUrl"
	"Diploma/utils"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type EventDelivery struct {
	eventUsecase event.Usecase
}

func NewEventDelivery(eventU event.Usecase) *EventDelivery {
	return &EventDelivery{
		eventUsecase: eventU,
	}
}

func (eD *EventDelivery) GetExternalEvents(c *gin.Context) {
	kudaGoURL := kudagoUrl.NewKudaGoUrl(kudagoUrl.MainKudaGoEventURL, &http.Client{Timeout: 10 * time.Second})
	page := utils.GetPageQueryParamFromRequest(c)
	kudaGoURL.AddPage(page)

	kudaGoEvents := &models.KudaGoEvents{}
	eventErr := make(chan error, 1)
	go kudaGoURL.SendKudagoRequestAndParseToStruct(kudaGoEvents, eventErr)

	var userID int
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	if <-eventErr != nil {
		utils.SendMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		myEventResult, err := eD.convertResultToReadyToBeGivenEvent(userID, &result)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		events.Events = append(events.Events, *myEventResult)
	}
	c.JSON(http.StatusOK, events)
}

func (eD *EventDelivery) GetCloseExternalEvents(c *gin.Context) {
	kudaGoURL := kudagoUrl.NewKudaGoUrl(kudagoUrl.MainKudaGoEventURL, &http.Client{Timeout: 10 * time.Second})
	page := utils.GetPageQueryParamFromRequest(c)
	kudaGoURL.AddPage(page)

	longitude := getLongitudeQueryParamFromRequest(c)
	kudaGoURL.AddLongitude(longitude)

	latitude := getLatitudeQueryParamFromRequest(c)
	kudaGoURL.AddLatitude(latitude)

	kudaGoURL.AddRadius()

	kudaGoEvents := &models.KudaGoEvents{}
	eventErr := make(chan error, 1)
	go kudaGoURL.SendKudagoRequestAndParseToStruct(kudaGoEvents, eventErr)

	var userID int
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	if <-eventErr != nil {
		utils.SendMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		myEventResult, err := eD.convertResultToReadyToBeGivenEvent(userID, &result)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		events.Events = append(events.Events, *myEventResult)
	}
	c.JSON(http.StatusOK, events)
}

func getLongitudeQueryParamFromRequest(c *gin.Context) string {
	longitudeParam := c.DefaultQuery("lon", "37.6155600")
	return longitudeParam
}

func getLatitudeQueryParamFromRequest(c *gin.Context) string {
	latitudeParam := c.DefaultQuery("lat", "55.7522200")
	return latitudeParam
}

func (eD *EventDelivery) GetTodayEvents(c *gin.Context) {
	kudaGoURL := kudagoUrl.NewKudaGoUrl(kudagoUrl.MainKudaGoEventURL, &http.Client{Timeout: 10 * time.Second})
	page := utils.GetPageQueryParamFromRequest(c)
	kudaGoURL.AddPage(page)
	kudaGoURL.AddActualSince()
	kudaGoURL.AddActualUntil()

	kudaGoEvents := &models.KudaGoEvents{}
	eventErr := make(chan error, 1)
	go kudaGoURL.SendKudagoRequestAndParseToStruct(kudaGoEvents, eventErr)
	
	var userID int
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	if <-eventErr != nil {
		utils.SendMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		myEventResult, err := eD.convertResultToReadyToBeGivenEvent(userID, &result)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		events.Events = append(events.Events, *myEventResult)
	}
	c.JSON(http.StatusOK, events)
}

func (eD *EventDelivery) convertResultToReadyToBeGivenEvent(userID int, result *models.KudaGoResult) (*models.MyEvent, error) {
	if result.Place.IsStub {
		return nil, errors.New("stub event")
	}
	myEventResult := utils.ToMyEvent(result)
	IsLiked, err := eD.eventUsecase.CheckKudaGoFavourite(userID, myEventResult.KudaGoID)
	if err != nil {
		log.Println("Liked error:", err.Error())
	}
	myEventResult.IsLiked = IsLiked
	return &myEventResult, nil
}

func (eD *EventDelivery) GetExternalEvent(c *gin.Context) {
	var userID int
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}

	KudaGoEventUrl := kudagoUrl.NewKudaGoUrl(kudagoUrl.KudaGoEventURL, &http.Client{Timeout: 10 * time.Second})
	KudaGoPlaceUrl := kudagoUrl.NewKudaGoUrl(kudagoUrl.KudaGoPlaceUrl, &http.Client{Timeout: 10 * time.Second})

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}
	KudaGoEventUrl.AddEventId(eventIDStr)
	KudaGoEventUrl.AddEventFields()

	placeID := c.Param("place_id")
	KudaGoPlaceUrl.AddPlaceId(placeID)
	KudaGoPlaceUrl.AddPlaceFields()

	KudaGoEvent := &models.KudaGoResult{}
	KudaGoPlace := &models.KudaGoPlaceResult{}

	eventErr := make(chan error, 1)
	placeErr := make(chan error, 1)
	go KudaGoEventUrl.SendKudagoRequestAndParseToStruct(KudaGoEvent, eventErr)
	go KudaGoPlaceUrl.SendKudagoRequestAndParseToStruct(KudaGoPlace, placeErr)
	peopleCount, err := eD.eventUsecase.GetPeopleCountWithEventCreatedIfNecessary(eventID)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, customErrors.ErrPostgres.Error())
		return
	}

	isGoing := false
	isFavourite := false
	var meetingErr error
	var favouriteErr error

	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		defer wg.Done()
		isGoing, meetingErr = eD.eventUsecase.CheckKudaGoMeeting(userID, eventID)
		if meetingErr != nil {
			return
		}
	}()

	go func(){
		defer wg.Done()
		isFavourite, favouriteErr = eD.eventUsecase.CheckKudaGoFavourite(userID, eventID)
		if favouriteErr != nil {
			return
		}
	}()

	wg.Wait()

	if <-eventErr != nil {
		utils.SendMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	if <-placeErr != nil {
		utils.SendMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	event := utils.ToMyEvent(KudaGoEvent)
	eventAndPlace := &models.KudaGoPlaceAndEvent{
		Event:       event,
		Place:       *KudaGoPlace,
		PeopleCount: peopleCount,
		IsGoing:     isGoing,
		IsFavourite: isFavourite,
	}
	c.JSON(http.StatusOK, eventAndPlace)
}

func (eD *EventDelivery) SwitchEventMeeting(c *gin.Context) {
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendMessage(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	err = eD.eventUsecase.SwitchEventMeeting(au.UserId, eventID)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendMessage(c, http.StatusOK, "OK")
}

func (eD *EventDelivery) SwitchEventFavourite(c *gin.Context) {
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		utils.SendMessage(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	eventIDStr := c.Param("event_id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	KudaGoEventUrl := kudagoUrl.NewKudaGoUrl(kudagoUrl.KudaGoEventURL, &http.Client{Timeout: 10 * time.Second})
	KudaGoEventUrl.AddEventId(eventIDStr)
	KudaGoEventUrl.AddEventFields()

	KudaGoEvent := &models.KudaGoResult{}
	eventErr := make(chan error, 1)
	
	go KudaGoEventUrl.SendKudagoRequestAndParseToStruct(KudaGoEvent, eventErr)

	err = eD.eventUsecase.SwitchEventFavourite(au.UserId, eventID)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	if <-eventErr != nil {
		utils.SendMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	resultEvent, err := eD.convertResultToReadyToBeGivenEvent(au.UserId, KudaGoEvent)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resultEvent)
}

func (eD *EventDelivery) GetFavourites(c *gin.Context) {
	var requestUserID int
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		requestUserID = 0
	} else {
		requestUserID = au.UserId
	}


	userIDString := c.Param("user_id")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		utils.SendMessage(c, http.StatusBadRequest, err.Error())
		return
	}

	favouriteEvents := models.MyEvents{}
	favouriteEvents.Events = []models.MyEvent{}

	favouriteEventIDs, err := eD.eventUsecase.GetFavouriteKudagoEventsIDs(userID)
	log.Println(favouriteEventIDs)
	if err != nil {
		utils.SendMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	var wg sync.WaitGroup
	httpClient :=  &http.Client{Timeout: 10 * time.Second}
	for index, eventID := range favouriteEventIDs {
		wg.Add(1)
		go func(index int, eventID int) {	
			defer wg.Done()
			kudagoEvent := &models.KudaGoResult{}
			eventErr := make(chan error, 1)
			kudaGoURL := kudagoUrl.NewKudaGoUrl(kudagoUrl.KudaGoEventURL, httpClient)
			kudaGoURL.AddEventId(fmt.Sprintf("%d", eventID))
			kudaGoURL.AddEventFields()
			kudaGoURL.SendKudagoRequestAndParseToStruct(kudagoEvent, eventErr)
			if <-eventErr != nil {
				return
			}
			myEvent, err := eD.convertResultToReadyToBeGivenEvent(requestUserID, kudagoEvent)
			if err != nil {
				log.Println(err.Error())
				return
			}
			favouriteEvents.Events = append(favouriteEvents.Events, *myEvent)
		}(index, eventID)
	}
	wg.Wait()
	c.JSON(http.StatusOK, favouriteEvents)
}

func (eD *EventDelivery) SearchKudaGoEvent(c *gin.Context) {
	searchingEvent := c.Query("q")
	page := utils.GetPageQueryParamFromRequest(c)

	httpClient :=  &http.Client{Timeout: 10 * time.Second}
	kudaGoURL := kudagoUrl.NewKudaGoUrl(kudagoUrl.KudaGoSearchURL, httpClient)
	kudaGoURL.AddSearchField(searchingEvent)
	kudaGoURL.AddPage(page)
	kudaGoURL.AddPageSize()
	log.Println(kudaGoURL.GetUrl())

	kudaGoEvents := &models.KudaGoSearchResults{}
	eventErr := make(chan error, 1)
	go kudaGoURL.SendKudagoRequestAndParseToStruct(kudaGoEvents, eventErr)
	if <-eventErr != nil {
		utils.SendMessage(c, http.StatusMisdirectedRequest, "kudago error")
		return
	}

	var userID int
	au, err := utils.GetAUFromContext(c)
	if err != nil {
		userID = 0
	} else {
		userID = au.UserId
	}
	
	events := &models.MyEvents{}
	for _, result := range kudaGoEvents.Results {
		myEventResult, err := eD.convertSearchResultToReadyToBeGivenEvent(userID, &result)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		events.Events = append(events.Events, *myEventResult)
	}
	c.JSON(http.StatusOK, events)
}

func (eD *EventDelivery) convertSearchResultToReadyToBeGivenEvent(userID int, result *models.KudaGoSearchResult) (*models.MyEvent, error) {
	if result.Place.IsStub {
		return nil, errors.New("stub event")
	}
	myEventResult := utils.ToMyEventSearch(result)
	IsLiked, _ := eD.eventUsecase.CheckKudaGoFavourite(userID, myEventResult.KudaGoID)
	myEventResult.IsLiked = IsLiked
	return &myEventResult, nil
}
