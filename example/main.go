package main

import (
    "log"
    "net/http"
    "time"
    chart "github.com/wcharczuk/go-chart"
    "github.com/wcharczuk/go-chart/drawing"
    "github.com/intercloud/gotsgen"
)

const SAMPLES = 200
const SCALE   = "24h"

func drawChart(res http.ResponseWriter, req *http.Request) {
    var graph *chart.Chart
    var ts *chart.TimeSeries

    ts = &chart.TimeSeries{
        Style: chart.Style{
             Show:        true,                           //note; if we set ANY other properties, we must set this to true.
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

    gts := gotsgen.New(start, duration/SAMPLES, SAMPLES)
    gts.Init("rand")
    // add generated data to chart TS
    for i := 0 ; i < len(gts.TS.XValues); i++ {
        ts.XValues = append(ts.XValues, gts.TS.XValues[i])
        ts.YValues = append(ts.YValues, gts.TS.YValues[i])
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


