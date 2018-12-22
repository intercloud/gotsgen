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
}

func TestInit(t *testing.T) {
    duration, _ := time.ParseDuration("24h")
    end := time.Now()
    start := end.Add(-duration)

    tsGen := New(start, duration/200, 200)


    tsGen.Init("rand")
}

