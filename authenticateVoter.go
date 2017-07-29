package main

type AuthenticateVoterRequest struct {
	ReferendumId   string `json:"referendumId"`
	VoterId        string `json:"voterId"`
	OrganizationId string `json:"organizationId"`
}

func authenticateVoter(referendumID string, voterID string, organizationID string) error {
	err := jsonReq(
		"http://localhost:3001/api/v1/organization/referendum/voter/authenticate",
		AuthenticateVoterRequest{
			ReferendumId:   referendumID,
			VoterId:        voterID,
			OrganizationId: organizationID,
		},
	)
	return err
}
