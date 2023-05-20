package metamap

type (

	// ChileRegistroCivil
	ChileRegistroCivilRequest struct {
		RunNumber    string `json:"runNumber"`
		DocumentType string `json:"documentType"`
		// 3-letter country code as per ISO 3166-1 alpha-3. You can visit https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3 country code reference
		Nationality    string `json:"nationality"`
		DocumentNumber string `json:"documentNumber"`
		CallbackUrl    string `json:"callbackUrl"`
	}

	ChileRegistroCivilResponse struct {
		Error interface{} `json:"error"`
		Data  struct {
			DocumentType string `json:"documentType"`
		} `json:"data"`
	}
)

/*
ChileRegistroCivil verify that a user's RUN number against the Chilean Civil Registry database.

This method takes in the ChileRegistroCivilRequest{} struct as a parameter.

NOTE: 3-letter country code as per ISO 3166-1 alpha-3.You can visit https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3 for the country code reference.

MetaMap connects with the Chilean Civil Registry (Servicio de Registro Civil e Identificación / SRCEI) to validate that the Rol Único Nacional (RUN) number present in the ID card for foreigners exists and its currently valid.
*/
func (c *Client) ChileRegistroCivil(req ChileRegistroCivilRequest) (*ChileRegistroCivilResponse, error) {
	url := "govchecks/v1/cl/registro-civil"
	method := MethodPOST
	var response ChileRegistroCivilResponse
	c.IsBasic = false
	if err := c.newRequest(method, url, req, response); err != nil {
		return &ChileRegistroCivilResponse{}, err
	}

	return &response, nil
}
