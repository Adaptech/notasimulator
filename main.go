package main

import (
	"flag"
	"github.com/golang-plus/uuid"
	"log"
)

type organization struct {
	Name string `json:"name"`
}

type referendum struct {
	Name     string   `json:"name"`
	Proposal string   `json:"proposal"`
	Options  []string `json:"options"`
}

type Election struct {
	Organization        organization
	Referendum          referendum
	NoOfVoters          int      `json:"noOfVoters"`
	PercentVotesPerHour []string `json:"percentVotesPerHour"`
}

type user struct {
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

type address struct {
	StreetAddress       string `json:"streetAddress"`
	PostOfficeBoxNumber string `json:"postOfficeBoxNumber"`
	AddressLocality     string `json:"addressLocality"`
	AddressRegion       string `json:"addressRegion"`
	PostalCode          string `json:"postalCode"`
	AddressCountry      string `json:"addressCountry"`
}

var (
	notaAddr string
	gesAddr  string
)

func parseFlags() {
	flag.IntVar(&duration, "d", 10, "Duration of election")
	flag.StringVar(&notaAddr, "notaAddr", "http://localhost:3001", "NOTA Address")
	flag.StringVar(&gesAddr, "gesAddr", "tcp://127.0.0.1:1113", "EventStore Address")
	flag.Parse()
}

func main() {
	parseFlags()

	users, err := readUserData("data/users.json")
	if err != nil {
		log.Fatalf("Error reading user data: %v", err)
	}

	log.Println("Users/Voters:", len(users))

	election, err := readElectionData("data/election.json")
	if err != nil {
		log.Fatalf("Error reading election data: %v", err)
	}

	log.Println(election.Organization.Name)
	log.Println(election.NoOfVoters)
	log.Println("Creating Election Admin ...")

	electionAdminID, err := createElectionAdmin("admin")
	if err != nil {
		log.Fatalf("Error creating election admin %v: %v", electionAdminID, err)
	}

	log.Println("Creating Organization ...")

	organizationGUID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Error creating organization UUID: %v", err)
	}

	newOrganizationID, err := createOrganization(organizationGUID.String(), election.Organization.Name)
	if err != nil {
		log.Fatalf("Error creating organization: %v", err)
	}

	log.Println("Creating Referendum ...")

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

	log.Println(election.Referendum.Name)
	log.Println(election.Referendum.Proposal)
	log.Println("Registering Voters ...")

	go errorStream()

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
				createErr("CreatingVoterError", err)
			}
		} else {
			break
		}
	}

	log.Println("Opening Polls ...")

	if err := openPolls(newReferendumID); err != nil {
		createErr("OpeningPollsError", err)
	}

	log.Println("Authenticating voters and casting votes ...")

	options := append(election.Referendum.Options, "None of the above")

	for count, option := range options {
		log.Printf("   %v: %v\n", count, option)
	}

	runElection(users, election.NoOfVoters, newReferendumID, newOrganizationID, options)

	if err := closePolls(newReferendumID); err != nil {
		createErr("ClosingPollsError", err)
	}
}
