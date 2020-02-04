package main

import (
	"fmt"
	"gotestlib/gotestlib"
	"time"
)

type Callback struct{}

func (c *Callback) SendResult(photos *gotestlib.PhotoWrapper, err error) {
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(photos)
}

func main() {

	api := gotestlib.NewApi()

	api.GetPhotos(&Callback{})

	time.Sleep(10 * time.Second)
}
