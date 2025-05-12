package helper

import (
	"testing"
	"time"
)

func TestBusinessDayFromStart(t *testing.T) {

	res := BusinessDayFromStart(time.Date(2006, 1, 2, 15, 04, 05, 0, time.UTC), 5)
	expectedRes := time.Date(2006, 1, 9, 15, 04, 05, 0, time.UTC)

	if !res.Equal(expectedRes) {
		t.Errorf("BusinessDaysFromStart(2006-01-02 15:04:05.000 UTC, 5) = %v, want %v", res, expectedRes)
	}
}
