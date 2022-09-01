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
		`const colors = {
'#000000': [0,0,0],
'#FFFFFF': [255,255,255],
'#FF0000': [255,0,0],
'#00FF00': [0,255,0],
'#0000FF': [0,0,255],
'#A8516E': [168,81,110],
'#4682BF': [70,130,191],
'#490206': [73,2,6],
'#A1E9DE': [161,233,222],
'#7B6608': [123,102,8],
};

function timer(tries) {
	var al = {};
	for (let [k, v] of Object.entries(colors)) {
		var i = 0;
    	while (i < tries) {
			var t0 = performance.now();
			start(...v);
			const d = (performance.now() - t0);
			if (d > 0) {
				al[k] = al[k] === undefined ? [] : al[k];
				al[k].push(d);
				i++;
			}
		}
    }
    return al;
}

timer(%d);`, repetition,
	)
}

type vec []float64

type Output struct {
	URL             string         `json:"url"`
	ElapsedColorRaw map[string]vec `json:"elapsed"`
	StatsColorUS    map[string]struct {
		Mean uint16 `json:"mean"`
		SEM  uint16 `json:"sem"`
	} `json:"stats_color_us"`
	StatsUS struct {
		Mean uint16 `json:"mean"`
		SEM  uint16 `json:"sem"`
	} `json:"stats_us"`
}

type stats struct {
	mean float64
	sem  float64
}

func (v vec) mean() float64 {
	var s float64
	for _, el := range v {
		s += el
	}
	return s / float64(len(v))
}

func (v vec) sigma(m float64) float64 {
	var s float64
	for _, el := range v {
		s += (el - m) * (el - m)
	}
	den := float64(len(v))
	return math.Sqrt(s/den/(den-1)) / sigmaCorrectionFactor(v)
}

func (v vec) CalculateStats() stats {
	if len(v) < 2 {
		panic("too few elements")
	}
	m := v.mean()
	return stats{
		mean: m,
		sem:  v.sigma(m),
	}
}

// https://en.wikipedia.org/wiki/Unbiased_estimation_of_standard_deviation#Estimating_the_standard_deviation_of_the_sample_mean
func sigmaCorrectionFactor(v []float64) float64 {
	l := float64(len(v))
	return 1. - 1/4/l - 7/32/l/l - 19/128/l/l/l
}

func (o *Output) Stats() {
	s := map[string]stats{}
	var V vec
	for k, v := range o.ElapsedColorRaw {
		s[k] = v.CalculateStats()
		V = append(V, v...)
	}

	for k, v := range s {
		o.StatsColorUS = map[string]struct {
			Mean uint16 `json:"mean"`
			SEM  uint16 `json:"sem"`
		}{
			k: {Mean: uint16(v.mean * 1000), SEM: uint16(v.sem * 1000)},
		}
	}

	st := V.CalculateStats()
	o.StatsUS.Mean = uint16(st.mean * 1000)
	o.StatsUS.SEM = uint16(st.sem * 1000)
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
		chromedp.EvaluateAsDevTools(timer(reps), &o.ElapsedColorRaw),
	); err != nil {
		return o, err
	}

	o.Stats()

	return o, nil
}
