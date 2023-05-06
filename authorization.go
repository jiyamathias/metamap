package metamap

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

// Authorization method makes a request to metemap and returns a JWT authorization token to access the rest of metamap services
func (c *Client) Authorization() (AuthorizationResponse, error) {
	url := c.BaseUrl + "oauth"
	method := MethodPOST
	payload := AuthorizationRequest{GrantType: "client_credentials"}

	var response AuthorizationResponse
	if err := c.newRequest(method, url, payload, &response); err != nil {
		return AuthorizationResponse{}, err
	}

	// set the  getenerated jwt accestoken to the AccessToken field in the client struct
	c.AccessToken = response.Payload.User.ID
	return response, nil
}
