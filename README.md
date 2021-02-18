# NameAuditor GO SDK
A Golang SDK for [NameAuditor API](https://docs.nameauditor.com), A fast & powerful API that returns a well-parsed & accurate WHOIS information of a domain name in JSON formats.

### Install
```
go get github.com/softpi/nameauditor-go-sdk
```

### API key

* Log in or sign up for your [RapidAPI](https://rapidapi.com/marketplace) account.
* Navigate to any API documentation page by searching for or clicking on one of the collections from the homepage.
* Scroll down to the “Header Parameters” section of the API console.
* Your API Key should be visible in the “X-RapidAPI-Key” field.

### Example
```go
package main

import (
	"fmt"

	nauditor_sdk "github.com/softpi/nameauditor-go-sdk"
)

func main() {
    // New Client
	client := nauditor_sdk.NewClient("Replace_API_Key")

	// GET /ping
	if err := client.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	// GET /tlds
	tlds, err := client.GetTLDs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tlds)

	// GET whois/{domain}
	record, err := client.GetWhois("google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(record)
}
```