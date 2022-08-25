package speedtest

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/chromedp/chromedp"
)

func timer(repetition int) string {
	return fmt.Sprintf(
		`function g(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min) + min);
}

function timer(tries) {
    var al = [], i = 0;
    while (i < tries) {
        var t0 = performance.now();
        start(...[0, 0, 0].map(() => g(1, 255)));
        const d = (performance.now() - t0) * 1000;
		if (d > 0) {
			al.push(d);
        	i++;
		}
    }
    return al;
}

timer(%d);`, repetition,
	)
}

type Output struct {
	URL       string    `json:"url"`
	ElapsedUS []float64 `json:"elapsed_us"`
	MeanUS    uint16    `json:"mean_us"`
	SemUM     uint16    `json:"std_error_of_mean_us"`
}

func (o *Output) Stats() {
	mean := o.mean()
	o.SemUM = uint16(o.sigma(mean))
	o.MeanUS = uint16(mean)
}

func (o *Output) mean() float64 {
	var s float64
	for _, el := range o.ElapsedUS {
		s += el
	}
	return s / float64(len(o.ElapsedUS))
}

func (o *Output) sigma(m float64) float64 {
	var s float64
	for _, el := range o.ElapsedUS {
		s += (el - m) * (el - m)
	}
	den := float64(len(o.ElapsedUS))
	return math.Sqrt(s / den / (den - 1))
}

// Run performs speedtest by emulating a browser and timing the js function execution.
func Run(url string, reps int) (Output, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	o := Output{URL: url}

	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(o.URL),
		chromedp.Sleep(1000*time.Millisecond),
		chromedp.EvaluateAsDevTools(timer(reps), &o.ElapsedUS),
	); err != nil {
		return o, err
	}

	o.Stats()

	return o, nil
}
