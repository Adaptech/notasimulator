package main

type createReferendumRequest struct {
	ReferendumID   string   `json:"referendumId"`
	OrganizationID string   `json:"organizationId"`
	Name           string   `json:"name"`
	Proposal       string   `json:"proposal"`
	Options        []string `json:"options"`
}

func createReferendum(id string, organizationID string, name string, proposal string, options []string) (referendumID string, err error) {
	err = jsonReq(
		"/api/v1/organization/referendum/create",
		createReferendumRequest{
			ReferendumID:   id,
			OrganizationID: organizationID,
			Name:           name,
			Proposal:       proposal,
			Options:        options,
		},
	)
	return id, err
}
