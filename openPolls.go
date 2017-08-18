package main

type openPollsRequest struct {
	ReferendumID string `json:"referendumId"`
}

func openPolls(referendumID string) error {
	err := jsonReq(
		"/api/v1/organization/referendum/polls/open",
		openPollsRequest{
			ReferendumID: referendumID,
		},
	)
	return err
}
