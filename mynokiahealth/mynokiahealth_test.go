package mynokiahealth

import (
	"reflect"
	"testing"
)

func TestGetTodayBodyMeasure(t *testing.T) {
	u := NewNokiaHealthUser()
	GetTodayBodyMeasure(u)
}

func TestGetYesterdayBodyMeasure(t *testing.T) {
	u := NewNokiaHealthUser()
	getYesterdayBodyMeasure(u)
}

func TestWeekAgoBodyMeasure(t *testing.T) {
	u := NewNokiaHealthUser()
	getWeekAgoBodyMeasure(u)
}

func TestDiffTodayYesterdayMeasure(t *testing.T) {
	u := NewNokiaHealthUser()
	DiffTodayYesterdayMeasure(u)
}

func TestDiffTodayWeekAgoMeasure(t *testing.T) {
	u := NewNokiaHealthUser()
	DiffTodayWeekAgoMeasure(u)

}

func TestDiffTodayWeightGoal(t *testing.T) {
	u := NewNokiaHealthUser()
	res := DiffTodayWeightGoal(u)
	resType := reflect.TypeOf(res).String()
	if resType != "string" {
		t.Errorf("Expected: string, Actual: %s\n", resType)
	}
}
