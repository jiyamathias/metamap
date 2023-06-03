# metamap
A Golang SDK for interacting with the metamap API

# Please ensure to create issues in this repo if :
- You encounter any error while using this package and that issue would be attended to immediately.

# Get Started
- In other to use this package, you need to first create an account with MEtaMap via https://metamap.com to get your `Client_Id` and `Client_Secret`

# Installation
To install this package, you need to install [Go](https://golang.org/) and set your Go workspace first.
1. You can use the below Go command to install metamap
```sh
$ go get -u github.com/iqquee/metamap
```
2. Import it in your code:
```sh
import "github.com/iqquee/metamap"
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

# GovChecks: Argentina

- ## ArgentinaDNI
ArgentinaDNI verify a user's DNI (Documento Nacional de Identidad) based on card number and issue date.

This method takes in the ArgentinaDNIRequest{} struct as a parameter.
### Use this object payload to implement the ArgentinaDNI() method
```go
type ArgentinaDNIRequest struct {
	// Document number from either a National ID or a driver license.
	DocumentNumber string `json:"documentNumber"`
	DateOfBirth    string `json:"dateOfBirth"`
	// M for male, F for female.
	Gender string `json:"gender"`
	// Document issue date.
	DateOfIssue string                 `json:"dateOfIssue"`
	CallbackUrl string                 `json:"callbackUrl"`
	// Use Metadata to add internal references for your outputs/webhooks.
	Metadata    map[string]interface{} `json:"metadata"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ArgentinaDNIRequest{
		DocumentNumber: "your document number",
		DateOfBirth:      "your data of birth e.g 2023-05-10",
		Gender:       "M", // for male
		DateOfIssue: "your data of issue",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ArgentinaDNI(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## ArgentinaRENAPER
ArgentinaRENAPER verify a user's DNI (Documento Nacional de Identidad) number and identity.

NOTE: Use the ArgentinaDNI() to validate the submitted DNI card.

This method takes in the ArgentinaRENAPERRequest{} struct as a parameter.
### Use this object payload to implement the ArgentinaRENAPER() method
```go
type ArgentinaRENAPERRequest struct {
	// Document number from either a National ID or a driver license.
	DocumentNumber string `json:"documentNumber"`
	DateOfBirth    string `json:"dateOfBirth"`
	// M for male, F for female.
	Gender string `json:"gender"`
	// Document issue date.
	DateOfIssue string                 `json:"dateOfIssue"`
	CallbackUrl string                 `json:"callbackUrl"`
	// Use Metadata to add internal references for your outputs/webhooks.
	Metadata    map[string]interface{} `json:"metadata"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ArgentinaRENAPERRequest{
		DocumentNumber: "your document number",
		DateOfBirth:      "your data of birth e.g 2023-05-10",
		Gender:       "M", // for male
		DateOfIssue: "your data of issue",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ArgentinaRENAPER(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## ArgentinaRENAPERPremium
ArgentinaRENAPERPremium verify a user's DNI (Documento Nacional de Identidad) number and identity.

NOTE: You do not need to validate the submitted DNI card parameters via ArgentinaDNI() as this check is implemented in here.

This method takes in the ArgentinaRENAPERPremiumRequest{} struct as a parameter.
### Use this object payload to implement the ArgentinaRENAPERPremium() method
```go
type ArgentinaRENAPERPremiumRequest struct {
	// Document number from either a National ID or a driver license.
	DocumentNumber string `json:"documentNumber"`
	// E.g "2023-05-10"
	DateOfBirth string `json:"dateOfBirth"`
	// Date of document issue E.g "2023-05-10"
	IssueDate string `json:"issueDate"`
	// M for male, F for female.
	Gender      string `json:"gender"`
	CallbackUrl string `json:"callbackUrl"`
	// Use Metadata to add internal references for your outputs/webhooks.
	Metadata map[string]interface{} `json:"metadata"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ArgentinaRENAPERPremiumRequest{
		DocumentNumber: "your document number",
		DateOfBirth:      "your data of birth e.g 2023-05-10",
		Gender:       "M", // for male
		IssueDate: "your data of issue",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ArgentinaRENAPERPremium(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## ArgentinaRENAPERExtended
ArgentinaRENAPERExtended verify a user's DNI (Documento Nacional de Identidad) number and identity against multiple databases.

NOTE: Use the ArgentinaDNI() to validate the submitted DNI card.

This method takes in the ArgentinaRENAPERExtendedRequest{} struct as a parameter.
### Use this object payload to implement the ArgentinaRENAPERExtended() method
```go
type ArgentinaRENAPERExtendedRequest struct {
	// Document number from either a National ID or a driver license
	DocumentNumber string `json:"documentNumber"`
	FullName       string `json:"fullName"`
	CallbackUrl    string `json:"callbackUrl"`
	// Use Metadata to add internal references for your outputs/webhooks.
	Metadata map[string]interface{} `json:"metadata"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ArgentinaRENAPERExtendedRequest{
		DocumentNumber: "your document number",
		FullName:      "your full name",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ArgentinaRENAPERExtended(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

# GovChecks: Brazil

- ## BrazilCNPJValidation
BrazilCNPJValidation validate a business's National Registry of Legal Entities number.

This method takes in the BrazilCNPJValidationRequest{} struct as a parameter.
### Use this object payload to implement the BrazilCNPJValidation() method
```go
type BrazilCNPJValidationRequest struct {
	Cnpj        string `json:"cnpj"`
	CallbackUrl string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.BrazilCNPJValidationRequest{
		Cnpj: "your cnpj number",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.BrazilCNPJValidation(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## BrazilCNPJExtendedValidation
BrazilCNPJExtendedValidation validate a business's National Registry of Legal Entities number.

This method takes in the BrazilCNPJExtendedValidationRequest{} struct as a parameter.
### Use this object payload to implement the BrazilCNPJExtendedValidation() method
```go
type BrazilCNPJExtendedValidationRequest struct {
	Cnpj        string `json:"cnpj"`
	CallbackUrl string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.BrazilCNPJExtendedValidationRequest{
		Cnpj: "your cnpj number",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.BrazilCNPJExtendedValidation(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## BrazilCPFValidation
BrazilCPFValidation verify a user's CPF number and identity.

This method takes in the BrazilCPFValidationRequest{} struct as a parameter.
### Use this object payload to implement the BrazilCPFValidation() method
```go
type BrazilCPFValidationRequest struct {
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

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.BrazilCPFValidationRequest{
		CpfNumber: "012.345.678-90",
		FullName: "JOHN DOE",
		DateOfBirth: "1900-01-01",
		MothersName: "JANE DOE",
		FathersName: "JAMES DOE",
		DocumentType: "national-id",
		DocumentNumber: "0123456789",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.BrazilCPFValidation(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```


- ## BrazilCPFLEGACY
BrazilCPFLEGACY verify a user's CPF number and identity.

NOTE: This version of the CPF check only handles individual validation requests. Use the new version of the BrazilCPFValidation() which can handle batch validation requests.

This method takes in the BrazilCPFLEGACYRequest{} struct as a parameter.
### Use this object payload to implement the BrazilCPFValidation() method
```go
type BrazilCPFLEGACYRequest struct {
	CpfNumber      string `json:"cpfNumber"`
	FullName       string `json:"fullName"`
	DocumentNumber string `json:"documentNumber"`
	DateOfBirth    string `json:"dateOfBirth"`
	DateOfExpiry   string `json:"dateOfExpiry"`
	DocumentType   string `json:"documentType"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.BrazilCPFLEGACYRequest{
		CpfNumber: "012.345.678-90",
		FullName: "JOHN DOE",
		DateOfBirth: "1900-01-01",
		DateOfExpiry: "2021-12-31",
		DocumentType: "national-id",
		DocumentNumber: "0123456789",
	}
	resp, err := client.BrazilCPFLEGACY(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## BrazilCPFLight
BrazilCPFLight verify a user's CPF number and identity.

This method takes in the BrazilCPFLightRequest{} struct as a parameter.
### Use this object payload to implement the BrazilCPFLight() method
```go
type BrazilCPFLightRequest struct {
	Cpf         string `json:"cpf"`
	CallbackUrl string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.BrazilCPFLightRequest{
		Cpf: "01234567",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.BrazilCPFLight(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

# GovChecks: Chile

- ## ChileRegistroCivil
ChileRegistroCivil verify that a user's RUN number against the Chilean Civil Registry database.

This method takes in the ChileRegistroCivilRequest{} struct as a parameter.
### Use this object payload to implement the ChileRegistroCivil() method
```go
type ChileRegistroCivilRequest struct {
	RunNumber    string `json:"runNumber"`
	DocumentType string `json:"documentType"`
	Nationality    string `json:"nationality"`
	DocumentNumber string `json:"documentNumber"`
	CallbackUrl    string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ChileRegistroCivilRequest{
		RunNumber: "run number",
		DocumentType : "document type",
		Nationality: "CHL",
		DocumentNumber: "document number",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ChileRegistroCivil(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## ChileCriminalCertificate
ChileCriminalCertificate verify that a user's background certificate is valid.

This method takes in the ChileCriminalCertificateRequest{} struct as a parameter.
### Use this object payload to implement the ChileCriminalCertificate() method
```go
type ChileCriminalCertificateRequest struct {
	SheetNumber      string `json:"sheetNumber"`
	VerificationCode string `json:"verificationCode"`
	CallbackUrl      string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ChileCriminalCertificateRequest{
		SheetNumber: "012345678901",
		VerificationCode: "1ab23cd45e6f",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ChileCriminalCertificate(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

# GovChecks: Colombia

- ## ColombiaMigrationInstitute
ColombiaMigrationInstitute verify that a user has a valid immigration status in Colombia.

This method takes in the ColombiaMigrationInstituteRequest{} struct as a parameter.

NOTE:  DateOfIssue fieid should be in this format 2022-01-01.
### Use this object payload to implement the ColombiaMigrationInstitute() method
```go
type ColombiaMigrationInstituteRequest struct {
	DocumentNumber string `json:"documentNumber"`
	DateOfIssue    string `json:"dateOfIssue"`
	CallbackUrl    string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ColombiaMigrationInstituteRequest{
		DocumentNumber: "document number",
		DateOfIssue: "2022-01-01",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ColombiaMigrationInstitute(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## ColombiaRegistraduria
ColombiaRegistraduria verify a user's ID card against the Colombian Civil Registry.

This method takes in the ColombiaRegistraduriaRequest{} struct as a parameter.
### Use this object payload to implement the ColombiaRegistraduria() method
```go
type ColombiaRegistraduriaRequest struct {
	DocumentNumber string `json:"documentNumber"`
	CallbackUrl    string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ColombiaRegistraduriaRequest{
		DocumentNumber: "document number",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ColombiaRegistraduria(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```


- ## ColombiaRegistraduriaLight
ColombiaMigrationInstitute verify that a user has a valid immigration status in Colombia.

This method takes in the ColombiaRegistraduriaLightRequest{} struct as a parameter.

NOTE:  DateOfIssue fieid should be in this format 2022-01-01.
### Use this object payload to implement the ColombiaRegistraduriaLight() method
```go
type ColombiaRegistraduriaLightRequest struct {
	DocumentNumber string `json:"documentNumber"`
	DateOfIssue    string `json:"dateOfIssue"`
	CallbackUrl    string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ColombiaRegistraduriaLightRequest{
		DocumentNumber: "document number",
		DateOfIssue: "2022-01-01",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ColombiaRegistraduriaLight(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## ColombiaUnifiedLegalSearch
ColombiaUnifiedLegalSearch check a user's full name against Colombian police records.

This method takes in the ColombiaUnifiedLegalSearchRequest{} struct as a parameter.
### Use this object payload to implement the ColombiaUnifiedLegalSearch() method
```go
type ColombiaUnifiedLegalSearchRequest struct {
	FullName    string `json:"fullName"`
	CallbackUrl string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ColombiaUnifiedLegalSearchRequest{
		FullName: "JOHN DOE",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ColombiaUnifiedLegalSearch(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## ColombiaRUES
ColombiaRUES check a business's national tax ID against Colombian Single Business and Social Registry.

This method takes in the ColombiaRUESRequest{} struct as a parameter.
### Use this object payload to implement the ColombiaRUES() method
```go
type ColombiaRUESRequest struct {
	Nit         string `json:"nit"`
	CallbackUrl string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.ColombiaRUESRequest{
		Nit: "000000000",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.ColombiaRUES(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

# GovChecks: Nigeria
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
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
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
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
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
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.NigeriaVotingIDNumberRequest{
		DocumentNumber: "your voters card number",
		FirstName:      "John",
		LastName:       "Doe",
		DateOfBirth:    "2023-05-01",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
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
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.NigeriaDriversLicenceRequest{
		DocumentNumber: "your drivers licence",
		FirstName:      "John",
		LastName:       "Doe",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.NigeriaDriversLicence(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```


- ## NigeriaCAC
NigeriaCAC verifies a company's CAC (Corporate Affairs Commission) number.
### Use this object payload to implement the NigeriaCorporateAffairsCommission() method
```go
type NigeriaCACRequest struct {
	RegistrationNumber string `json:"registrationNumber"`
	CallbackUrl        string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.NigeriaCACRequest{
		RegistrationNumber: "your registration number",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.NigeriaCAC(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## NigeriaTaxIdentificationNumber
NigeriaTaxIdentificationNumber verifies a company's Tax Identification Number.
### Use this object payload to implement the NigeriaCorporateAffairsCommission() method
```go
type NigeriaTaxIdentificationNumberRequest struct {
	TaxNumber   string `json:"taxNumber"`
	CallbackUrl string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.NigeriaTaxIdentificationNumberRequest{
		TaxNumber: "your tax number",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.NigeriaTaxIdentificationNumber(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## NigeriaCACAffiliates
NigeriaCACAffiliates verifies a major shareholder names and titles if a company's Corporate Affairs Commission number is valid.
### Use this object payload to implement the NigeriaCACAffiliates() method
```go
type NigeriaCACAffiliatesRequest struct {
	RegistrationNumber string `json:"registrationNumber"`
	CallbackUrl        string `json:"callbackUrl"`
}

```
```go
package main

import (
	"fmt"
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.NigeriaCACAffiliatesRequest{
		RegistrationNumber: "your registration number",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.NigeriaCACAffiliates(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```

- ## NigeriaBVN
verifies a BVN (Bank Verification Number).
### Use this object payload to implement the NigeriaBVN() method
```go
type NigeriaBVNRequest struct {
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
	metamap "github.com/iqquee/metamap"
)

func main() {
	clientId := "your metamap client id"
	clientSecret := "your metamap client secrete"
	httpClient := hhttp.DefaultClient
	client := metamap.New(&httpClient, clientId, clientSecret)

	req := metamap.NigeriaBVNRequest{
		DocumentNumber: "your BVN number",
		FirstName: "John",
		LastName: "Doe",
		CallbackUrl: "https://webhook.site/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
	}
	resp, err := client.NigeriaBVN(req)
	if err != nil {
		fmt.Println("This is the error: ", err)
		return
	}

	fmt.Println("This is the response: ", resp)
}
```