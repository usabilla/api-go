package main

import (
	"fmt"
	"os"

	"github.com/usabilla/gobilla"
)

func main() {
	args := os.Args[1:]
	// the zero index is the key and the first index is the secret
	gb := gobilla.New(args[0], args[1])

	b := gb.Buttons()
	buttons, err := b.Get(map[string]string{"limit": "2"}) // limit for buttons does not work
	if err != nil {
		panic(err)
	}
	f := b.Feedback()
	for _, button := range buttons.Items {
		feedback, err := f.Get(button.ID, nil)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Feedback for button with id: %s\n%s\n", button.ID, feedback.Items)
	}

	c := gb.Campaigns()
	campaigns, err := c.Get(map[string]string{"limit": "1"})
	if err != nil {
		panic(err)
	}
	r := c.Results()
	for _, campaign := range campaigns.Items {
		results, err := r.Get(campaign.ID, map[string]string{"limit": "2"})
		if err != nil {
			panic(err)
		}
		for _, result := range results.Items {
			fmt.Printf("Result for campaign with id: %s\n%+v\n", campaign.ID, result)
		}
	}
}
