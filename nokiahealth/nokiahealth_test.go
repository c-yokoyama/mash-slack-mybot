package nokiahealth

import (
	"testing"
)

func TestGetTodayBodyMeasure(t *testing.T) {
	u := NewNokiaHealthUser()
	GetTodayBodyMeasure(u)
}
