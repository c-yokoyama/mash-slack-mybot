package mynokiahealth

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	. "github.com/jrmycanady/nokiahealth"
)

type MeasureData struct {
	Weight     string
	FatRatio   string
	FatWight   string
	MuscleMass string
}

const redisURL = "redis:6379"

// Initial goal value
var weightGoal = "70.5"
var rc *redis.Client

// InitMyNokiaHealth intialize Nokiahealth User and redis connection
func InitMyNokiaHealth() User {
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

	// Connect to redis
	rc = redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	s, err := rc.Get("goal").Result()
	if err == redis.Nil {
		err = rc.Set("goal", weightGoal, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		// Update weightGoal with stored value
		fmt.Println("mynokiahealth: value is stored in redis: ", s)
		weightGoal = s
	}
	fmt.Println("mynokiahealth: Connected to redis!")

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
	res.Weight = strconv.FormatFloat(m.ParseData().Weights[0].Kgs, 'g', 4, 64)
	res.FatRatio = strconv.FormatFloat(m.ParseData().FatRatios[0].Ratio, 'g', 4, 64)
	res.FatWight = strconv.FormatFloat(m.ParseData().FatMassWeights[0].Kgs, 'g', 4, 64)
	res.MuscleMass = strconv.FormatFloat(m.ParseData().MuscleMasses[0].Mass, 'g', 4, 64)

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

func diffStrWithFloat(a string, b string) string {
	f32a, _ := strconv.ParseFloat(a, 32)
	f32b, _ := strconv.ParseFloat(b, 32)
	diff := f32a - f32b
	res := strconv.FormatFloat(diff, 'g', 4, 32)
	if diff > 0 {
		res = "+" + res
	}
	return res
}

func DiffTodayYesterdayMeasure(u User) MeasureData {
	res := MeasureData{}
	today := GetTodayBodyMeasure(u)
	yestday := getYesterdayBodyMeasure(u)

	res.Weight = diffStrWithFloat(today.Weight, yestday.Weight)
	res.FatWight = diffStrWithFloat(today.FatWight, yestday.FatWight)
	res.MuscleMass = diffStrWithFloat(today.MuscleMass, yestday.MuscleMass)

	//fmt.Printf("DiffTodayYesterdayMeasure: res: %v\n", res)

	return res
}

func DiffTodayWeekAgoMeasure(u User) MeasureData {
	res := MeasureData{}
	today := GetTodayBodyMeasure(u)
	weekago := getWeekAgoBodyMeasure(u)

	res.Weight = diffStrWithFloat(today.Weight, weekago.Weight)
	res.FatWight = diffStrWithFloat(today.FatWight, weekago.FatWight)
	res.MuscleMass = diffStrWithFloat(today.MuscleMass, weekago.MuscleMass)

	//fmt.Printf("DiffTodayWeekAgoMeasure: res: %v\n", res)

	return res
}

func DiffTodayWeightGoal(u User) string {
	today := GetTodayBodyMeasure(u)
	return diffStrWithFloat(today.Weight, getWeightGoal())

}

func SetWeightGoal(goal float64) {
	weightGoal = strconv.FormatFloat(goal, 'g', 4, 64)
	err := rc.Set("goal", weightGoal, 0).Err()
	if err != nil {
		log.Fatal(err)
	}
}

func getWeightGoal() string {
	// Check redis update
	s, err := rc.Get("goal").Result()
	if err == redis.Nil {
		err = rc.Set("goal", weightGoal, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
	} else if err != nil {
		log.Fatal(err)
	} else {
		// Update weightGoal with stored value
		weightGoal = s
	}

	return weightGoal
}
