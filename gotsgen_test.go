package gotsgen

import (
	"testing"
	"time"
)

func TestQuery(t *testing.T) {
	duration, _ := time.ParseDuration("24h")
	end := time.Now()
	start := end.Add(-duration)

	ts, err := Query(start, end, 200, "rand")

	if err != nil {
		t.Errorf("Unexpected error %s\n", err.Error())
	}
	if len(ts.XValues) != 200 {
		t.Errorf("Expected time series to have 200 values got %d\n", len(ts.XValues))
	}

	ts, err = Query(start, end, 200, "fake")
	if err == nil || err.Error() != "Unknown generator type" {
		t.Errorf("Expected error Unknown generator type but got %v\n", err)
	}

	ts, err = Query(start, start, 200, "norm")
	if err == nil || err.Error() != "Bad time range" {
		t.Errorf("Expected error Unknown generator type but got %v\n", err)
	}

	ts, err = Query(end, start, 200, "deriv")
	if err == nil || err.Error() != "Bad time range" {
		t.Errorf("Expected error Unknown generator type but got %v\n", err)
	}
}
