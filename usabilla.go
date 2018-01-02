/*
Copyright (c) 2018 Usabilla

Permission is hereby granted, free of charge, to any person obtaining a
copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish, dis-
tribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the fol-
lowing conditions:

The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABIL-
ITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT
SHALL THE AUTHOR BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
IN THE SOFTWARE.
*/

// Package usabilla provides a wrapper around Usabilla Public API.
//
// https://usabilla.com/api
package usabilla

import "net/http"

// Usabilla is the client that exposes all the resources of the Usabilla API.
// You can provide a custom http client to change the way the client works.
type Usabilla struct {
	auth   auth
	Client *http.Client
}

// New creates a new Usabilla instance and sets the auth with key and secret.
// Client is the default http client. To change the way the client works
// provide a custom http client. Passing nil will use the http.DefaultClient
func New(key, secret string, customClient *http.Client) *Usabilla {
	client := http.DefaultClient
	if customClient != nil {
		client = customClient
	}
	return &Usabilla{
		auth: auth{
			key:    key,
			secret: secret,
		},
		Client: client,
	}
}

// Buttons encapsulates the button resource.
func (usabilla *Usabilla) Buttons() *Buttons {
	return &Buttons{
		resource: resource{
			auth: usabilla.auth,
		},
		client: usabilla.Client,
	}
}

// Campaigns encapsulates the campaign resource.
func (usabilla *Usabilla) Campaigns() *Campaigns {
	return &Campaigns{
		resource: resource{
			auth: usabilla.auth,
		},
		client: usabilla.Client,
	}
}

// Apps encapsulates the app resource.
func (usabilla *Usabilla) Apps() *Apps {
	return &Apps{
		resource: resource{
			auth: usabilla.auth,
		},
		client: usabilla.Client,
	}
}

// AppsCampaigns encapsulates the AppCampaigns resource
func (usabilla *Usabilla) AppCampaigns() *AppCampaigns {
	return &AppCampaigns{
		resource: resource{
			auth: usabilla.auth,
		},
		client: usabilla.Client,
	}
}

// EmailButtons encapsulates the email button resource.
func (usabilla *Usabilla) EmailButtons() *EmailButtons {
	return &EmailButtons{
		resource: resource{
			auth: usabilla.auth,
		},
		client: usabilla.Client,
	}
}

// Inpage encapsulates the app resource.
func (usabilla *Usabilla) InpageWidgets() *InpageWidgets {
	return &InpageWidgets{
		resource: resource{
			auth: usabilla.auth,
		},
		client: usabilla.Client,
	}
}
