package gotsgen

import (
    "testing"
    "time"
    "github.com/davecgh/go-spew/spew"
)

func TestNew(t *testing.T) {

    duration, _ := time.ParseDuration("24h")
    end := time.Now()
    start := end.Add(-duration)

    tsGen := New(start, duration, 200)
    spew.Dump(tsGen)
    if tsGen.Samples != 200 {
         t.Error("Expected 200, got ", tsGen.Samples)
    }
}

func TestInit(t *testing.T) {
    duration, _ := time.ParseDuration("24h")
    end := time.Now()
    start := end.Add(-duration)

    tsGen := New(start, duration/200, 200)

    spew.Dump(tsGen)

    tsGen.Init("rand")
    spew.Dump(tsGen)
}

