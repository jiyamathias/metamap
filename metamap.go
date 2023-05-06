package metamap

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	//Http methods for the http request
	MethodPOST = "POST"
	MethodGET  = "GET"
)

// Client is the struct holding all of the variables to be used
type Client struct {
	Http         *http.Client
	AccessToken  string
	AuthToken    string
	ClientId     string
	ClientSecret string
	BaseUrl      string
}

// Endocode takes in a string and encoded it into a base64 string
func Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// New is a method that takes in several parameters in other to initiate the metamap configs
func New(h *http.Client, clientId, clientSecret string) *Client {
	baseUrl := "https://api.getmati.com/"
	//setup the authrozation token
	token := fmt.Sprintf("%s:%s", clientId, clientSecret)
	encodedToken := Encode(token)
	newToken := fmt.Sprintf("Basic %s", encodedToken)

	return &Client{
		Http:         h,
		AuthToken:    newToken,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		BaseUrl:      baseUrl,
	}
}

// newRequest method makes a http request to the metamap server and decodes the server response into the resp(esponse) parameter passed into the newRequest method
func (c *Client) newRequest(method, reqURL string, reqBody interface{}, resp interface{}) error {
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return errors.Wrap(err, "http client ::: unable to marshal request struct")
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, reqURL, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", c.AuthToken)
	if err != nil {
		return errors.Wrap(err, "http client ::: unable to create request body")
	}

	res, err := c.Http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client ::: client failed to execute request")
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != 204 {
		return errors.Errorf("http client ::: invalid status code received, expected 200/204, got %v", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "http client ::: unable to unmarshal response body")
	}

	return nil
}
