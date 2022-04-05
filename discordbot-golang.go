package main

import (
	"context"
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/disgord/std"
	"os"
)

func main() {
	client := disgord.New(disgord.Config{
		BotToken: os.Getenv("DISCORD_TOKEN"),
		Intents:  disgord.AllIntents(),
	})
	defer client.Gateway().StayConnectedUntilInterrupted()
	content, _ := std.NewMsgFilter(context.Background(), client)
	content.SetPrefix("ping")

	client.Gateway().
		WithMiddleware(content.HasPrefix).
		MessageCreate(func(s disgord.Session, evt *disgord.MessageCreate) {
			_, _ = evt.Message.Reply(context.Background(), s, "pong")
		})
	fmt.Println("Hello World")
}
