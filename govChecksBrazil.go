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
