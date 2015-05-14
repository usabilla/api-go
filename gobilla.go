package gobilla

// Gobilla is the client that exposes all the resources of the Usabilla API.
type Gobilla struct {
	auth auth
}

// New creates a new Gobilla instance and sets the auth with key and secret.
func New(key, secret string) *Gobilla {
	return &Gobilla{
		auth: auth{
			key:    key,
			secret: secret,
		},
	}
}

// Buttons encapsulates the button resource.
func (gb *Gobilla) Buttons() Buttons {
	return Buttons{
		resource: resource{
			auth: gb.auth,
		},
	}
}

// Campaigns encapsulates the campaign resource.
func (gb *Gobilla) Campaigns() Campaigns {
	return Campaigns{
		resource: resource{
			auth: gb.auth,
		},
	}
}

// Apps encapsulates the app resource.
func (gb *Gobilla) Apps() Apps {
	return Apps{
		resource: resource{
			auth: gb.auth,
		},
	}
}

// EmailButtons encapsulates the email button resource.
func (gb *Gobilla) EmailButtons() EmailButtons {
	return EmailButtons{
		resource: resource{
			auth: gb.auth,
		},
	}
}
