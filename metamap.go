package metamap

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

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
	/*
		IsBasic is used to determine what kind of request is being being made so that the request header  can be set accordingly.

		This value should be set to true only for if the request header authorization is Basic and for Bearer it should be set to false
	*/
	IsBasic bool

	TokenValidity time.Time // metamap JWT token is only valid for 1hr and afterwards a new JWT token needs to be created intothet to access the resource
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

	return &Client{
		Http:         h,
		AuthToken:    encodedToken,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		BaseUrl:      baseUrl,
	}
}

/*
newRequest makes a http request to the metamap server and decodes the server response into the resp(esponse) parameter passed into the newRequest method

Without the JWT access token, we can't make request to access the reest of MetaMap api and so before making any request, the JWT access token is first verified using the client.VerifyTokenValidity method
to confirm the validity of the JWT token if it has not expired yet and if it has expired, it request for a new JWT token.
*/
func (c *Client) newRequest(method, reqURL string, reqBody interface{}, resp interface{}) error {
	//JWT access token is first verified before making any request
	if err := c.VerifyTokenValidity(); err != nil {
		return err
	}

	newURL := c.BaseUrl + reqURL
	var body io.Reader

	if reqBody != nil {
		bb, err := json.Marshal(reqBody)
		if err != nil {
			return errors.Wrap(err, "http client ::: unable to marshal request struct")
		}
		body = bytes.NewReader(bb)
	}

	req, err := http.NewRequest(method, newURL, body)
	req.Header.Set("Content-Type", "application/json")
	if c.IsBasic {
		authToken := fmt.Sprintf("Basic %s", c.AuthToken)
		req.Header.Set("Authorization", authToken)
	} else {
		accessToken := fmt.Sprintf("Bearer %s", c.AccessToken)
		req.Header.Set("Authorization", accessToken)
	}

	if err != nil {
		return errors.Wrap(err, "http client ::: unable to create request body")
	}

	res, err := c.Http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http client ::: client failed to execute request")
	}

	if res.StatusCode != http.StatusOK {
		return errors.Errorf("http client ::: invalid status code received, expected 200, got %v", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return errors.Wrap(err, "http client ::: unable to unmarshal response body")
	}

	return nil
}
