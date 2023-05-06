package metamap

type (

	//national ID
	NigeriaNINRequest struct {
		//VNIN is a required parameter that must be passed in
		VNIN        string `json:"vNIN"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		DateOfBirth string `json:"dateOfBirth"`
		//Metadata is an optional parameter
		Metadata map[string]interface{} `json:"metadata"`
		//CallbackUrl is a required parameter that must be passed in
		CallbackUrl string `json:"callbackUrl"`
	}

	NigeriaNINResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			ID          string `json:"id"`
			FirstName   string `json:"firstName"`
			LastName    string `json:"lastName"`
			DateOfBirth string `json:"dateOfBirth"`
			Gender      string `json:"gender"`
			Phone       string `json:"phone"`
			VNin        string `json:"vNin"`
			Photo       string `json:"photo"`
		} `json:"data"`
	}

	//voters card
	NigeriaVotingIDRequest struct {
		DocumentNumber string `json:"documentNumber"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		DateOfBirth    string `json:"dateOfBirth"`
		CallbackUrl    string `json:"callbackUrl"`
	}

	NigeriaVotingIDResponse struct {
		Status string      `json:"status"`
		ID     string      `json:"id"`
		Error  interface{} `json:"error"`
		Data   struct {
			FirstName     string `json:"firstName"`
			LastName      string `json:"lastName"`
			MiddleName    string `json:"middlename"`
			Gender        string `json:"gender"`
			Occupation    string `json:"occupation"`
			StateOfOrigin string `json:"stateOfOrigin"`
			LgaOfOrigin   string `json:"lgaOfOrigin"`
			OnSpokenLang  string `json:"onspokenlang"`
			PollingUnit   string `json:"pollingUnit"`
		} `json:"data"`
		Timestamp string `json:"timestamp"`
	}

	//drivers licence
	NigeriaDriversLicenceRequest struct {
		DocumentNumber string `json:"documentNumber"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		CallbackUrl    string `json:"callbackUrl"`
	}

	NigeriaDriversLicenceResponse struct {
		Status string      `json:"status"`
		ID     string      `json:"id"`
		Error  interface{} `json:"error"`
		Data   struct {
			LicenseNo           string `json:"license_no"`
			FirstName           string `json:"firstName"`
			LastName            string `json:"lastName"`
			IssuedDate          string `json:"issuedDate"`
			ExpirationDate      string `json:"expirationDate"`
			StateOfIssue        string `json:"stateOfIssue"`
			Gender              string `json:"gender"`
			Birthdate           string `json:"birthdate"`
			MiddleName          string `json:"middlename"`
			GovernmentFaceMatch struct {
				Photo   string `json:"photo"`
				Matched bool   `json:"matched"`
				Score   int    `json:"score"`
			} `json:"governmentFaceMatch"`
		} `json:"data"`
		Timestamp string `json:"timestamp"`
	}
)

/*
NigeriaVirtualNIN verifies NIN(National Identification Number).

This method takes in the NigeriaNINRequest{} struct as a parameter. Both the VNIN  and CallbackUrl field are required parameter that must be passed in for the request.

NOTE: The format for the DateOfBirth is yyyy-mm-dd.

It also takes in an optional parameter `metadata` which is a map[string]interface that is passed in to add custom data to the request.
*/
func (c *Client) NigeriaVirtualNIN(req NigeriaNINRequest) (*NigeriaNINResponse, error) {
	url := "govchecks/v1/ng/vnin"
	method := MethodPOST
	var response NigeriaNINResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &NigeriaNINResponse{}, err
	}

	return &response, nil
}

/*
NigeriaVotingID verifies Voting Identity Number.

This method takes in the NigeriaVotingIDRequest{} struct as a parameter.

NOTE: The format for the DateOfBirth is yyyy-mm-dd.
*/
func (c *Client) NigeriaVotingID(req NigeriaVotingIDRequest) (*NigeriaVotingIDResponse, error) {
	url := "govchecks/v1/ng/vin"
	method := MethodPOST
	var response NigeriaVotingIDResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &NigeriaVotingIDResponse{}, err
	}

	return &response, nil
}

/*
NigeriaDriversLicence verifies drivers licence.

This method takes in the NigeriaDriversLicenceRequest{} struct as a parameter.
*/
func (c *Client) NigeriaDriversLicence(req NigeriaDriversLicenceRequest) (*NigeriaDriversLicenceResponse, error) {
	url := "govchecks/v1/ng/dl"
	method := MethodPOST
	var response NigeriaDriversLicenceResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &NigeriaDriversLicenceResponse{}, err
	}

	return &response, nil
}
