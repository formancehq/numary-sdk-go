package ledgerclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

const (
	AuthEndpoint = "https://api.numary.cloud/auth/authenticate/tokens"
)

type cloudRoundTripper struct {
	client        *http.Client
	rt            http.RoundTripper
	url           string
	token         string
	resolvedToken string
}

func (c *cloudRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	if c.resolvedToken != "" {
		return c.rt.RoundTrip(request)
	}
	rsp, err := c.client.Post(c.url, "application/json", bytes.NewBufferString(fmt.Sprintf(`{"strategy": "m2m", "token": "%s"}`, c.token)))
	if err != nil {
		return nil, errors.Wrap(err, "getting token")
	}

	type Token struct {
		Data struct {
			JWT string `json:"jwt"`
		} `json:"data"`
	}
	tok := &Token{}
	err = json.NewDecoder(rsp.Body).Decode(tok)
	if err != nil {
		return nil, errors.Wrap(err, "decoding auth server response")
	}
	if request.Header == nil {
		request.Header = http.Header{}
	}
	request.Header.Set("Authorization", "Bearer "+tok.Data.JWT)
	return c.rt.RoundTrip(request)
}

var _ http.RoundTripper = &cloudRoundTripper{}

func WithCloudClient(client *http.Client, cloudUrl, token string) *http.Client {
	newClient := *client
	if newClient.Transport == nil {
		newClient.Transport = http.DefaultTransport
	}
	newClient.Transport = &cloudRoundTripper{
		client: client,
		rt:     newClient.Transport,
		url:    cloudUrl,
		token:  token,
	}
	return &newClient
}

func FetchToken(client *http.Client, cloudUrl, token string) (string, error) {
	rsp, err := client.Post(cloudUrl, "application/json", bytes.NewBufferString(fmt.Sprintf(`{"strategy": "m2m", "token": "%s"}`, token)))
	if err != nil {
		return "", errors.Wrap(err, "getting token")
	}

	type Token struct {
		Data struct {
			JWT string `json:"jwt"`
		} `json:"data"`
	}
	tok := &Token{}
	err = json.NewDecoder(rsp.Body).Decode(tok)
	if err != nil {
		return "nil", errors.Wrap(err, "decoding auth server response")
	}
	return tok.Data.JWT, nil
}
