package main

import (
	"fmt"
	"log"
	"mash-slack-mybot/nokiahealth"
	"os"
	"strings"
	//"mash-slack-mybot/nokiahealth"
	"github.com/nlopes/slack"
)

func getMyMeasures() string {
	// @bot measure hogehogeでswitch
	return "your-measures-value"
}

func main() {
	api := slack.New(os.Getenv("TOKEN"))
	botUserID := os.Getenv("BOTUSER_ID")

	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(true)

	nokiahealth.InitUser()

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("EventReceived: ")
		switch ev := msg.Data.(type) {

		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			//fmt.Println("Infos:", ev.Info)
			//fmt.Println("Connection counter:", ev.ConnectionCount)

		case *slack.MessageEvent:
			fmt.Printf("Message: %v\n", ev)

			// Msg includes mention to botUser
			if strings.HasPrefix(ev.Msg.Text, "<@"+botUserID+">") {
				args := strings.Split(ev.Msg.Text, " ")
				if len(args) <= 1 {
					rtm.SendMessage(rtm.NewOutgoingMessage("機能一覧はこちらです！", ev.Channel))
					break
				}
				switch args[1] {
				case "measure":
					fmt.Println(getMyMeasures())
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
