package main

type castVoteRequest struct {
	ReferendumID string `json:"referendumId"`
	VoterID      string `json:"voterId"`
	Vote         string `json:"vote"`
}

func castVote(referendumID string, voterID string, vote string) error {
	err := jsonReq(
		"/api/v1/organization/referendum/vote",
		castVoteRequest{
			ReferendumID: referendumID,
			VoterID:      voterID,
			Vote:         vote,
		},
	)
	return err
}
