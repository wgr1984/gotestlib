package main

import (
	"fmt"
	"gotestlib/gotestlib"
)

func main() {

	api := gotestlib.NewApi()

	photos, err := api.GetPhotos()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(photos)
}
