/*
 * Copyright (c) 2019 InterCloud
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gotsgen

import (
    "time"
    "errors"
    "math/rand"
)

// TimeSeries contains time series measurements.
type TimeSeries struct {
    XValues []time.Time
    YValues []float64
}

// TSGen contains informations to generate a time series.
type TSGen struct {
  Start time.Time
  Period time.Duration
  Samples uint
  TS *TimeSeries}

// Init starts the generator.
// It takes a type of generator as parameter.
// The valid types are:
//
//     "rand":  values are random generated, see `math/rand` package (Float64())
//     "norm":  values are generated from a normal distribution, see `math/rand` package (NormFloat64())
//     "deriv": values are generated from the dicrete derivative of a continuously and randomly increasing counter
//
// You can use it like this:
//
//      gts.Init("rand")
func (g TSGen) Init(t string) error {
    typeFunc := map[string]interface{}{
        "rand": g.addRandomData,
        "norm": g.addNormalData,
        "deriv": g.addDerivativeData,
    }
    if _, ok := typeFunc[t]; !ok {
        return errors.New("Unknown generator type")
    }
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
    typeFunc[t].(func(*rand.Rand))(r)

    return nil
}

// New creates a new generator.
// It takes 3 parameters:
//
//     - start    is the starting point of thetime series
//     - period   is the sampling rate
//     - samples  is the number of measurements of the timeseries
//
// You can use it like this:
//
//      duration, _ := time.ParseDuration("24h")
//      end := time.Now()
//      start := end.Add(-duration)
//      gts = gotsgen.New(start, duration/200, 200)
func New(start time.Time, period time.Duration, samples uint) *TSGen {
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
