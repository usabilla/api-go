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
		count := 0
		fmt.Printf("START PRINTING FEEDBACK FROM BUTTON: %s\n", button.ID)
		for item := range b.Feedback().Iterate(button.ID, nil) {
			fmt.Printf("FEEDBACK %s\n", item.ID)
			count++
		}
		fmt.Printf("RECEIVED %d FEEDBACK ITEMS\n", count)
	}
	fmt.Printf("RECEIVED FEEDBACK FROM %d BUTTONS\n", buttons.Count)
}

func campaigns(gb *gobilla.Gobilla) {
	c := gb.Campaigns()

	campaigns, err := c.Get(nil)
	if err != nil {
		panic(err)
	}

	for _, campaign := range campaigns.Items {
		count := 0
		fmt.Printf("START PRINTING CAMPAIGN RESULTS FROM CAMPAIGN: %s\n", campaign.ID)
		for result := range c.Results().Iterate(campaign.ID, nil) {
			fmt.Printf("CAMPAIGN RESULT %s\n", result.ID)
			count++
		}
		fmt.Printf("RECEIVED %d CAMPAIGN RESULTS\n", count)
	}
	fmt.Printf("RECEIVED CAMPAIGN RESULTS FROM %d CAMPAIGNS\n", campaigns.Count)
}

func main() {
	args := os.Args[1:]

	gb := gobilla.New(args[0], args[1])

	buttons(gb)

	campaigns(gb)
}
