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
NewAPI is an API factory, which also sets the key and secret.
*/
func NewAPI(k, s string) *API {
	key = k
	secret = s
	return &API{}
}
