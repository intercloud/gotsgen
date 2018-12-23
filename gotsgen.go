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
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// TimeSeries contains time series measurements.
type TimeSeries struct {
	XValues []time.Time
	YValues []float64
}

// Query generates a time series from the given type.
// It takes 3 parameters:
//
//     - start    is the starting point of the time series
//     - end      is the end of time series scale
//     - samples  is the number of measurements of the time series
//     - type     is the category of generator
//
// The valid types of generators are:
//
//     "rand":  values are random generated, see `math/rand` package (Float64())
//     "norm":  values are generated from a normal distribution, see `math/rand` package (NormFloat64())
//     "deriv": values are generated from the dicrete derivative of a continuously and randomly increasing counter
//
// Example:
//
//      duration, _ := time.ParseDuration("24h")
//      end := time.Now()
//      start := end.Add(-duration)
//      ts, err := Query(start, end, 200, "rand")
func Query(start time.Time, end time.Time, samples uint, t string) (*TimeSeries, error) {
	typeFunc := map[string]interface{}{
		"rand":  addRandomData,
		"norm":  addNormalData,
		"deriv": addDerivativeData,
	}

	ts := &TimeSeries{
		XValues: []time.Time{},
		YValues: []float64{},
	}

	// todo test if end date != start date and if end date > start date
	if !end.After(start) || end.Equal(start) {
		return ts, errors.New("Bad time range")
	}
	// test if generator type is valid
	if _, ok := typeFunc[t]; !ok {
		return ts, errors.New("Unknown generator type")
	}

	duration := end.Sub(start).Seconds() / float64(samples)
	period, _ := time.ParseDuration(fmt.Sprintf("%fs", duration))
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// call the generator (see generators.go) corresponding to the given category
	typeFunc[t].(func(*TimeSeries, time.Time, time.Duration, uint, *rand.Rand))(ts, start, period, samples, r)

	return ts, nil
}
