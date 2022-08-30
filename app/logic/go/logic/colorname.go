package logic

type color struct {
	Name    string
	R, G, B uint8
}

func dist(x, x0 uint8) uint8 {
	if x > x0 {
		return x - x0
	}
	return x0 - x
}

func Name(r, g, b uint8) string {
	output := ""
	const maxDistance = 10
	maxDistanceTotal := uint8(3 * maxDistance)

	for _, color := range colorNameLookupTable {
		var totalDist uint8
		d := dist(color.R, r)
		if d > maxDistance {
			continue
		}
		totalDist += d

		d = dist(color.G, g)
		if d > maxDistance {
			continue
		}
		totalDist += d

		d = dist(color.B, b)
		if d > maxDistance {
			continue
		}
		totalDist += d

		if totalDist == 0 {
			return color.Name
		}
		if totalDist < maxDistanceTotal {
			output = color.Name
			maxDistanceTotal = totalDist
		}
	}

	return output
}
