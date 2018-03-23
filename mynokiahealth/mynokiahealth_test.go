package mynokiahealth

import (
	"reflect"
	"testing"
)

func TestgetBodyMeasureWithDay(t *testing.T) {
	u := NewNokiaHealthUser()
	getBodyMeasureWithDay(u, 0)
	getBodyMeasureWithDay(u, -1)
	getBodyMeasureWithDay(u, -7)
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
