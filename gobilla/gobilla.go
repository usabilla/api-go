package main

import (
	"fmt"
	"os"

	"github.com/usabilla/gobilla"
)

func main() {
	args := os.Args[1:]
	// the zero index is the key and the first index is the secret
	api := gobilla.NewClient(args[0], args[1])

	// limit for buttons does not work
	buttons, err := api.Buttons.Get(map[string]string{"limit": "2"})
	if err != nil {
		panic(err)
	}
	for _, button := range buttons.Items {
		feedback, err := api.Buttons.Feedback.Get(button.ID, nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Feedback for button with id: %s\n%s\n", button.ID, feedback.Items)
	}

	campaigns, err := api.Campaigns.Get(nil)
	if err != nil {
		panic(err)
	}
	for _, campaign := range campaigns.Items {
		results, err := api.Campaigns.Results.Get(campaign.ID, map[string]string{"limit": "2"})
		if err != nil {
			panic(err)
		}
		for _, result := range results.Items {
			fmt.Printf("Result for campaign with id: %s\n%+v\n", campaign.ID, result)
		}
	}
}
