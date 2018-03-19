package nokiahealth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	. "github.com/jrmycanady/nokiahealth"
)

type MeasureData struct {
	weight     float64
	fatRatio   float64
	fatWight   float64
	muscleMass float64
}

// InitUser initializes nokiahealth user with user credentislas
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

func GetTodayBodyMeasure(u User) MeasureData {
	res := MeasureData{}
	p := BodyMeasuresQueryParams{}
	day := time.Now().AddDate(0, 0, -1)
	l := 1
	p.StartDate = &day
	p.Limit = &l

	m, err := u.GetBodyMeasures(&p)
	if err != nil {
		log.Fatal(err)
	}
	res.weight = m.ParseData().Weights[0].Kgs
	res.fatRatio = m.ParseData().FatRatios[0].Ratio
	res.fatWight = m.ParseData().FatMassWeights[0].Kgs
	res.muscleMass = m.ParseData().MuscleMasses[0].Mass

	fmt.Println(res)

	return res
}

func GetYesterdayBodyMeasure(u User) {

}

func GetWeekAgoBodyMeasure(u User) {

}

// 1週間前、1日前との比較値、cronで見せる用
func NotifyMeasures(user User) {

}
