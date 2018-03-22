package main

import (
	"fmt"
	"log"
	"mash-slack-mybot/mynokiahealth"
	"os"
	"strings"

	"github.com/jrmycanady/nokiahealth"
	"github.com/nlopes/slack"
)

var botHelp = `
@
`

func getMyMeasures(u nokiahealth.User, args []string) string {
	// @bot measure hogehogeでswitch
	return "your-measures-value"
}

func main() {
	api := slack.New(os.Getenv("TOKEN"))
	botUserID := os.Getenv("BOTUSER_ID")

	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	nokiaUser := mynokiahealth.NewNokiaHealthUser()

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
					rtm.SendMessage(rtm.NewOutgoingMessage("機能一覧です！"+botHelp, ev.Channel))
					break
				}
				switch args[1] {
				case "measure":
					res := getMyMeasures(nokiaUser, args[2:])
					fmt.Println("Measure Res: " + res)
					rtm.SendMessage(rtm.NewOutgoingMessage(res, ev.Channel))
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
