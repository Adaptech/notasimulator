package main

type createOrganizationRequest struct {
	OrganizationID  string `json:"organizationId"`
	Name            string `json:"name"`
	ElectionAdminID string `json:"electionAdminId"`
}

func createOrganization(id string, name string) (organizationID string, err error) {
	err = jsonReq(
		"/api/v1/organization/create",
		createOrganizationRequest{
			OrganizationID:  id,
			Name:            name,
			ElectionAdminID: "admin-1",
		},
	)
	return id, err
}
