package main

type CreateElectionAdminRequest struct {
	ElectionAdminId string `json:"electionAdminId"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	Address         `json:"address"`
}

func createElectionAdmin(id string) (adminId string, err error) {
	err = jsonReq(
		"http://localhost:3001/api/v1/electionadmin/create",
		CreateElectionAdminRequest{
			ElectionAdminId: id,
			Firstname:       "Joe",
			Lastname:        "Admin",
			Address: Address{
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
