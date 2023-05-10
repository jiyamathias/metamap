package metamap

type (

	// ArgentinaDNI
	ArgentinaDNIRequest struct {
		// Document number from either a National ID or a driver license.
		DocumentNumber string `json:"documentNumber"`
		// "2023-05-10"
		DateOfBirth string `json:"dateOfBirth"`
		// M for male, F for female.
		Gender string `json:"gender"`
		// Document issue date.
		DateOfIssue string `json:"dateOfIssue"`
		CallbackUrl string `json:"callbackUrl"`
		// Use Metadata to add internal references for your outputs/webhooks.
		Metadata map[string]interface{} `json:"metadata"`
	}

	ArgentinaDNIResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			DateOfIssue    string `json:"dateOfIssue"`
			DocumentNumber string `json:"documentNumber"`
		} `json:"data"`
	}

	// ArgentinaRENAPER
	ArgentinaRENAPERRequest struct {
		// Document number from either a National ID or a driver license.
		DocumentNumber string `json:"documentNumber"`
		// "2023-05-10"
		DateOfBirth string `json:"dateOfBirth"`
		// M for male, F for female.
		Gender string `json:"gender"`
		// Document issue date.
		DateOfIssue string `json:"dateOfIssue"`
		CallbackUrl string `json:"callbackUrl"`
		// Use Metadata to add internal references for your outputs/webhooks.
		Metadata map[string]interface{} `json:"metadata"`
	}

	ArgentinaRENAPERResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			DateOfBirth  string      `json:"dateOfBirth"`
			FullName     string      `json:"fullName"`
			DniNumber    string      `json:"dniNumber"`
			Cuit         string      `json:"cuit"`
			PhoneNumbers []string    `json:"phoneNumbers"`
			Deceased     interface{} `json:"deceased"`
		} `json:"data"`
	}
)

/*
ArgentinaDNI verify a user's DNI (Documento Nacional de Identidad) based on card number and issue date.

This method takes in the ArgentinaDNIRequest{} struct as a parameter.
*/
func (c *Client) ArgentinaDNI(req ArgentinaDNIRequest) (*ArgentinaDNIResponse, error) {
	url := "govchecks/v1/ar/dni"
	method := MethodPOST
	var response ArgentinaDNIResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ArgentinaDNIResponse{}, err
	}

	return &response, nil
}

/*
ArgentinaRENAPER verify a user's DNI (Documento Nacional de Identidad) number and identity.

This method takes in the ArgentinaRENAPERRequest{} struct as a parameter.
*/
func (c *Client) ArgentinaRENAPER(req ArgentinaRENAPERRequest) (*ArgentinaRENAPERResponse, error) {
	url := "govchecks/v1/ar/renaper"
	method := MethodPOST
	var response ArgentinaRENAPERResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ArgentinaRENAPERResponse{}, err
	}

	return &response, nil
}
