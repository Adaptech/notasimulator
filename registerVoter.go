package main

type RegisterVoterRequest struct {
	VoterId        string `json:"voterId"`
	OrganizationId string `json:"organizationId"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Address        `json:"address"`
}

func registerVoter(id string, organizationID string, firstname string, lastname string, streetAddress string, postOfficeBoxNumber string, addressLocality string, addressRegion string, postalCode string, addressCountry string) error {
	err := jsonReq(
		"http://localhost:3001/api/v1/organization/voter/register",
		RegisterVoterRequest{
			VoterId:        id,
			OrganizationId: organizationID,
			Firstname:      firstname,
			Lastname:       lastname,
			Address: Address{
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
