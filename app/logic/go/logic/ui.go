package logic

import (
	"errors"
)

func ui(name string, isWarm bool) (html string) {
	if name == "" {
		name = "Not found"
	}
	t := "Cool"
	if isWarm {
		t = "Warm"
	}

	html = `<div><label for="output_name" id="output_label">Color Name:</label><output name="color_name" id="output_name"> ` +
		name + `</output></div><div><label for="output_type" id="output_label">Color Type:</label><output name="color_type" id="output_type"> ` + t + `</output></div>`

	return
}

func UI(r, g, b float64) (string, error) {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return "", errors.New("wrong RGB input")
	}

	isWarm, err := Type(r, g, b)
	if err != nil {
		return "", err
	}

	return ui(Name(r, g, b), isWarm), nil
}
