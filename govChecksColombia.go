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
