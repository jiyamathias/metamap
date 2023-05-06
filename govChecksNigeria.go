package metamap

// {
// 	"error": null,
// 	"data": {
// 	  "id": "1",
// 	  "firstName": "John",
// 	  "lastName": "Doe",
// 	  "dateOfBirth": "04-04-1944",
// 	  "gender": "M",
// 	  "phone": "12341234567890",
// 	  "vNin": "JZ426633988976CH",
// 	  "photo": "https://media.getmati.com/file?location=eyUkp7zR7LSTnV4x4lJma7L0"
// 	}
//   }

type (
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
)

/*
VirtualNIN verifies NIN(National Identification Number).

This method takes in the NigeriaNINRequest{} struct as a parameter. Both the VNIN  and CallbackUrl field are required parameter that must be passed in for the request.

It also takes in an optional parameter `metadata` which is a map[string]interface that is passed in to add custom data to the request.
*/
func (c *Client) VirtualNIN(req NigeriaNINRequest) (*NigeriaNINResponse, error) {
	url := c.BaseUrl + "govchecks/v1/ng/vnin"
	method := MethodPOST
	var response NigeriaNINResponse
	if err := c.newRequest(method, url, req, response); err != nil {
		return &NigeriaNINResponse{}, err
	}

	return &response, nil
}
