package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/composer22/docker-deploy-server/server"
)

// Client represents an instance of a connection to the server.
type Client struct {
	Token string `json:"bearerToken"` // The API authorization token to the server.
	Url   string `json:"URL"`         // The URL to the docker-deploy-server endpoint.
}

// New is a factory function that returns a new client instance.
func New(token string, url string) *Client {
	return &Client{
		Token: token,
		Url:   url,
	}
}

// Version prints the version of the client then exits.
func Version() string {
	return fmt.Sprintf("%s version %s\n", applicationName, version)
}

// Deploy submits a deploy request to the server.
func (c *Client) Deploy(repository string, tag string, environment string) (string, error) {
	// Create the payload.

	dr := &server.DeployRequest{
		ImageName:   repository,
		ImageTag:    tag,
		Environment: environment,
	}
	payload, err := json.Marshal(dr)
	if err != nil {
		return "", err
	}

	// Send the request.
	req, err := http.NewRequest(httpPost, fmt.Sprintf("%s%s", c.Url, httpRouteV1Deploy),
		bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return "", err
	}
	return c.sendRequest(req)
}

// Status prints out the status of a previous deploy.
func (c *Client) Status(deployID string) (string, error) {
	req, err := http.NewRequest(httpGet, fmt.Sprintf("%s%s%s", c.Url, httpRouteV1Status, deployID), nil)
	if err != nil {
		return "", err
	}
	return c.sendRequest(req)
}

// sendRequest sends a request to the server and prints the result.
func (c *Client) sendRequest(req *http.Request) (string, error) {
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	cl := &http.Client{}
	resp, err := cl.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
