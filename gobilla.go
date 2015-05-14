package gobilla

import (
	"github.com/usabilla/gobilla/auth"
	"github.com/usabilla/gobilla/l4a"
	"github.com/usabilla/gobilla/l4e"
	"github.com/usabilla/gobilla/l4w"
	"github.com/usabilla/gobilla/resource"
)

// Gobilla is the client that exposes all the resources of the Usabilla API.
type Gobilla struct {
	auth auth.Auth
}

// New creates a new Gobilla instance and sets the auth with key and secret.
func New(key, secret string) *Gobilla {
	return &Gobilla{
		auth: auth.Auth{
			Key:    key,
			Secret: secret,
		},
	}
}

// Buttons encapsulates the button resource.
func (gb *Gobilla) Buttons() l4w.Buttons {
	return l4w.Buttons{
		Resource: resource.Resource{
			Auth: gb.auth,
		},
	}
}

// Campaigns encapsulates the campaign resource.
func (gb *Gobilla) Campaigns() l4w.Campaigns {
	return l4w.Campaigns{
		Resource: resource.Resource{
			Auth: gb.auth,
		},
	}
}

// Apps encapsulates the app resource.
func (gb *Gobilla) Apps() l4a.Apps {
	return l4a.Apps{
		Resource: resource.Resource{
			Auth: gb.auth,
		},
	}
}

// EmailButtons encapsulates the email button resource.
func (gb *Gobilla) EmailButtons() l4e.EmailButtons {
	return l4e.EmailButtons{
		Resource: resource.Resource{
			Auth: gb.auth,
		},
	}
}
