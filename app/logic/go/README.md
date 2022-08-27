# Logic in Go

## Optimisation

### Colorname Feature

- Use of the euclidian distance squared value leads to 25% boost  
- RGB encoding as `uint8` instead of `float64` and the Manhattan distance instead of the Euclidean leads to the boost of ~4x

```bash
GOMAXPROCS=1 go test -bench=. -benchmem -memprofile profile .
goos: darwin
goarch: arm64
pkg: edgecomputing/logic
BenchmarkName-10                         2835823               415.2 ns/op             0 B/op          0 allocs/op
BenchmarkNameSquareDistance-10           3545172               335.3 ns/op             0 B/op          0 allocs/op
BenchmarkNameV2-10                      10209282               115.7 ns/op             0 B/op          0 allocs/op
```

```go
package logic

import "math"

type color struct {
	Name    string
	R, G, B float64
}

var colorNameLookupTable = []color{
	{Name: "Aero", R: 124, G: 185, B: 232},
	{Name: "Aero blue", R: 201, G: 255, B: 229},
	{Name: "African violet", R: 178, G: 132, B: 190},
	{Name: "Air Force blue (RAF)", R: 93, G: 138, B: 168},
	{Name: "Black", R: 0, G: 0, B: 0},
	{Name: "Almond", R: 239, G: 222, B: 205},
	{Name: "Amaranth", R: 229, G: 43, B: 80},
	{Name: "Amaranth deep purple", R: 159, G: 43, B: 104},
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

func NameSquareDistance(r, g, b float64) string {
	const distMaxThreshold = 50.
	output := ""
	distance := distMaxThreshold

	for _, color := range colorNameLookupTable {

		d := (color.R-r)*(color.R-r) +
			(color.G-g)*(color.G-g) +
			(color.B-b)*(color.B-b)

		if d == 0 {
			return color.Name
		}

		if d < distance {
			output = color.Name
			distance = d
		}

	}

	return output
}

type colorV2 struct {
	Name    string
	R, G, B uint8
}

var colorNameLookupTableV2 = []colorV2{
	{Name: "Aero", R: 124, G: 185, B: 232},
	{Name: "Aero blue", R: 201, G: 255, B: 229},
	{Name: "African violet", R: 178, G: 132, B: 190},
	{Name: "Air Force blue (RAF)", R: 93, G: 138, B: 168},
	{Name: "Black", R: 0, G: 0, B: 0},
	{Name: "Almond", R: 239, G: 222, B: 205},
	{Name: "Amaranth", R: 229, G: 43, B: 80},
	{Name: "Amaranth deep purple", R: 159, G: 43, B: 104},
}


func dist(x, x0 uint8) uint8 {
	if x > x0 {
		return x - x0
	}
	return x0 - x
}

func NameV2(r, g, b uint8) string {
	output := ""
	const maxDistance = 5
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

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Name(0, 0, 0)
	}
}

func BenchmarkNameSquareDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NameSquareDistance(0, 0, 0)
	}
}

func BenchmarkNameV2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NameV2(0, 0, 0)
	}
}
```


### Colortype Feature

```bash
 GOMAXPROCS=1 go test -bench=. -benchmem -memprofile profile .
goos: darwin
goarch: arm64
pkg: edgecomputing/logic
BenchmarkType            2243859               519.2 ns/op             0 B/op          0 allocs/op
BenchmarkTypeUint8       2475934               486.2 ns/op             0 B/op          0 allocs/op
```

- `uint8` instead of float64` to encode the color and the tree's split threshold leads to 10% speed gain

### Overall improvement

The optimisation listed above led to the factor of 1.5 speed gain.

```bash
GOMAXPROCS=1 go test -bench=. -benchmem -memprofile profile .
goos: darwin
goarch: arm64
pkg: edgecomputing/logic
BenchmarkStartStatusQuo          1318202               906.3 ns/op             0 B/op          0 allocs/op
BenchmarkStart                   1929783               603.5 ns/op             0 B/op          0 allocs/op
```
