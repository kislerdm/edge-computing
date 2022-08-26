package logic

import "errors"

type Output struct {
	Name   string
	IsWarm bool
}

func Start(r, g, b float64) (Output, error) {
	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return Output{}, errors.New("wrong RGB input")
	}

	isWarm, err := Type(r, g, b)
	if err != nil {
		return Output{}, err
	}

	return Output{
		Name:   Name(r, g, b),
		IsWarm: isWarm,
	}, nil
}
