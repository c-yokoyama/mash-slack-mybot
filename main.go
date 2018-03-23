package main

import (
	"fmt"
	"log"
	"mash-slack-mybot/mynokiahealth"
	"os"
	"strconv"
	"strings"

	"github.com/jrmycanady/nokiahealth"
	"github.com/nlopes/slack"
	"gopkg.in/robfig/cron.v2"
)

var rtm *slack.RTM
var nokiaUser nokiahealth.User

const cronSpec = "TZ=Asia/Tokyo 30 11 * * * *"

func getHelpStr(botUserID string) string {
	help := "機能一覧です！\n" +
		"<@" + botUserID + "> measure: 本日(最新)の測定結果を表示します\n" +
		"<@" + botUserID + "> measure goal: 本日(最新)の測定結果と目標体重の差分を表示します\n" +
		"<@" + botUserID + "> measure set goal <value>: 目標体重を設定します\n" +
		"<@" + botUserID + "> help: 機能一覧を表示します\n" +
		"\n"
	return help
}

func getMeasureDetail(u nokiahealth.User) string {
	today := mynokiahealth.GetTodayBodyMeasure(u)
	diffDay := mynokiahealth.DiffTodayYesterdayMeasure(u)
	diffWeek := mynokiahealth.DiffTodayWeekAgoMeasure(u)
	diffGoal := mynokiahealth.DiffTodayWeightGoal(u)

	res := "最新の測定結果は、\n" +
		"-------------------\n" +
		"体重 `" + today.Weight + "kg` \n" +
		"体脂肪率 `" + today.FatRatio + "%` \n" +
		"体脂肪量 `" + today.FatWight + "kg `\n" +
		"筋肉量 `" + today.MuscleMass + "kg `\n" +
		"-------------------\n" +
		"です！\n" +
		"前回測定時との差分は、\n" +
		"-------------------\n" +
		"体重 `" + diffDay.Weight + "kg` \n" +
		"体脂肪量 `" + diffDay.FatWight + "kg `\n" +
		"筋肉量 `" + diffDay.MuscleMass + "kg `\n" +
		"-------------------\n" +
		"です！\n" +
		"約1週間前の測定時との差分は、\n" +
		"-------------------\n" +
		"体重 `" + diffWeek.Weight + "kg` \n" +
		"体脂肪量 `" + diffWeek.FatWight + "kg `\n" +
		"筋肉量 `" + diffWeek.MuscleMass + "kg `\n" +
		"-------------------\n" +
		"です！\n" +
		"目標体重との差分は `" + diffGoal + "kg` です、頑張りましょう。\n"

	return res
}

func sendCronMessage() {
	res := getMeasureDetail(nokiaUser)
	cronChanID := os.Getenv("CRON_CHAN_ID")
	fmt.Println("cronChanID: " + cronChanID)
	rtm.SendMessage(rtm.NewOutgoingMessage(res, cronChanID))
}

func getMyMeasures(u nokiahealth.User, args []string) string {
	if len(args) == 0 {
		return getMeasureDetail(u)
	}

	switch args[0] {
	case "goal":
		diffGoal := mynokiahealth.DiffTodayWeightGoal(u)
		return "最新の測定結果と目標体重との差分は `" + diffGoal + "kg` です、頑張りましょう。\n"
	case "set":
		if args[1] != "goal" || len(args) != 3 {
			return "`measure set goal <value>` ですよ。"
		}
		goal, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return "値が間違っています。"
		}
		mynokiahealth.SetWeightGoal(goal)
		return "目標を設定しました。"

	default:
		return "コマンドが間違っています。"
	}
}

func main() {
	api := slack.New(os.Getenv("TOKEN"))
	botUserID := os.Getenv("BOTUSER_ID")

	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)
	rtm = api.NewRTM()
	go rtm.ManageConnection()

	nokiaUser = mynokiahealth.NewNokiaHealthUser()
	c := cron.New()
	c.AddFunc(cronSpec, sendCronMessage)
	c.Start()

	for msg := range rtm.IncomingEvents {
		fmt.Print("EventReceived: ")
		switch ev := msg.Data.(type) {
		case *slack.ConnectedEvent:
			//fmt.Println("Infos:", ev.Info)
			//fmt.Println("Connection counter:", ev.ConnectionCount)
		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)
			// Msg includes mention to botUser
			if strings.HasPrefix(ev.Msg.Text, "<@"+botUserID+">") {
				args := strings.Split(ev.Msg.Text, " ")
				if len(args) <= 1 {
					rtm.SendMessage(rtm.NewOutgoingMessage(getHelpStr(botUserID), ev.Channel))
					break
				}
				switch args[1] {
				case "measure":
					res := getMyMeasures(nokiaUser, args[2:])
					rtm.SendMessage(rtm.NewOutgoingMessage(res, ev.Channel))

				case "help":
					rtm.SendMessage(rtm.NewOutgoingMessage(getHelpStr(botUserID), ev.Channel))
				}
			}
		case *slack.PresenceChangeEvent:
			//fmt.Printf("Presence Change: %v\n", ev)

		case *slack.LatencyReport:
			//fmt.Printf("Current latency: %v\n", ev.Value)

		case *slack.RTMError:
			fmt.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			fmt.Printf("Invalid credentials")
			return

		default:
			// fmt.Printf("Unexpected: %v\n", msg.Data)
		}
	}
}
