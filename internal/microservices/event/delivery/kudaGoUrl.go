package delivery

import (
	"fmt"
	"time"
)

const mainKudaGoEventURL = `https://kudago.com/public-api/v1.4/events/?fields=id,dates,title,images,location&order_by=-id&page_size=10`

const twentyKilometers = 20000

type KudaGoUrl struct {
	url string
}

func NewKudaGoUrl() *KudaGoUrl {
	return &KudaGoUrl{url: mainKudaGoEventURL}
}

func (kgUrl *KudaGoUrl) AddPage(page int) {
	kgUrl.url = fmt.Sprintf("%s&%s=%d",kgUrl.url,"page", page)
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