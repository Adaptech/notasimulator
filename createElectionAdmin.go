package main

type createElectionAdminRequest struct {
	ElectionAdminID string  `json:"electionAdminId"`
	Firstname       string  `json:"firstname"`
	Lastname        string  `json:"lastname"`
	Address         address `json:"address"`
}

func createElectionAdmin(id string) (adminID string, err error) {
	err = jsonReq(
		"/api/v1/electionadmin/create",
		createElectionAdminRequest{
			ElectionAdminID: id,
			Firstname:       "Joe",
			Lastname:        "Admin",
			Address: address{
				StreetAddress:       "405 E. Stueben",
				PostOfficeBoxNumber: "",
				AddressLocality:     "Bingen",
				AddressRegion:       "WA",
				PostalCode:          "98605",
				AddressCountry:      "US",
			},
		},
	)
	return id, err
}
