package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"genpages/speedtest"
)

func main() {
	var url, p string
	var reps int

	flag.StringVar(&url, "url", "", "url to call")
	flag.StringVar(&p, "output", "", "Output to store results")
	flag.IntVar(&reps, "reps", 1000, "number of repetitions to calculate duration")
	flag.Parse()

	if url == "" {
		log.Fatalln("url must be set")
	}
	if p == "" {
		log.Fatalln("path to store Output must be set")
	}

	o, err := speedtest.Run(url, reps)
	if err != nil {
		log.Fatalln(err)
	}

	b, err := json.Marshal(o)
	if err != nil {
		log.Fatalln(err)
	}

	if err := os.WriteFile(p, b, 0664); err != nil {
		log.Fatalln(err)
	}
}
