package metamap

type (
	// BrazilCNPJValidation
	BrazilCNPJValidationRequest struct {
		Cnpj        string `json:"cnpj"`
		CallbackUrl string `json:"callbackUrl"`
	}

	BrazilCNPJValidationResponse struct {
		Status int         `json:"status"`
		Error  interface{} `json:"error"`
		Data   struct {
			CompanyName         string      `json:"companyName"`
			CommercialName      string      `json:"commercialName"`
			Cpnj                string      `json:"cpnj"`
			DateOfOpening       string      `json:"dateOfOpening"`
			CompanySize         string      `json:"companySize"`
			Status              string      `json:"status"`
			StatusReason        interface{} `json:"statusReason"`
			StatusDate          string      `json:"statusDate"`
			SpecialStatus       interface{} `json:"specialStatus"`
			SpecialStatusReason interface{} `json:"specialStatusReason"`
			MainActivity        string      `json:"mainActivity"`
			SecondaryActivities []string    `json:"secondaryActivities"`
			CompanyType         string      `json:"companyType"`
			Address             struct {
				Street       string `json:"street"`
				Number       string `json:"number"`
				ZipCode      string `json:"zipCode"`
				District     string `json:"district"`
				Municipality string `json:"municipality"`
				State        string `json:"state"`
				Details      string `json:"details"`
			} `json:"address"`
			Contact struct {
				Email       string `json:"email"`
				PhoneNumber string `json:"phoneNumber"`
			} `json:"contact"`
			EFR interface{} `json:"EFR"`
		} `json:"data"`
		Timestamp string `json:"timestamp"`
	}

	// BrazilCNPJExtendedValidation
	BrazilCNPJExtendedValidationRequest struct {
		Cnpj        string `json:"cnpj"`
		CallbackUrl string `json:"callbackUrl"`
	}

	BrazilCNPJExtendedValidationResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			RegistrationNumber  string `json:"registrationNumber"`
			TipoEstabelecimento string `json:"tipoEstabelecimento"`
			CompanyName         string `json:"companyName"`
			LegalName           string `json:"legalName"`
			CadastralStatus     struct {
				RegistrationStatus string `json:"registrationStatus"`
				StatusDate         string `json:"statusDate"`
				Reason             string `json:"reason"`
			} `json:"cadastralStatus"`
			LegalNature struct {
				LegalCode        string `json:"legalCode"`
				LegalDescription string `json:"legalDescription"`
			} `json:"legalNature"`
			RegistrationDate string `json:"registrationDate"`
			CnaePrincipal    struct {
				CnaeCode        string `json:"cnaeCode"`
				CnaeDescription string `json:"cnaeDescription"`
			} `json:"cnaePrincipal"`
			Address struct {
				Street            string `json:"street"`
				Locality          string `json:"locality"`
				AreaNumber        string `json:"areaNumber"`
				AdditionalAddress string `json:"additionalAddress"`
				Postcode          string `json:"postcode"`
				LocalitySector    string `json:"localitySector"`
				FederationUnit    string `json:"federationUnit"`
				Country           struct {
					Code        string `json:"code"`
					Description string `json:"description"`
				} `json:"country"`
				County struct {
					AdministrativeCode string `json:"administrativeCode"`
					AdministrativeName string `json:"administrativeName"`
				} `json:"county"`
			} `json:"address"`
			MunicipalitiesJurisdiction struct {
				Code        string `json:"code"`
				Description string `json:"description"`
			} `json:"municipalitiesJurisdiction"`
			Phones []struct {
				PhoneCode   string `json:"phoneCode"`
				PhoneNumber string `json:"phoneNumber"`
			} `json:"phones"`
			Email                 string `json:"email"`
			CapitalSocial         string `json:"capitalSocial"`
			CompanySize           string `json:"companySize"`
			SituacaoEspecial      string `json:"situacaoEspecial"`
			DataSituacaoEspecial  string `json:"dataSituacaoEspecial"`
			InformacoesAdicionais struct {
				OptanteSimples       string        `json:"optanteSimples"`
				OptanteMei           string        `json:"optanteMei"`
				ListaPeriodosSimples []interface{} `json:"listaPeriodosSimples"`
			} `json:"informacoesAdicionais"`
			ShareholdersInfo []struct {
				ShareholderType string `json:"shareholderType"`
				CpfNumber       string `json:"cpfNumber"`
				ShareholderName string `json:"shareholderName"`
				Qualificacao    string `json:"qualificacao"`
				AssignationDate string `json:"assignationDate"`
				Country         struct {
					Code        string `json:"code"`
					Description string `json:"description"`
				} `json:"country"`
				LegalRepresentation struct {
					RepresentativeCpf  string `json:"representativeCpf"`
					RepresentativeName string `json:"representativeName"`
					Qualificacao       string `json:"qualificacao"`
				} `json:"legalRepresentation"`
			} `json:"shareholdersInfo"`
		} `json:"data"`
	}

	// BrazilCPFValidationResponse
	BrazilCPFValidationRequest struct {
		CpfNumber   string `json:"cpfNumber"`
		FullName    string `json:"fullName"`
		DateOfBirth string `json:"dateOfBirth"`
		MothersName string `json:"mothersName"`
		FathersName string `json:"fathersName"`
		// nationalId or driving-license
		DocumentType   string `json:"documentType"`
		DocumentNumber string `json:"documentNumber"`
		// M for male, F for female. Returns null if there is no gender data.
		Gender      string `json:"gender"`
		CallbackUrl string `json:"callbackUrl"`
	}

	BrazilCPFValidationResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			CpfNumberMatched      bool   `json:"cpfNumberMatched"`
			Gender                string `json:"gender"`
			Nationality           string `json:"nationality"`
			TaxStatus             string `json:"taxStatus"`
			FullNameSimilarity    int    `json:"fullNameSimilarity"`
			DateOfBirthMatched    bool   `json:"dateOfBirthMatched"`
			DocumentType          string `json:"documentType"`
			DocumentNumberMatched bool   `json:"documentNumberMatched"`
			MothersNameSimilarity int    `json:"mothersNameSimilarity"`
			FathersNameSimilarity int    `json:"fathersNameSimilarity"`
		} `json:"data"`
	}

	// BrazilCPFLEGACY
	BrazilCPFLEGACYRequest struct {
		CpfNumber      string `json:"cpfNumber"`
		FullName       string `json:"fullName"`
		DocumentNumber string `json:"documentNumber"`
		DateOfBirth    string `json:"dateOfBirth"`
		DateOfExpiry   string `json:"dateOfExpiry"`
		DocumentType   string `json:"documentType"`
	}

	BrazilCPFLEGACYResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			CPFNumber      string `json:"CPFNumber"`
			Gender         string `json:"gender"`
			Nationality    string `json:"nationality"`
			TaxStatus      string `json:"taxStatus"`
			FullName       string `json:"fullName"`
			DateOfBirth    string `json:"dateOfBirth"`
			DocumentType   string `json:"documentType"`
			DocumentNumber string `json:"documentNumber"`
		} `json:"data"`
	}
)

/*
BrazilCNPJValidation validate a business's National Registry of Legal Entities number.

This method takes in the BrazilCNPJValidationRequest{} struct as a parameter.

MetaMap connects with the Brazilian Internal Revenue Service (Ministério da Fazenda / Treasury) to validate the National Registry of Legal Entities (Cadastro Nacional de Pessoas Jurídicas / CNPJ) number.
*/
func (c *Client) BrazilCNPJValidation(req BrazilCNPJValidationRequest) (*BrazilCNPJValidationResponse, error) {
	url := "govchecks/v1/br/cnpj"
	method := MethodPOST
	var response BrazilCNPJValidationResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &BrazilCNPJValidationResponse{}, err
	}

	return &response, nil
}

/*
BrazilCNPJExtendedValidation validate a business's National Registry of Legal Entities number.

This method takes in the BrazilCNPJExtendedValidationRequest{} struct as a parameter.

MetaMap connects with the Brazilian Internal Revenue Service (Ministério da Fazenda / Treasury) to validate the National Registry of Legal Entities (Cadastro Nacional de Pessoas Jurídicas / CNPJ) number.

If the CNPJ number is valid, the endpoint will also pull up additional information, including company registration data, type of business, phone, email, physical address, and information about shareholders, including their names and CPF numbers.
*/
func (c *Client) BrazilCNPJExtendedValidation(req BrazilCNPJExtendedValidationRequest) (*BrazilCNPJExtendedValidationResponse, error) {
	url := "govchecks/v1/br/cnpj-extended"
	method := MethodPOST
	var response BrazilCNPJExtendedValidationResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &BrazilCNPJExtendedValidationResponse{}, err
	}

	return &response, nil
}

/*
BrazilCPFValidation verify a user's CPF number and identity.

This method takes in the BrazilCPFValidationRequest{} struct as a parameter.

MetaMap connects with the Brazilian Internal Revenue Service (Ministério da Fazenda / Treasury) to validate that the Registration of Individuals (Cadastro de Pessoas Físicas / CPF) number present in the ID card exists and its owner matches the data obtained from it.
*/
func (c *Client) BrazilCPFValidation(req BrazilCPFValidationRequest) (*BrazilCPFValidationResponse, error) {
	url := "govchecks/v1/br/cpf-validation"
	method := MethodPOST
	var response BrazilCPFValidationResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &BrazilCPFValidationResponse{}, err
	}

	return &response, nil
}

/*
BrazilCPFLEGACY verify a user's CPF number and identity.

This method takes in the BrazilCPFLEGACYRequest{} struct as a parameter.

NOTE: This version of the CPF check only handles individual validation requests. Use the new version of the BrazilCPFValidation() which can handle batch validation requests.

MetaMap connects with the Brazilian IRS (Ministério da Fazenda / Treasury) to validate that the CPF (Cadastro de Pessoas Físicas / Registration of Individuals) number present in the ID card exists and its owner matches the data obtained from it.
*/
func (c *Client) BrazilCPFLEGACY(req BrazilCPFLEGACYRequest) (*BrazilCPFLEGACYResponse, error) {
	url := "govchecks/v1/br/cpf"
	method := MethodPOST
	var response BrazilCPFLEGACYResponse
	c.IsBasic = false
	//this method required multipart form data
	c.IsMultipartHeader = true
	if err := c.newRequest(method, url, req, response); err != nil {
		return &BrazilCPFLEGACYResponse{}, err
	}

	return &response, nil
}
