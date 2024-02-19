package main

import (
	"context" // Importing context package for managing contexts
	"fmt"     // Importing fmt package for formatted I/O
	"log"     // Importing log package for logging
	"os"      // Importing os package for operating system functionality
	"strconv" // Importing strconv package for string conversions
	"time"    // Importing time package for time-related operations

	"github.com/shomali11/slacker" // Importing slacker package for creating Slack bots
)

// printCommandEvents prints command events received from the bot
func printCommandEvents(analyticsChannel <-chan *slacker.CommandEvents) {
	for event := range analyticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}

func main() {
	// Setting environment variables for Slack bot token and app token
	os.Setenv("SLACK_BOT_TOKEN", "YOUR SLACK BOT TOKEN HERE")
	os.Setenv("SLACK_APP_TOKEN", "YOUR SLACK APP TOKEN HERE")

	// Creating a new Slack bot client
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))

	// Goroutine to print command events
	go printCommandEvents(bot.CommandEvents())

	// Registering a command handler for "my year of birth is <year>"
	bot.Command("my year of birth is <year>", &slacker.CommandDefinition{
		Description: "year of birth calculator", // Command description
		Example:     "my year of birth is 2020", // Example usage
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")  // Retrieving year parameter from the request
			yob, err := strconv.Atoi(year) // Converting year string to integer
			if err != nil {                // Handling conversion error
				fmt.Println(err)
			}
			age := (time.Now().Year() - yob)     // Calculating age based on current year and year of birth
			r := fmt.Sprintf("age is %d\n", age) // Constructing response message with calculated age
			response.Reply(r)                    // Replying to the Slack channel with the response message
		},
	})

	ctx, cancel := context.WithCancel(context.Background()) // Creating a new context
	defer cancel()                                          // Deferring the cancellation of the context
	err := bot.Listen(ctx)                                  // Starting the bot and listening for events
	if err != nil {
		log.Fatal(err) // Logging fatal error if bot encounters any issues
	}
}
