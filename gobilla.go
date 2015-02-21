package gobilla

var (
	key    string
	secret string
)

// Client is the client that exposes all the resources of the Usabilla API.
type Client struct {
	Buttons   Buttons
	Campaigns Campaigns
}

func (client *Client) Key(k string) {
	key = k
}

func (client *Client) Secret(s string) {
	secret = s
}

// NewClient creates a new Client instance and sets the key and secret.
func NewClient(k, s string) Client {
	key = k
	secret = s
	return Client{}
}
