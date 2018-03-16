package nokiahealth

import (
	"testing"
)

func TestGetTodayBodyMeasure(t *testing.T) {
	u := InitUser()
	GetTodayBodyMeasure(u)
}
