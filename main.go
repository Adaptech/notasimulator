package main

import (
	"fmt"

	"github.com/nu7hatch/gouuid"
)

func main() {

	users := readUserData("data/users.json")
	fmt.Println("Users/Voters:")
	fmt.Println(len(users))

	election := readElectionData("data/election.json")
	fmt.Println(election.Organization.Name)
	fmt.Println(election.Referendum.Name)
	fmt.Println(election.Referendum.Proposal)
	fmt.Println(election.NoOfVoters)

	fmt.Printf("Creating Election Admin ...")
	const electionAdminID = "admin"
	createElectionAdmin(electionAdminID)

	fmt.Printf("Creating Organization ...")
	aGUID, _ := uuid.NewV4()
	newOrganizationID := aGUID.String()
	createOrganization(newOrganizationID, election.Organization.Name)

	fmt.Printf("Creating Referendum ...")
	referendum := election.Referendum
	a2ndGUID, _ := uuid.NewV4()
	newReferendumID := a2ndGUID.String()
	createReferendum(newReferendumID, newOrganizationID, referendum.Name, referendum.Proposal, referendum.Options)

	fmt.Printf("Registering Voters ...")

	for cnt, user := range users {
		if cnt <= election.NoOfVoters {
			registerVoter(user.ID, newOrganizationID, user.Firstname, user.Lastname, user.StreetAddress, "", user.AddressLocality, user.AddressRegion, user.PostalCode, user.AddressCountry)
		} else {
			break
		}
	}

	fmt.Printf("Opening Polls ...")
	openPolls(newReferendumID)

	fmt.Printf("Authenticating voters and casting votes ...")
	for cnt, user := range users {
		if cnt <= election.NoOfVoters {
			authenticateVoter(newReferendumID, user.ID, newOrganizationID)

		} else {
			break
		}
	}
	closePolls(newReferendumID)
}

type organization struct {
	Name string `json:"name"`
}

type Referendum struct {
	Name     string   `json:"name"`
	Proposal string   `json:"proposal"`
	Options  []string `json:"options"`
}

type Election struct {
	Organization        organization
	Referendum          Referendum
	NoOfVoters          int      `json:"noOfVoters"`
	PercentVotesPerHour []string `json:"percentVotesPerHour"`
}

type User struct {
	ID              string `json:"userId"`
	Firstname       string `json:"firstname"`
	Lastname        string `json:"lastname"`
	StreetAddress   string `json:"streetAddress"`
	AddressLocality string `json:"addressLocality"`
	AddressRegion   string `json:"addressRegion"`
	PostalCode      string `json:"postalCode"`
	AddressCountry  string `json:"addressCountry"`
	Email           string `json:"email"`
	Password        string `json:"password"`
}
