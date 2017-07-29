package main

type OpenPollsRequest struct {
	ReferendumId string `json:"referendumId"`
}

func openPolls(referendumID string) error {
	err := jsonReq(
		"http://localhost:3001/api/v1/organization/referendum/polls/open",
		OpenPollsRequest{
			ReferendumId: referendumID,
		},
	)
	return err
}
