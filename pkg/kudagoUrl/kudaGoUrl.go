package kudagoUrl

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/goccy/go-json"
)

const MainKudaGoEventURL = `https://kudago.com/public-api/v1.4/events/?expand=place&fields=id,dates,title,images,location,place&order_by=-id&page_size=10`

const KudaGoEventURL = `https://kudago.com/public-api/v1.4/events/`

const KudaGoPlaceUrl = `https://kudago.com/public-api/v1.4/places/`

const KudaGoSearchURL = `https://kudago.com/public-api/v1.4/search/?expand=place`

const twentyKilometers = 20000

const pageSize = 10

type KudaGoUrl struct {
	url string
	httpClient *http.Client
}

func NewKudaGoUrl(kudaGoUrl string, httpClient *http.Client) *KudaGoUrl {
	return &KudaGoUrl{
		url: kudaGoUrl,
		httpClient: httpClient,
	}
}

func (kgUrl *KudaGoUrl) AddPage(page string) {
	kgUrl.url = fmt.Sprintf("%s&%s=%s",kgUrl.url,"page", page)
}

func (kgUrl *KudaGoUrl) AddLongitude(longitude string) {
	kgUrl.url = fmt.Sprintf("%s&%s=%s",kgUrl.url,"lon", longitude)
}

func (kgUrl *KudaGoUrl) AddLatitude(latitude string) {
	kgUrl.url = fmt.Sprintf("%s&%s=%s",kgUrl.url,"lat", latitude)
}

func (kgUrl *KudaGoUrl) AddRadius() {
	kgUrl.url = fmt.Sprintf("%s&%s=%d",kgUrl.url,"radius", twentyKilometers)
}

func (kgUrl *KudaGoUrl) AddActualSince() {
	kgUrl.url = fmt.Sprintf("%s&%s=%d",kgUrl.url,"actual_since", time.Now().Unix())
}

func (kgUrl *KudaGoUrl) AddActualUntil() {
	kgUrl.url = fmt.Sprintf("%s&%s=%d",kgUrl.url,"actual_until",time.Now().Add(time.Hour * 24).Unix())
}

func (kgUrl *KudaGoUrl) GetUrl() (string) {
	return kgUrl.url
}

func (kgUrl *KudaGoUrl) AddPlaceFields() {
	kgUrl.url = fmt.Sprintf("%s?fields=%s", kgUrl.url, "site_url,foreign_url,coords,address,title")
}

func (kgUrl *KudaGoUrl) AddEventFields() {
	kgUrl.url = fmt.Sprintf("%s?fields=%s", kgUrl.url,"id,dates,title,images,location,place,description,price")
}

func (kgUrl *KudaGoUrl) AddEventId(eventId string) {
	kgUrl.url = fmt.Sprintf("%s%s/", kgUrl.url, eventId)
}

func (kgUrl *KudaGoUrl) AddPlaceId(placeId string) {
	kgUrl.url = fmt.Sprintf("%s%s/", kgUrl.url, placeId)
}

func (kgUrl *KudaGoUrl) AddSearchField(searchQuery string) {
	kgUrl.url = fmt.Sprintf("%s&q=%s&ctype=event", kgUrl.url, searchQuery)
}

func (kgUrl *KudaGoUrl) AddPageSize() {
	kgUrl.url = fmt.Sprintf("%s&%s=%d",kgUrl.url,"page_size", pageSize)
}

func (kgUrl *KudaGoUrl) SendKudagoRequestAndParseToStruct(jsonUnmarshalStruct interface{}, errChan chan<- error) {
	resp, err := kgUrl.httpClient.Get(kgUrl.url)
	defer close(errChan)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	if(resp.Header.Get("Content-Type") != "application/json"){
		errChan<-errors.New("not json")
		return
	}

	err = json.NewDecoder(resp.Body).Decode(jsonUnmarshalStruct)
	if err != nil {
		errChan <- err
		return
	}
	errChan <- nil
}