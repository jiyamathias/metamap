package metamap

type (
	// ColombiaMigrationInstitute
	ColombiaMigrationInstituteRequest struct {
		DocumentNumber string `json:"documentNumber"`
		DateOfIssue    string `json:"dateOfIssue"`
		CallbackUrl    string `json:"callbackUrl"`
	}

	ColombiaMigrationInstituteResponse struct {
		Status int         `json:"status"`
		Error  interface{} `json:"error"`
		Data   struct {
			FullName        string `json:"fullName"`
			DateOfBirth     string `json:"dateOfBirth"`
			Nationality     string `json:"nationality"`
			Status          string `json:"status"`
			DateOfExpiry    string `json:"dateOfExpiry"`
			DateOfIssue     string `json:"dateOfIssue"`
			CertificateCode string `json:"certificateCode"`
		} `json:"data"`
		Timestamp string `json:"timestamp"`
	}

	// ColombiaRegistraduria
	ColombiaRegistraduriaRequest struct {
		DocumentNumber string `json:"documentNumber"`
		CallbackUrl    string `json:"callbackUrl"`
	}

	ColombiaRegistraduriaResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			AgeRange                     string `json:"ageRange"`
			BankAccountsCount            string `json:"bankAccountsCount"`
			CommercialIndustryDebtsCount string `json:"commercialIndustryDebtsCount"`
			DocumentNumber               string `json:"documentNumber"`
			EmissionDate                 string `json:"emissionDate"`
			FinancialIndustryDebtsCount  string `json:"financialIndustryDebtsCount"`
			FullName                     string `json:"fullName"`
			Gender                       string `json:"gender"`
			IssuePlace                   string `json:"issuePlace"`
			Name                         string `json:"name"`
			MiddleName                   string `json:"middleName"`
			Surname                      string `json:"surname"`
			SecondSurname                string `json:"secondSurname"`
			SavingAccountsCount          string `json:"savingAccountsCount"`
			SolidarityIndustryDebtsCount string `json:"solidarityIndustryDebtsCount"`
			ServiceIndustryDebtsCount    string `json:"serviceIndustryDebtsCount"`
		} `json:"data"`
	}

	// ColombiaRegistraduriaLight
	ColombiaRegistraduriaLightRequest struct {
		DocumentNumber string `json:"documentNumber"` // Document number for a National ID
		DateOfIssue    string `json:"dateOfIssue"`    // Document issue date
		CallbackUrl    string `json:"callbackUrl"`
	}

	ColombiaRegistraduriaLightResponse struct {
		Status int         `json:"status"`
		Error  interface{} `json:"error"`
		Id     string      `json:"id"`
		Data   struct {
			FullName       string `json:"fullName"`
			DocumentNumber string `json:"documentNumber"`
			EmissionDate   string `json:"emissionDate"`
			PlaceOfIssue   string `json:"placeOfIssue"`
			Status         string `json:"status"`
		} `json:"data"`
	}
)

/*
ColombiaMigrationInstitute verify that a user has a valid immigration status in Colombia.

This method takes in the ColombiaMigrationInstituteRequest{} struct as a parameter.

NOTE:  DateOfIssue fieid should be in this format 2022-01-01.

MetaMap connects with Colombia Migration (Migración Colombia) to validate that the user's name and ID issue date match the information associated with their residence permit.
*/
func (c *Client) ColombiaMigrationInstitute(req ColombiaMigrationInstituteRequest) (*ColombiaMigrationInstituteResponse, error) {
	url := "govchecks/v1/co/migraciones"
	method := MethodPOST
	var response ColombiaMigrationInstituteResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ColombiaMigrationInstituteResponse{}, err
	}

	return &response, nil
}

/*
ColombiaRegistraduria verify a user's ID card against the Colombian Civil Registry.

This method takes in the ColombiaRegistraduriaRequest{} struct as a parameter.

MetaMap connects with the Colombian Civil Registry (Resgistraduría Nacional del Estado Civil) to validate that the Cédula Number (Rol Único Nacional) and other data present in the ID card corresponds to their database. When validated, this API will also return the user's financial information, including loan counts and bank accounts counts.
*/
func (c *Client) ColombiaRegistraduria(req ColombiaRegistraduriaRequest) (*ColombiaRegistraduriaResponse, error) {
	url := "govchecks/v1/co/registraduria"
	method := MethodPOST
	var response ColombiaRegistraduriaResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ColombiaRegistraduriaResponse{}, err
	}

	return &response, nil
}

/*
ColombiaRegistraduriaLight verify a user's ID card against the Colombian Civil Registry.

This method takes in the ColombiaRegistraduriaLightRequest{} struct as a parameter.

NOTE:  DateOfIssue fieid should be in this format 2022-01-01.

MetaMap connects with the Colombian Civil Registry (Resgistraduría Nacional del Estado Civil) to validate that the Cédula Number (Rol Único Nacional) and other data present in the ID card corresponds to their database and is valid.
*/
func (c *Client) ColombiaRegistraduriaLight(req ColombiaRegistraduriaLightRequest) (*ColombiaRegistraduriaLightResponse, error) {
	url := "govchecks/v1/co/registraduria-light"
	method := MethodPOST
	var response ColombiaRegistraduriaLightResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ColombiaRegistraduriaLightResponse{}, err
	}

	return &response, nil
}
