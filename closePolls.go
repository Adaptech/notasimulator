package main

type ClosePollsRequest struct {
	ReferendumId string `json:"referendumId"`
}

func closePolls(referendumID string) error {
	err := jsonReq(
		"http://localhost:3001/api/v1/organization/referendum/polls/close",
		ClosePollsRequest{
			ReferendumId: referendumID,
		},
	)
	return err
}
