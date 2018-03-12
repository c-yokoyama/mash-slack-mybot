package nokiahealth

import "github.com/jrmycanady/nokiahealth"

var client = nokiahealth.NewClient("<consumer_key>", " <consumer_secret>", "<callback_url>")

func GetTodayBodyMeasure() {

}

func GetYesterdayBodyMeasure() {

}

func GetWeekAgoBodyMeasure() {

}

// 1週間前、1日前との比較値、cronで見せる用
func NotifyMeasures() {

}
