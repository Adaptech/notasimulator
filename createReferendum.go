package main

type CreateReferendumRequest struct {
	ReferendumId   string   `json:"referendumId"`
	OrganizationId string   `json:"organizationId"`
	Name           string   `json:"name"`
	Proposal       string   `json:"proposal"`
	Options        []string `json:"options"`
}

func createReferendum(id string, organizationID string, name string, proposal string, options []string) (referendumId string, err error) {
	err = jsonReq(
		"http://localhost:3001/api/v1/organization/referendum/create",
		CreateReferendumRequest{
			ReferendumId:   id,
			OrganizationId: organizationID,
			Name:           name,
			Proposal:       proposal,
			Options:        options,
		},
	)
	return id, err
}
