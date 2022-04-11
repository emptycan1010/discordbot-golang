package main

import (
	"context"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"os"
)

var log = logrus.New()

var commands = []*disgord.CreateApplicationCommand{
	{
		Name:        "test_command",
		Description: "just testing",
		Options: []*disgord.ApplicationCommandOption{
			{
				Name:        "test_option",
				Type:        disgord.OptionTypeString,
				Description: "testing options",
				Choices: []*disgord.ApplicationCommandOptionChoice{
					{
						Name:  "test_choice",
						Value: "test_val",
					},
				},
			},
		},
	},
}

func main() {
	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISCORD_TOKEN"),
		Logger:   log,
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

	client.Gateway().BotReady(func() {
		for i := range commands {
			// application command id is 0 here
			// on a ready event, the client is updated to store the application id
			// you can fetch the application id using the bot id (current user id) or copy it from
			// the discord page.
			if err = client.ApplicationCommand(0).Guild(923185824869285938).Create(commands[i]); err != nil {
				log.Fatal(err)
			}
		}
	})

	client.Gateway().InteractionCreate(func(s disgord.Session, h *disgord.InteractionCreate) {
		fmt.Printf("%+v", *h)
		err := s.SendInteractionResponse(context.Background(), h, &disgord.CreateInteractionResponse{
			Type: 4,
			Data: &disgord.CreateInteractionResponseData{
				Content:    "hello",
				Components: []*disgord.MessageComponent{},
			},
		})
		if err != nil {
			log.Error(err)
		}
	})
}
