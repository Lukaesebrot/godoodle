package godoodle

import (
	"encoding/json"
	"errors"

	"github.com/valyala/fasthttp"
)

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

// CreditsSpent checks how much credits the current client already used
func (client *Client) CreditsSpent() (int64, error) {
	// Acquire a request object
	request := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(request)

	// Acquire a response object
	response := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(response)

	// Prepare the request object
	request.SetRequestURI(EndpointCreditsSpent)
	request.Header.SetMethod("POST")
	request.Header.SetContentType("application/json")
	request.PostArgs().Set("clientId", client.clientID)
	request.PostArgs().Set("clientSecret", client.clientSecret)

	// Perform the request
	err := client.http.Do(request, response)
	if err != nil {
		return -1, err
	}

	// Parse the response into a json struct
	var responseStruct struct {
		Used  int64  `json:"used,omitempty"`
		Error string `json:"error,omitempty"`
	}
	err = json.Unmarshal(response.Body(), &responseStruct)
	if err != nil {
		return -1, err
	}

	// Return the corresponding values
	if responseStruct.Error != "" {
		return -1, errors.New(responseStruct.Error)
	}
	return responseStruct.Used, nil
}
