package main

import (
	"fmt"
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

func main() {

	users := readUserData("data/users.json")
	fmt.Println("Users/Voters:")
	fmt.Println(len(users))

	election := readElectionData("data/election.json")
	fmt.Println(election.Organization.Name)
	fmt.Println(election.Referendum.Name)
	fmt.Println(election.Referendum.Proposal)
	fmt.Println(election.Referendum.Options)
	fmt.Println(election.PercentVotesPerHour)

}
