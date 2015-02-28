package main

import (
	"fmt"
	"os"

	"github.com/usabilla/gobilla"
)

func buttons(key, secret string) {
	gb := gobilla.New(key, secret)

	b := gb.Buttons()
	buttons, err := b.Get(nil)
	if err != nil {
		fmt.Errorf("%s", err)
	}
	for _, button := range buttons.Items {
		fmt.Println(button)
		// feedback, err := b.Feedback().Get(button["id"], nil)
		// if err != nil {
		// 	fmt.Errorf("%s", err)
		// }
		// fmt.Printf("Feedback for button with id: %s\n%s\n", button.ID, feedback.Items)
	}
}

func campaigns(key, secret string) {
	gb := gobilla.New(key, secret)

	c := gb.Campaigns()
	campaigns, err := c.Get(map[string]string{"limit": "1"})
	if err != nil {
		panic(err)
	}
	for _, campaign := range campaigns.Items {
		results, err := c.Results().Get(campaign.ID, map[string]string{"limit": "2"})
		if err != nil {
			panic(err)
		}
		for _, result := range results.Items {
			fmt.Printf("Result for campaign with id: %s\n%+v\n", campaign.ID, result)
		}
	}
}

func main() {
	args := os.Args[1:]

	buttons(args[0], args[1])

	// campaigns(args[0], args[1])
}
