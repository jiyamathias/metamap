# metamap-go
A Golang SDK for interacting with the metamap API

# Please ensure to create issues in this repo if :
- You encounter any error while using this package and that issue would be attended to immediately.

# Get Started
- In other to use this package, you need to first create an account with MEtaMap via https://metamap.com to get your `Client_Id` and `Client_Secret`

# Installation
To install this package, you need to install [Go](https://golang.org/) and set your Go workspace first.
1. You can use the below Go command to install metamap-go
```sh
$ go get -u github.com/iqquee/metamap-go
```
2. Import it in your code:
```sh
import "github.com/iqquee/metamap-go"
```
## Note : All methods in this package returns two (2) things:
- [x] An object of the response
- [x] An error (if any)

# Quick start
```sh
# assume the following codes in example.go file
$ touch example.go
# open the just created example.go file in the text editor of your choice
```


## GovChecks: Nigeria
- ## NigeriaVirtualNIN
NigeriaVirtualNIN verifies NIN(National Identification Number).

NOTE: The MetaData field in the metamap.NigeriaNINRequest{} struct is an optional parameter which can contain extra data to add internal references.
### Use this object payload to implement the NigeriaVirtualNIN() method
```go
type NigeriaNINRequest struct {
		VNIN        string `json:"vNIN"`
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		DateOfBirth string `json:"dateOfBirth"`
		Metadata map[string]interface{} `json:"metadata"`
		//CallbackUrl is a required parameter that must be passed in
		CallbackUrl string `json:"callbackUrl"`
}
```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap-go"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := http.Client{}
	client := metamap.New(&httpClient, clientId, clientSecret)

	extraData := make(map[string]interface{})
	extraData["user_id"] = "08992e987"
	extraData["some extra data"] = "some extra data"
	req := metamap.NigeriaNINRequest{
		VNIN:        "your national identity number",
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: "2023-05-01",
		Metadata:    extraData,
		CallbackUrl: "your callback url",
	}
	resp, err := client.NigeriaVirtualNIN(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## NigeriaVotingIDNumber
NigeriaVotingIDNumber verifies NIN(National Identification Number).
### Use this object payload to implement the NigeriaVirtualNIN() method
```go
type NigeriaVotingIDNumberRequest struct {
		DocumentNumber string `json:"documentNumber"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		DateOfBirth    string `json:"dateOfBirth"`
		CallbackUrl    string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap-go"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := http.Client{}
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.NigeriaVotingIDNumberRequest{
		DocumentNumber: "your voters card number",
		FirstName:      "John",
		LastName:       "Doe",
		DateOfBirth:    "2023-05-01",
		CallbackUrl:    "your callback url",
	}
	resp, err := client.NigeriaVotingIDNumber(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## NigeriaDriversLicence
NigeriaDriversLicence verifies drivers licence.
### Use this object payload to implement the NigeriaDriversLicence() method
```go
type NigeriaDriversLicenceRequest struct {
		DocumentNumber string `json:"documentNumber"`
		FirstName      string `json:"firstName"`
		LastName       string `json:"lastName"`
		CallbackUrl    string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap-go"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := http.Client{}
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.NigeriaDriversLicenceRequest{
		DocumentNumber: "your drivers licence",
		FirstName:      "John",
		LastName:       "Doe",
		CallbackUrl:    "your callback url",
	}
	resp, err := client.NigeriaDriversLicence(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```