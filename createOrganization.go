package main

type CreateOrganizationRequest struct {
	OrganizationId  string `json:"organizationId"`
	Name            string `json:"name"`
	ElectionAdminId string `json:"electionAdminId"`
}

func createOrganization(id string, name string) (organizationId string, err error) {
	err = jsonReq(
		"http://localhost:3001/api/v1/organization/create",
		CreateOrganizationRequest{
			OrganizationId:  id,
			Name:            name,
			ElectionAdminId: "admin-1",
		},
	)
	return id, err
}
