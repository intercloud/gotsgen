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
  Period time.Duration
  Samples int
  TS *TimeSeries
}


func (g TSGen) addRandomData(r *rand.Rand) {
    t := g.Start
    for i := 0; i < g.Samples; i++ {
        g.TS.XValues = append(g.TS.XValues, t)
        g.TS.YValues = append(g.TS.YValues, r.Float64())
        t = t.Add(g.Period)
    }
}

func (g TSGen) addNormalData(r *rand.Rand) {
}

func (g TSGen) addDerivativeData(r *rand.Rand) {
//    c := r.Float64()
//    p := c
//    n := c + r.NormFloat64()
}


// Init ...
func (g TSGen) Init(t string) {
    typeFunc := map[string]interface{}{
        "rand": g.addRandomData,
        "norm": g.addNormalData,
        "deriv": g.addDerivativeData,
    }
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    typeFunc[t].(func(*rand.Rand))(r)
}

// New
func New(start time.Time, period time.Duration, samples int) *TSGen {
    ts := &TimeSeries{
        XValues: []time.Time{},
        YValues: []float64{},
    }
    tsGen := &TSGen{
        TS: ts,
        Start: start,
        Period: period,
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
