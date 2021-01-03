package utils

import (
	"image"
	"net/http"
)

func DownloadImage(url string) (image.Image, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	img, _,  err := image.Decode(response.Body)

	if err != nil {
		return nil, err
	}

	return img, nil
}