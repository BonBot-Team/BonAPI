package utils

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"net/http"
)

func DownloadImage(url string) (image.Image, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	buffer := &bytes.Buffer{}

	_, err = io.Copy(buffer, response.Body)

	if err != nil {
		return nil, err
	}

	if http.DetectContentType(buffer.Bytes()) == "image/jpeg" {
		img, err := jpeg.Decode(buffer)

		if err != nil {
			return nil, err
		}

		return img, nil
	} else {
		img, _, err := image.Decode(buffer)

		if err != nil {
			return nil, err
		}

		return img, nil
	}
}
