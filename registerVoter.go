package main

type registerVoterRequest struct {
	VoterID        string  `json:"voterId"`
	OrganizationID string  `json:"organizationId"`
	Firstname      string  `json:"firstname"`
	Lastname       string  `json:"lastname"`
	Address        address `json:"address"`
}

func registerVoter(id string, organizationID string, firstname string, lastname string, streetAddress string, postOfficeBoxNumber string, addressLocality string, addressRegion string, postalCode string, addressCountry string) error {
	err := jsonReq(
		"/api/v1/organization/voter/register",
		registerVoterRequest{
			VoterID:        id,
			OrganizationID: organizationID,
			Firstname:      firstname,
			Lastname:       lastname,
			Address: address{
				StreetAddress:       streetAddress,
				PostOfficeBoxNumber: postOfficeBoxNumber,
				AddressLocality:     addressLocality,
				AddressRegion:       addressRegion,
				PostalCode:          postalCode,
				AddressCountry:      addressCountry,
			},
		},
	)
	return err
}
