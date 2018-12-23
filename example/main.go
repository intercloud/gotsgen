package main

import (
	"github.com/intercloud/gotsgen"
	chart "github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
	"log"
	"net/http"
	"time"
)

// SAMPLES is the number od points of measurement
const SAMPLES = 200

// SCALE is the duration of the full time series
const SCALE = "24h"

func drawChart(res http.ResponseWriter, req *http.Request) {
	var graph *chart.Chart
	var ts *chart.TimeSeries

	ts = &chart.TimeSeries{
		Style: chart.Style{
			Show:        true,                            //note; if we set ANY other properties, we must set this to true.
			StrokeColor: drawing.ColorBlue,               // will supercede defaults
			FillColor:   drawing.ColorBlue.WithAlpha(64), // will supercede defaults
		},
		XValues: []time.Time{},
		YValues: []float64{},
	}

	graph = &chart.Chart{
		XAxis: chart.XAxis{
			Style:          chart.StyleShow(),
			ValueFormatter: chart.TimeHourValueFormatter,
		},
		Series: []chart.Series{ts},
	}

	duration, _ := time.ParseDuration(SCALE)
	end := time.Now()
	start := end.Add(-duration)

	gts, _ := gotsgen.Query(start, end, SAMPLES, "rand")
	// add generated data to chart TS
	for i := 0; i < len(gts.XValues); i++ {
		ts.XValues = append(ts.XValues, gts.XValues[i])
		ts.YValues = append(ts.YValues, gts.YValues[i])
	}

	if len(ts.XValues) == 0 {
		http.Error(res, "no data (yet)", http.StatusBadRequest)
		return
	}
	res.Header().Set("Content-Type", "image/png")
	if err := graph.Render(chart.PNG, res); err != nil {
		log.Printf("%v", err)
	}
}

func main() {
	http.HandleFunc("/", drawChart)
	log.Printf("Open http://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
