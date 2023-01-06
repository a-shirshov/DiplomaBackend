package delivery

import (
	"fmt"
	"time"
)

const mainKudaGoEventURL = `https://kudago.com/public-api/v1.4/events/?fields=id,dates,title,images,location,place&order_by=-id&page_size=10`

const KudaGoEventURL = `https://kudago.com/public-api/v1.4/events/`

const KudaGoPlaceUrl = `https://kudago.com/public-api/v1.4/places/`

const twentyKilometers = 20000

type KudaGoUrl struct {
	url string
}

func NewKudaGoUrl(kudaGoUrl string) *KudaGoUrl {
	return &KudaGoUrl{url: kudaGoUrl}
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
	kgUrl.url = fmt.Sprintf("%s?fields=%s", kgUrl.url,"id,dates,title,images,location,place")
}

func (kgUrl *KudaGoUrl) AddEventId(eventId string) {
	kgUrl.url = fmt.Sprintf("%s%s/", kgUrl.url, eventId)
}

func (kgUrl *KudaGoUrl) AddPlaceId(placeId string) {
	kgUrl.url = fmt.Sprintf("%s%s/", kgUrl.url, placeId)
}