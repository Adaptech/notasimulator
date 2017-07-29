package main

import (
	"fmt"
	"github.com/golang-plus/uuid"
	"log"
	"math/rand"
	"time"
)

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

type Address struct {
	StreetAddress       string `json:"streetAddress"`
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber"`
	AddressLocality     string `json:"addressLocality"`
	AddressRegion       string `json:"addressRegion"`
	PostalCode          string `json:"postalCode"`
	AddressCountry      string `json:"addressCountry"`
}

func main() {
	users, err := readUserData("data/users.json")
	if err != nil {
		log.Fatalf("Error reading user data: %v", err)
	}

	fmt.Println("Users/Voters:", len(users))

	election, err := readElectionData("data/election.json")
	if err != nil {
		log.Fatalf("Error reading election data: %v", err)
	}

	fmt.Println(election.Organization.Name)
	fmt.Println(election.NoOfVoters)
	fmt.Println("Creating Election Admin ...")

	electionAdminID, err := createElectionAdmin("admin")
	if err != nil {
		log.Fatalf("Error creating election admin %v: %v", electionAdminID, err)
	}

	fmt.Println("Creating Organization ...")

	organizationGUID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Error creating organization UUID: %v", err)
	}

	newOrganizationID, err := createOrganization(organizationGUID.String(), election.Organization.Name)
	if err != nil {
		log.Fatalf("Error creating organization: %v", err)
	}

	fmt.Println("Creating Referendum ...")

	referendumGUID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Error creating referendum UUID: %v", err)
	}

	newReferendumID, err := createReferendum(
		referendumGUID.String(),
		newOrganizationID,
		election.Referendum.Name,
		election.Referendum.Proposal,
		election.Referendum.Options,
	)
	if err != nil {
		log.Fatalf("Error creating referendum: %v", err)
	}

	fmt.Println(election.Referendum.Name)
	fmt.Println(election.Referendum.Proposal)
	fmt.Println("Registering Voters ...")

	for count, user := range users {
		if count <= election.NoOfVoters {
			if err := registerVoter(
				user.ID,
				newOrganizationID,
				user.Firstname,
				user.Lastname,
				user.StreetAddress,
				"",
				user.AddressLocality,
				user.AddressRegion,
				user.PostalCode,
				user.AddressCountry,
			); err != nil {
				log.Printf("Error creating voter: %v", err)
			}
		} else {
			break
		}
	}

	fmt.Println("Opening Polls ...")

	if err := openPolls(newReferendumID); err != nil {
		log.Fatalf("Error opening polls: %v", err)
	}

	fmt.Println("Authenticating voters and casting votes ...")

	options := append(election.Referendum.Options, "None of the above")

	for count, option := range options {
		fmt.Printf("   %v: %v\n", count, option)
	}

	choices := len(options)

	rand.Seed(time.Now().UTC().UnixNano())

	for count, user := range users {
		if count < election.NoOfVoters {
			err := authenticateVoter(newReferendumID, user.ID, newOrganizationID)
			if err != nil {
				log.Printf("Error authenticating voter: %v", err)
			}

			voteToCast := rand.Intn(choices)

			fmt.Printf("Choice: %v, voting %v.\n", voteToCast, options[voteToCast])

			if err := castVote(newReferendumID, user.ID, options[voteToCast]); err != nil {
				log.Printf("Error casting vote: %v", err)
			}

		} else {
			break
		}
	}

	if err := closePolls(newReferendumID); err != nil {
		log.Fatalf("Error closing polls: %v", err)
	}
}
