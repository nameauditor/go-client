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
	record, err := client.GetWhois("nameauditor.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(record)
}

// Output
{
        "UUID": "c0mesvium8seq4vkga90",
        "ID": "2529097792_DOMAIN_COM-VRSN",
        "Full": "nameauditor.com",
        "Name": "nameauditor",
        "Extension": "com",
        "Availability": "registered",
        "Length": 11,
        "Age": "38 weeks 4 days 19 hours",
        "TimeToExpire": "13 weeks 3 days 5 hours",
        "TimeSinceLastUpdate": "38 weeks 4 days 16 hours",
        "ContainNumber": false,
        "ContainHyphen": false,
        "CreationDate": "2020-05-22T14:58:45Z",
        "UpdatedDate": "2020-05-22T18:18:32Z",
        "ExpirationDate": "2021-05-22T14:58:45Z",
        "Status": [
            "clientTransferProhibited"
        ],
        "Nameservers": {
            "CORTNEY.NS.CLOUDFLARE.COM": [
                "108.162.192.87",
                "173.245.58.87",
                "172.64.32.87",
                "2606:4700:50::adf5:3a57",
                "2803:f800:50::6ca2:c057",
                "2a06:98c1:50::ac40:2057"
            ],
            "KARL.NS.CLOUDFLARE.COM": [
                "173.245.59.190",
                "108.162.193.190",
                "172.64.33.190",
                "2606:4700:58::adf5:3bbe",
                "2a06:98c1:50::ac40:21be",
                "2803:f800:50::6ca2:c1be"
            ]
        },
        "DNSSEC": "unsigned",
        "RegistrarID": "1068",
        "RegistrarName": "NameCheap, Inc.",
        "RegistrarAbuseContactEmail": "abuse@namecheap.com",
        "WhoisServer": "whois.namecheap.com",
        "Contact": {
            "Registrant": {
                "ID": "",
                "Name": "WhoisGuard Protected",
                "Organization": "WhoisGuard, Inc.",
                "Street": "P.O. Box 0823-03411 ",
                "City": "Panama",
                "Province": "Panama",
                "PostalCode": "",
                "Country": "PA",
                "Phone": "+507.8365503",
                "Fax": "+51.17057182",
                "Email": "19b83006b557452eb54fc2d0cff8c6f7.protect@whoisguard.com"
            },
            "Admin": {
                "ID": "",
                "Name": "WhoisGuard Protected",
                "Organization": "WhoisGuard, Inc.",
                "Street": "P.O. Box 0823-03411 ",
                "City": "Panama",
                "Province": "Panama",
                "PostalCode": "",
                "Country": "PA",
                "Phone": "+507.8365503",
                "Fax": "+51.17057182",
                "Email": "19b83006b557452eb54fc2d0cff8c6f7.protect@whoisguard.com"
            },
            "Tech": {
                "ID": "",
                "Name": "WhoisGuard Protected",
                "Organization": "WhoisGuard, Inc.",
                "Street": "P.O. Box 0823-03411 ",
                "City": "Panama",
                "Province": "Panama",
                "PostalCode": "",
                "Country": "PA",
                "Phone": "+507.8365503",
                "Fax": "+51.17057182",
                "Email": "19b83006b557452eb54fc2d0cff8c6f7.protect@whoisguard.com"
            }
        },
        "FirstWhoisRaw": "Domain Name: NAMEAUDITOR.COM...",
        "LastWhoisRaw": "Domain name: nameauditor.com...",
        "CreatedAt": "2021-02-17T10:22:54.474154566Z"
    }
```