package gobilla

var (
	key    string
	secret string
)

/*
API is the client that exposes all the resources of the Usabilla API.
*/
type API struct {
	Buttons   Buttons
	Campaigns Campaigns
}

/*
NewAPI is used to create the API and set a key and secret.
*/
func NewAPI(k, s string) *API {
	key = k
	secret = s
	return &API{}
}
