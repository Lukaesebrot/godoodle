package godoodle

import "github.com/valyala/fasthttp"

// Client represents an API client
type Client struct {
	http         *fasthttp.Client
	clientID     string
	clientSecret string
}

// New creates a new API client
func New(clientID, clientSecret string) *Client {
	return &Client{
		http: &fasthttp.Client{
			Name: "godoodle",
		},
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}
