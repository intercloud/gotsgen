package gotsgen

import (
    "testing"
    "time"
)

func TestNew(t *testing.T) {

    duration, _ := time.ParseDuration("24h")
    end := time.Now()
    start := end.Add(-duration)

    tsGen := New(start, duration/200, 200)
    if tsGen.Samples != 200 {
         t.Error("Expected 200 samples, got ", tsGen.Samples)
    }
    if tsGen.Start != start {
         t.Errorf("Expected start date to be %v, got %v\n", start, tsGen.Start)
    }
    if tsGen.Period != duration/200 {
         t.Errorf("Expected period to be %v, got %v\n", duration/200, tsGen.Period)
    }
}

func TestInit(t *testing.T) {
    duration, _ := time.ParseDuration("24h")
    end := time.Now()
    start := end.Add(-duration)

    tsGen := New(start, duration/200, 200)

    if len(tsGen.TS.XValues) != 0 {
         t.Errorf("Expected time series to have no value got %d\n", len(tsGen.TS.XValues))
    }
    err := tsGen.Init("rand")
    if err != nil {
         t.Errorf("Unexpected error %s\n", err.Error())
    }
    if len(tsGen.TS.XValues) != 200 {
         t.Errorf("Expected time series to have 200 values got %d\n", len(tsGen.TS.XValues))
    }

    err = tsGen.Init("fake")
    if err == nil || err.Error() != "Unknown generator type" {
         t.Errorf("Expected error Unknown generator type but got %v\n", err)
    }
}

