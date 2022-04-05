package main

import (
	"context"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"os"
)

//var log = logrus.New()
//
//var commands = []*disgord.CreateApplicationCommand{
//	{
//		Name:        "test_command",
//		Description: "just testing",
//		Options: []*disgord.ApplicationCommandOption{
//			{
//				Name:        "test_option",
//				Type:        disgord.OptionTypeString,
//				Description: "testing options",
//				Choices: []*disgord.ApplicationCommandOptionChoice{
//					{
//						Name:  "test_choice",
//						Value: "test_val",
//					},
//				},
//			},
//		},
//	},
//}

func main() {
	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISCORD_TOKEN"),
		Intents:  disgord.AllIntents(),
	})
	defer func(gateway disgord.GatewayQueryBuilder) {
		err := gateway.StayConnectedUntilInterrupted()
		if err != nil {

		}
	}(client.Gateway())

	u, err := client.BotAuthorizeURL(disgord.PermissionUseSlashCommands, []string{
		"bot",
		"applications.commands",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(u)

	content, _ := std.NewMsgFilter(context.Background(), client)
	content.SetPrefix("ping")
	client.Gateway().
		WithMiddleware(content.HasPrefix).
		MessageCreate(func(s disgord.Session, evt *disgord.MessageCreate) {
			_, _ = evt.Message.Reply(context.Background(), s, "pong")
		})
}
