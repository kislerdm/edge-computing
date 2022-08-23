package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

func timer(repetition int) string {
	return fmt.Sprintf(
		`function g(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min) + min);
};

function timer(tries) {
    var al = [], i = 0;
    while (i < tries) {
        var t0 = performance.now();
        start(...[0, 0, 0].map(() => g(1, 255)));
        al.push((performance.now() - t0) * 1000);
        i++;
    }
    return al;
};

timer(%d);`, repetition,
	)
}

type output struct {
	URL       string    `json:"url"`
	ElapsedUS []float64 `json:"elapsed_us"`
	MeanUS    float64   `json:"mean_us"`
}

func (o *output) mean() {
	var s float64
	for _, el := range o.ElapsedUS {
		s += el
	}
	o.MeanUS = s / float64(len(o.ElapsedUS))
}

func main() {
	var p string
	var reps int
	var o output

	flag.StringVar(&o.URL, "url", "", "url to call")
	flag.StringVar(&p, "output", "", "output to store results")
	flag.IntVar(&reps, "reps", 1000, "number of repetitions to calculate duration")
	flag.Parse()

	if o.URL == "" {
		log.Fatalln("url must be set")
	}
	if p == "" {
		log.Fatalln("path to store output must be set")
	}

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	if err := chromedp.Run(
		ctx,
		chromedp.Navigate(o.URL),
		chromedp.Sleep(1000*time.Millisecond),
	); err != nil {
		log.Fatalln(err)
	}

	if err := chromedp.Run(
		ctx,
		chromedp.EvaluateAsDevTools(timer(reps), &o.ElapsedUS),
	); err != nil {
		log.Fatalln(err)
	}

	o.mean()

	b, err := json.Marshal(o)
	if err != nil {
		log.Fatalln(err)
	}

	if err := os.WriteFile(p, b, 0664); err != nil {
		log.Fatalln(err)
	}
}
