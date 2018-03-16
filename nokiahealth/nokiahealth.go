package nokiahealth

import (
	"fmt"
	"os"

	"github.com/jrmycanady/nokiahealth"
)

// InitUser initializes nokiahealth user with user credentislas
func InitUser() {
	client := nokiahealth.NewClient(os.Getenv("NOKIA_COUSUMER_KEY"), os.Getenv("NOKIT_CONSUMER_SECRET"), "")
	u := client.GenerateUser(os.Getenv("NOKIA_TOKEN"), os.Getenv("NOKIA_SECRET"), os.Getenv("NOKIA_USERID"))

	m, err := u.GetBodyMeasures(nil)
	fmt.Println(m)
}

func GetTodayBodyMeasure() {

}

func GetYesterdayBodyMeasure() {

}

func GetWeekAgoBodyMeasure() {

}

// 1週間前、1日前との比較値、cronで見せる用
func NotifyMeasures() {

}
