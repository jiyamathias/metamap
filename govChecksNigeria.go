package metamap

type (

	//National ID
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

	//Voters Card
	NigeriaVotingIDNumberRequest struct {
		DocumentNumber string `json:"documentNumber"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		DateOfBirth    string `json:"dateOfBirth"`
		CallbackUrl    string `json:"callbackUrl"`
	}

	NigeriaVotingIDNumberResponse struct {
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

	//Drivers Licence
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

	//Corporate Affairs Commission
	NigeriaCACRequest struct {
		RegistrationNumber string `json:"registrationNumber"`
		CallbackUrl        string `json:"callbackUrl"`
	}

	NigeriaCACResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			Type             string `json:"type"`
			CompanyName      string `json:"companyName"`
			CacNumber        string `json:"cacNumber"`
			Status           string `json:"status"`
			CompanyAddress   string `json:"companyAddress"`
			CompanyEmail     string `json:"companyEmail"`
			RegistrationDate string `json:"registrationDate"`
		} `json:"data"`
	}

	// Tax Identification Number
	NigeriaTaxIdentificationNumberRequest struct {
		TaxNumber   string `json:"taxNumber"`
		CallbackUrl string `json:"callbackUrl"`
	}

	NigeriaTaxIdentificationNumberResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			CompanyName  string `json:"companyName"`
			FirsNumber   string `json:"FirsNumber"`
			CacNumber    string `json:"cacNumber"`
			JtbNumber    string `json:"JtbNumber"`
			TaxOffice    string `json:"taxOffice"`
			CompanyPhone string `json:"companyPhone"`
			CompanyEmail string `json:"companyEmail"`
		} `json:"data"`
	}

	//Corporate Affairs Commission Affilaite
	NigeriaCACAffiliatesRequest struct {
		RegistrationNumber string `json:"registrationNumber"`
		CallbackUrl        string `json:"callbackUrl"`
	}

	NigeriaCACAffiliatesResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			Shareholders []struct {
				Name                string `json:"name"`
				Position            string `json:"position"`
				Status              string `json:"status"`
				DateOfBirth         string `json:"dateOfBirth"`
				PhoneNumber         string `json:"phoneNumber"`
				Email               string `json:"email"`
				City                string `json:"city"`
				Address             string `json:"address"`
				IdType              string `json:"idType"`
				IdNumber            string `json:"idNumber"`
				Shares              string `json:"shares"`
				AccreditationNumber string `json:"accreditationNumber"`
			} `json:"shareholders"`
		} `json:"data"`
	}

	// Bank Verification Number
	NigeriaBVNRequest struct {
		DocumentNumber string `json:"documentNumber"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		CallbackUrl    string `json:"callbackUrl"`
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
NigeriaVotingIDNumber verifies Voting Identity Number.

This method takes in the NigeriaVotingIDNumberRequest{} struct as a parameter.

NOTE: The format for the DateOfBirth is yyyy-mm-dd.
*/
func (c *Client) NigeriaVotingIDNumber(req NigeriaVotingIDNumberRequest) (*NigeriaVotingIDNumberResponse, error) {
	url := "govchecks/v1/ng/vin"
	method := MethodPOST
	var response NigeriaVotingIDNumberResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &NigeriaVotingIDNumberResponse{}, err
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

/*
NigeriaCAC verifies a company's CAC (Corporate Affairs Commission) number.

This method takes in the NigeriaCACRequest{} struct as a parameter.
*/
func (c *Client) NigeriaCAC(req NigeriaCACRequest) (*NigeriaCACResponse, error) {
	url := "govchecks/v1/ng/cac"
	method := MethodPOST
	var response NigeriaCACResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &NigeriaCACResponse{}, err
	}

	return &response, nil
}

/*
NigeriaTaxIdentificationNumber verifies a company's Tax Identification Number.

This method takes in the NigeriaTaxIdentificationNumberRequest{} struct as a parameter.
*/
func (c *Client) NigeriaTaxIdentificationNumber(req NigeriaTaxIdentificationNumberRequest) (*NigeriaTaxIdentificationNumberResponse, error) {
	url := "govchecks/v1/ng/tin"
	method := MethodPOST
	var response NigeriaTaxIdentificationNumberResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &NigeriaTaxIdentificationNumberResponse{}, err
	}

	return &response, nil
}

/*
NigeriaCACAffiliates verifies a major shareholder names and titles if a company's Corporate Affairs Commission number is valid.

This method takes in the NigeriaCACAffiliatesRequest{} struct as a parameter.
*/
func (c *Client) NigeriaCACAffiliates(req NigeriaCACAffiliatesRequest) (*NigeriaCACAffiliatesResponse, error) {
	url := "govchecks/v1/ng/cac-affiliates"
	method := MethodPOST
	var response NigeriaCACAffiliatesResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &NigeriaCACAffiliatesResponse{}, err
	}

	return &response, nil
}

/*
NigeriaBVN verifies a BVN (Bank Verification Number).

This method takes in the NigeriaBVNRequest{} struct as a parameter.
*/
func (c *Client) NigeriaBVN(req NigeriaBVNRequest) (*interface{}, error) {
	url := "govchecks/v1/ng/bvn"
	method := MethodPOST
	var response interface{}
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return nil, err
	}

	return &response, nil
}
