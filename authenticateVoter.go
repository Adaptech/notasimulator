package main

type authenticateVoterRequest struct {
	ReferendumID   string `json:"referendumId"`
	VoterID        string `json:"voterId"`
	OrganizationID string `json:"organizationId"`
}

func authenticateVoter(referendumID string, voterID string, organizationID string) error {
	err := jsonReq(
		"/api/v1/organization/referendum/voter/authenticate",
		authenticateVoterRequest{
			ReferendumID:   referendumID,
			VoterID:        voterID,
			OrganizationID: organizationID,
		},
	)
	return err
}
