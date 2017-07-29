package main

type CastVoteRequest struct {
	ReferendumId string `json:"referendumId"`
	VoterId      string `json:"voterId"`
	Vote         string `json:"vote"`
}

func castVote(referendumID string, voterID string, vote string) error {
	err := jsonReq(
		"http://localhost:3001/api/v1/organization/referendum/vote",
		CastVoteRequest{
			ReferendumId: referendumID,
			VoterId:      voterID,
			Vote:         vote,
		},
	)
	return err
}
