package main

import (
	"fmt"
	"os"

	"github.com/usabilla/gobilla"
)

func buttons(gb *gobilla.Gobilla) {
	b := gb.Buttons()

	buttons, err := b.Get(nil)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	for _, button := range buttons.Items {
		resp, err := b.Feedback().Get(button.ID, nil)
		if err != nil {
			fmt.Errorf("%s", err)
		}
		count := 0
		fmt.Printf("START PRINTING FEEDBACK FOR BUTTON: %s\n", button.ID)
		for _, item := range resp.Items {
			fmt.Printf("FEEDBACK %s\n", item.ID)
			count++
		}
		fmt.Printf("RECEIVED %d FEEDBACK ITEMS\n", count)
	}
	fmt.Printf("RECEIVED FEEDBACK FROM %d BUTTONS\n", buttons.Count)
}

func buttonsIterator(gb *gobilla.Gobilla) {
	b := gb.Buttons()

	buttons, err := b.Get(nil)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	for _, button := range buttons.Items {
		count := 0
		fmt.Printf("START PRINTING FEEDBACK FOR BUTTON: %s\n", button.ID)
		for item := range b.Feedback().Iterate(button.ID, nil) {
			fmt.Printf("FEEDBACK %s\n", item.ID)
			count++
		}
		fmt.Printf("RECEIVED %d FEEDBACK ITEMS\n", count)
	}
	fmt.Printf("RECEIVED FEEDBACK FROM %d BUTTONS\n", buttons.Count)
}

func main() {
	key := os.Getenv("USABILLA_API_KEY")
	secret := os.Getenv("USABILLA_API_SECRET")

	// You can pass a custom http.Client
	// We pass nil to use the http.DefaultClient
	gb := gobilla.New(key, secret, nil)

	// Uses a simple GET to get all feedback items for all buttons.
	buttons(gb)

	// Uses a channel of feedback items, and once all items have been
	// consumed and the response HasMore then it fires a new request
	// for all feedback items for all buttons.
	buttonsIterator(gb)
}
