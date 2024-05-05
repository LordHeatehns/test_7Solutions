package service

import (
	"errors"

	"github.com/go-resty/resty/v2"
)

func GetText() (string, error) {
	client := resty.New()
	text := ""
	resp, err := client.R().
		SetResult(text).
		Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return text, err
	}

	if resp.StatusCode() != 200 {
		return text, errors.New("internet server error")
	}

	text = string(resp.Body())

	return text, nil

}
