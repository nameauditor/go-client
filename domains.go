package nauditor_sdk

import "time"

// TLDsList
type TLDsList struct {
	Total int
	TLDs  []string
}

// Domain
type Domain struct {
	UUID string
	//
	ID                  string
	Full                string
	Name                string
	Extension           string
	Availability        string
	Length              int
	Age                 string
	TimeToExpire        string
	TimeSinceLastUpdate string
	ContainNumber       bool
	ContainHyphen       bool
	//
	CreationDate   time.Time
	UpdatedDate    time.Time
	ExpirationDate time.Time
	//
	Status      []string
	Nameservers map[string][]string
	//
	DNSSEC string

	// registrar stuff
	RegistrarID                string
	RegistrarName              string
	RegistrarAbuseContactEmail string
	WhoisServer                string
	//
	Contact struct {
		Registrant struct {
			ID           string
			Name         string
			Organization string
			Street       string
			City         string
			Province     string
			PostalCode   string
			Country      string
			Phone        string
			Fax          string
			Email        string
		}
		Admin struct {
			ID           string
			Name         string
			Organization string
			Street       string
			City         string
			Province     string
			PostalCode   string
			Country      string
			Phone        string
			Fax          string
			Email        string
		}
		Tech struct {
			ID           string
			Name         string
			Organization string
			Street       string
			City         string
			Province     string
			PostalCode   string
			Country      string
			Phone        string
			Fax          string
			Email        string
		}
	}
	//
	FirstWhoisRaw string
	LastWhoisRaw  string
	//
	CreatedAt time.Time
}
