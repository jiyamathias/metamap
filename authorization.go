package metamap

import "time"

type (
	// AuthorizationRequest is the authorization request model
	AuthorizationRequest struct {
		GrantType string `json:"grant_type"` // client_credentials
	}

	// AuthorizationResponse is the authorization response model
	AuthorizationResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		Payload     struct {
			User struct {
				ID string `json:"_id"`
			} `json:"user"`
		} `json:"payload"`
	}
)

// Authorization makes a request to metemap and returns a JWT authorization token to access the rest of metamap services
func (c *Client) Authorization() (*AuthorizationResponse, error) {
	url := c.BaseUrl + "oauth"
	method := MethodPOST
	payload := AuthorizationRequest{GrantType: "client_credentials"}

	var response AuthorizationResponse
	if err := c.newRequest(method, url, payload, &response); err != nil {
		return &AuthorizationResponse{}, err
	}

	// set the  getenerated jwt accestoken to the AccessToken field in the client struct
	c.AccessToken = response.Payload.User.ID
	c.TokenValidity = time.Now() //save the current time to the TimeValidity field

	return &response, nil
}

/*
VerifyTokenValidity is a metjod used to verify the JWT token validity

Since the JWT token MetaMap gives to have access to it's resources is only valid for 1 hour. The VerifyTokenValidity method
verifies the time that the JWT token was created against the the current time to confirm if the its already been a hour or not.
If it has, the VerifyTokenValidity method makes a request to MetaMap to get a new JWT token using  client.Authorization method.
*/
func (c *Client) VerifyTokenValidity() error {
	t := c.TokenValidity
	// we use 55 minutes as the token validity time to cover up for the time it took to assign a new time to the client.TokenValidity field to avloid errors of any sort
	fiftyFiveMinsAgo := t.Add(-55 * time.Minute)
	if time.Now().After(fiftyFiveMinsAgo) {
		// the token have expires, make a new token request
		resp, err := c.Authorization() // just calling thus method requests a new JWT token from MetaMap and assings it to the client.TokenValidity field
		if err != nil {
			return err
		}
		// assign the newly generated JWT token to the client.AccessToken field and the current tiime to the client.TokenValidity field
		c.AccessToken = resp.Payload.User.ID
		c.TokenValidity = time.Now()
	}

	//the token is still valid and so user the aleady saved token
	return nil
}
