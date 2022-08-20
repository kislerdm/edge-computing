package logic

import "math"

type color struct {
	Name    string
	R, G, B float64
}

func Name(r, g, b float64) string {
	const distMaxThreshold = 50.
	output := ""
	distance := distMaxThreshold

	for _, color := range colorNameLookupTable {

		dR := color.R - r
		dG := color.G - g
		dB := color.B - b

		if dR == 0 && dG == 0 && dB == 0 {
			return color.Name
		}

		if d := math.Sqrt(dR*dR + dG*dG + dB*dB); d < distance {
			output = color.Name
			distance = d
		}

	}

	return output
}
