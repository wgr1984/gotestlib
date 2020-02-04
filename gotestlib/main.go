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

type PhotosCallback interface {
	SendResult(photos *PhotoWrapper, err error)
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

func (api *Api) GetPhotos(callback PhotosCallback) {
	go func() {
		resp, err := api.client.Get(sampleUrl)
		if err != nil {
			callback.SendResult(nil, err)
			return
		}

		dec := json.NewDecoder(resp.Body)
		defer resp.Body.Close()

		var photos []Photo

		if err := dec.Decode(&photos); err != nil {
			callback.SendResult(nil, err)
			return
		}

		callback.SendResult(&PhotoWrapper{photos}, nil)
	}()
}
