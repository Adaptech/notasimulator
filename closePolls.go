package main

type closePollsRequest struct {
	ReferendumID string `json:"referendumId"`
}

func closePolls(referendumID string) error {
	err := jsonReq(
		"/api/v1/organization/referendum/polls/close",
		closePollsRequest{
			ReferendumID: referendumID,
		},
	)
	return err
}
