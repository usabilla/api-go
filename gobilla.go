package gobilla

// Gobilla is the client that exposes all the resources of the Usabilla API.
type Gobilla struct {
	Key, Secret string
}

// Buttons encapsulate the button resource.
func (gb *Gobilla) Buttons() Buttons {
	au := auth{
		key:    gb.Key,
		secret: gb.Secret,
	}
	return Buttons{
		resource: resource{
			auth: au,
			uri:  buttonURI,
		},
	}
}

// Campaigns encapsulate the campaign resource.
func (gb *Gobilla) Campaigns() Campaigns {
	au := auth{
		key:    gb.Key,
		secret: gb.Secret,
	}
	return Campaigns{
		resource: resource{
			auth: au,
			uri:  campaignURI,
		},
	}
}
