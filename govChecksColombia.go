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

	// ColombiaUnifiedLegalSearch
	ColombiaUnifiedLegalSearchRequest struct {
		FullName    string `json:"fullName"`
		CallbackUrl string `json:"callbackUrl"`
	}

	ColombiaUnifiedLegalSearchResponse struct {
		Status int    `json:"status"`
		Id     string `json:"id"`
		Error  struct {
			Type    string `json:"type"`
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
		Data struct {
			Results []struct {
				IdProceso            interface{} `json:"idProceso"`
				LlaveProceso         string      `json:"llaveProceso"`
				FechaProceso         string      `json:"fechaProceso"`
				FechaUltimaActuacion string      `json:"fechaUltimaActuacion"`
				Despacho             string      `json:"despacho"`
				Departamento         string      `json:"departamento"`
				SujetosProcesales    string      `json:"sujetosProcesales"`
			} `json:"results"`
		} `json:"data"`
		Timestamp string `json:"timestamp"`
	}

	// ColombiaRUES
	ColombiaRUESRequest struct {
		Nit         string `json:"nit"`
		CallbackUrl string `json:"callbackUrl"`
	}
	ColombiaRUESResponse struct {
		Status int         `json:"status"`
		Error  interface{} `json:"error"`
		Data   struct {
			CompanyName               string      `json:"companyName"`
			Nit                       string      `json:"nit"`
			ShortName                 interface{} `json:"shortName"`
			Municipality              string      `json:"municipality"`
			Category                  string      `json:"category"`
			Status                    string      `json:"status"`
			EnrollmentNumber          string      `json:"enrollmentNumber"`
			LastRenewal               string      `json:"lastRenewal"`
			RenewalDate               string      `json:"renewalDate"`
			EnrollmentDate            string      `json:"enrollmentDate"`
			ExpirationDate            string      `json:"expirationDate"`
			EnrollmentStatus          string      `json:"enrollmentStatus"`
			AnnullmentReason          string      `json:"annullmentReason"`
			CompanyType               string      `json:"companyType"`
			OrganizationType          string      `json:"organizationType"`
			EnrollmentCategory        string      `json:"enrollmentCategory"`
			LastUpdate                string      `json:"lastUpdate"`
			EconomicActivities        []string    `json:"economicActivities"`
			MainLegalRepresentative   string      `json:"mainLegalRepresentative"`
			OtherLegalRepresentatives []string    `json:"otherLegalRepresentatives"`
			Entities                  []struct {
				EntityName       string `json:"entityName"`
				EnrollmentNumber string `json:"enrollmentNumber"`
				EnrollmentDate   string `json:"enrollmentDate"`
			} `json:"entities"`
		} `json:"data"`
		Timestamp string `json:"timestamp"`
	}

	// ColombiaPPT
	ColombiaPPTRequest struct {
		Rumv        string `json:"rumv"`        // User's RUMV number
		Dni         string `json:"dni"`         // User's DNI number
		DateOfBirth string `json:"dateOfBirth"` // Date of Birth format is DD-MM-YYYY
		CallbackUrl string `json:"callbackUrl"` // Callback URL to receive request results.
	}
	ColombiaPPTResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			FullName     string `json:"fullName"`
			Phone        string `json:"phone"`
			Email        string `json:"email"`
			DocumentType string `json:"documentType"`
			Dni          string `json:"dni"`
			Status       string `json:"status"`
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

/*
ColombiaUnifiedLegalSearch check a user's full name against Colombian police records.

This method takes in the ColombiaUnifiedLegalSearchRequest{} struct as a parameter.

MetaMap searches through Colombian police records for a user's past or current warrants or arrests cases.
*/
func (c *Client) ColombiaUnifiedLegalSearch(req ColombiaUnifiedLegalSearchRequest) (*ColombiaUnifiedLegalSearchResponse, error) {
	url := "govchecks/v1/co/unified-criminal-search"
	method := MethodPOST
	var response ColombiaUnifiedLegalSearchResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ColombiaUnifiedLegalSearchResponse{}, err
	}

	return &response, nil
}

/*
ColombiaRUES check a business's national tax ID against Colombian Single Business and Social Registry.

This method takes in the ColombiaRUESRequest{} struct as a parameter.

MetaMap searches through the Colombian Single Business and Social Registry (Registro Unico Empresarial y Social, RUES) to validate a business's national tax ID (Número de Identificación Tributaria, NIT).
*/
func (c *Client) ColombiaRUES(req ColombiaRUESRequest) (*ColombiaRUESResponse, error) {
	url := "govchecks/v1/co/rues"
	method := MethodPOST
	var response ColombiaRUESResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ColombiaRUESResponse{}, err
	}

	return &response, nil
}

/*
ColombiaPPT verify the validity of the Colombian PPT (Permiso por protección temporal = Temporary Protection Permit).

This method takes in the ColombiaPPTRequest{} struct as a parameter.

NOTE: Date of Birth format is DD-MM-YYYY

MetaMap searches through the Colombian PPT registry to validate the Permiso por protección temporal document against the RUMV Number, the Identity Document Number, and the date of birth.
*/
func (c *Client) ColombiaPPT(req ColombiaPPTRequest) (*ColombiaPPTResponse, error) {
	url := "govchecks/v1/co/ppt"
	method := MethodPOST
	var response ColombiaPPTResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ColombiaPPTResponse{}, err
	}

	return &response, nil
}
