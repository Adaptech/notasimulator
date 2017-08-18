package main

import (
	"log"
	"math/rand"
	"time"
)

var (
	votes          int
	bursts         = 4
	duration       int
	durationSec    int
	offsetVotes    int
	burstVotes     int
	votesPerBurst  int
	remainderVotes int
	votesChan      = make(chan string)
)

func runBurstVotes() {
	waitTime := durationSec / bursts
	timer := time.Tick(time.Duration(waitTime))

	for b := 0; b < bursts; b++ {
		go func() {
			select {
			case <-timer:
				for x := 0; x < votesPerBurst; x++ {
					go func() {
						time.Sleep(time.Duration(rand.Intn(waitTime / 2)))
						votesChan <- "burst vote"
					}()
				}
			}
		}()
	}
}

func runRemainderBurstVotes() {
	if remainderVotes > 0 {
		randWait := rand.Intn(durationSec)
		time.Sleep(time.Duration(randWait))

		for x := 0; x < remainderVotes; x++ {
			go func() {
				time.Sleep(time.Duration(rand.Intn(randWait / 3)))
				votesChan <- "remaineder burst vote"
			}()
		}
	}
}

func runOffsetVotes() {
	for x := 0; x < offsetVotes; x++ {
		go func() {
			randWait := rand.Intn(durationSec)
			time.Sleep(time.Duration(randWait))
			votesChan <- "offset vote"
		}()
	}
}

func runElection(users []user, voterCount int, newReferendumID string, newOrganizationID string, options []string) {
	rand.Seed(time.Now().UTC().UnixNano())

	votes = voterCount
	durationSec = duration * 1000000000
	offsetVotes = rand.Intn(votes)
	burstVotes = votes - offsetVotes
	votesPerBurst = burstVotes / bursts
	remainderVotes = burstVotes % bursts

	go runOffsetVotes()
	go runBurstVotes()
	go runRemainderBurstVotes()

	for count, usr := range users {
		if count == votes {
			break
		}
		select {
		case <-votesChan:
			go func() {
				err := authenticateVoter(newReferendumID, usr.ID, newOrganizationID)
				if err != nil {
					createErr("AuthenticatingVoterError", err)
				}

				voteToCast := rand.Intn(len(options))

				log.Printf("Choice: %v, voting %v.\n", voteToCast, options[voteToCast])

				if err := castVote(newReferendumID, usr.ID, options[voteToCast]); err != nil {
					for x := 0; x < 10; x++ {
						createErr("CastingVoteError", err)
						time.Sleep(time.Duration(time.Second))
						if err := castVote(newReferendumID, usr.ID, options[voteToCast]); err == nil {
							break
						}
					}
				}
			}()
		}
	}

	log.Println("Votes sent!")
}
