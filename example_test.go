package gobilla_test

import (
	"fmt"

	"github.com/usabilla/gobilla"
)

func ExampleButtons() {
	gb := gobilla.New("key", "secret")

	b := gb.Buttons()
	buttons, err := b.Get(map[string]string{"limit": "2"})
	if err != nil {
		fmt.Errorf("%s", err)
	}
	for _, button := range buttons.Items {
		feedback, err := b.Feedback().Get(button.ID, nil)
		if err != nil {
			fmt.Errorf("%s", err)
		}
		fmt.Printf("Feedback for button with id: %s\n%s\n", button.ID, feedback.Items)
	}
	// Output:
	//
}

func ExampleCampaigns() {
	gb := gobilla.New("key", "secret")

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
	// Output:
	//
}
