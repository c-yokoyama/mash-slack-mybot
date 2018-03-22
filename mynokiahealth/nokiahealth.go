package mynokiahealth

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	. "github.com/jrmycanady/nokiahealth"
)

type MeasureData struct {
	weight     float32
	fatRatio   float32
	fatWight   float32
	muscleMass float32
}

var goalWeight float32 = 70.5

// NewNokiaHealthUser initializes nokiahealth user with user credentislas
func NewNokiaHealthUser() User {
	client := NewClient(os.Getenv("NOKIA_COUSUMER_KEY"), os.Getenv("NOKIT_CONSUMER_SECRET"), "")
	userid, _ := strconv.Atoi(os.Getenv("NOKIA_USERID"))
	u := client.GenerateUser(os.Getenv("NOKIA_TOKEN"), os.Getenv("NOKIA_SECRET"), userid)

	m, err := u.GetBodyMeasures(nil)

	if err != nil {
		log.Fatal(err)
	}

	if m.Status.String() != "OperationWasSuccessful" {
		s := "Fail to generate user: " + m.Status.String()
		log.Fatal(errors.New(s))
	}
	return u
}

func getBodyMeasureWithDay(u User, d int) MeasureData {
	res := MeasureData{}
	p := BodyMeasuresQueryParams{}
	day := time.Now().AddDate(0, 0, d)
	p.EndDate = &day

	m, err := u.GetBodyMeasures(&p)
	if err != nil {
		log.Fatal(err)
	}
	res.weight = float32(m.ParseData().Weights[0].Kgs)
	res.fatRatio = float32(m.ParseData().FatRatios[0].Ratio)
	res.fatWight = float32(m.ParseData().FatMassWeights[0].Kgs)
	res.muscleMass = float32(m.ParseData().MuscleMasses[0].Mass)

	//fmt.Printf("Day: %v, Res: %v\n", day, res)

	return res
}

func GetTodayBodyMeasure(u User) MeasureData {
	return getBodyMeasureWithDay(u, 0)
}

func getYesterdayBodyMeasure(u User) MeasureData {
	return getBodyMeasureWithDay(u, -1)
}

func getWeekAgoBodyMeasure(u User) MeasureData {
	return getBodyMeasureWithDay(u, -7)

}

func DiffTodayYesterdayMeasure(u User) MeasureData {
	res := MeasureData{}
	today := GetTodayBodyMeasure(u)
	yestday := getYesterdayBodyMeasure(u)

	res.weight = today.weight - yestday.weight
	res.fatRatio = today.fatRatio - yestday.fatRatio
	res.fatWight = today.fatWight - yestday.fatWight
	res.muscleMass = today.muscleMass - yestday.muscleMass

	//fmt.Printf("DiffTodayYesterdayMeasure: res: %v\n", res)

	return res
}

func DiffTodayWeekAgoMeasure(u User) MeasureData {
	res := MeasureData{}
	today := GetTodayBodyMeasure(u)
	weekago := getWeekAgoBodyMeasure(u)

	res.weight = today.weight - weekago.weight
	res.fatRatio = today.fatRatio - weekago.fatRatio
	res.fatWight = today.fatWight - weekago.fatWight
	res.muscleMass = today.muscleMass - weekago.muscleMass

	//fmt.Printf("DiffTodayWeekAgoMeasure: res: %v\n", res)

	return res
}

func DiffTodayWeightGoal(u User) float32 {
	today := GetTodayBodyMeasure(u)
	//四捨五入してかえす
	return today.weight - getWeightGoal()

}

func SetWeightGoal(goal float32) {
	goalWeight = goal
}

func getWeightGoal() float32 {
	return goalWeight
}
