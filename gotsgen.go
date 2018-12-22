package gotsgen

import (
    "time"
    "math/rand"
)

// TimeSeries ...
type TimeSeries struct {
    XValues []time.Time
    YValues []float64
}

type TSGen struct {
  Start time.Time
  Duration time.Duration
  Samples int
  TS *TimeSeries
  Rand *rand.Rand
}

func (g TSGen) addRandomData() {
}

func (g TSGen) addNormalData() {
}

func (g TSGen) addDerivativeData() {
//    c := r.Float64()
//    p := c
//    n := c + r.NormFloat64()
}

// Init ...
func (g TSGen) Init(start time.Time, duration time.Duration) {
    g.Start = start
    g.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    g.Duration = duration
}

// New
func New(ts *TimeSeries, samples int) *TSGen {
    tsGen := &TSGen{
        Samples: samples,
    }
    return tsGen
}

// AddData ...
func AddData(ts *TimeSeries, start time.Time, d time.Duration, samples int) {
    t := start
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    c := r.Float64()
    p := c
    n := c + r.NormFloat64()
    for i := 0; i < samples; i++ {
        c = n
        n = c + r.NormFloat64()
        ts.XValues = append(ts.XValues, t)
        ts.YValues = append(ts.YValues, (n - p)/2 + 10)
        t = t.Add(d)
    }
}
