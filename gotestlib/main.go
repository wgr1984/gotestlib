package gotestlib

import (
	"encoding/json"
	"net/http"
)

const sampleUrl = "https://jsonplaceholder.typicode.com/photos"

type Api struct {
	client *http.Client
}

type Photo struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}

type PhotoWrapper struct {
	items []Photo
}

// var (
// 	items []MyType
// )

func (p *PhotoWrapper) GetItemsCount() int {
	return len(p.items)
}

func (p *PhotoWrapper) GetItem(i int) *Photo {
	if i >= 0 && i < len(p.items) {
		return &p.items[i]
	}
	return nil
}

func NewApi() *Api {
	return &Api{client: http.DefaultClient}
}

func (api *Api) GetPhotos() (*PhotoWrapper, error) {

	resp, err := api.client.Get(sampleUrl)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	defer resp.Body.Close()

	var photos []Photo

	if err := dec.Decode(&photos); err != nil {
		return nil, err
	}

	return &PhotoWrapper{photos}, nil
}
