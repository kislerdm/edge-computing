package logic

type Output struct {
	Name   string
	IsWarm bool
}

func Start(r, g, b uint8) (Output, error) {
	isWarm, err := Type(r, g, b)
	if err != nil {
		return Output{}, err
	}

	return Output{
		Name:   Name(r, g, b),
		IsWarm: isWarm,
	}, nil
}
