# Online Voting Simulator for NOTA

*This simulates an election for the "None Of The Above" (NOTA) online voting system. (https://github.com/Adaptech/nota)*


## Getting Started

1. Start NOTA and https://geteventstore.com on localhost as described at https://github.com/Adaptech/nota
2. (Optional) Edit ```data/election.json```. The maximum noOfVoters in an election is 20,000.
3. ```go build && ./notasimulator```.
4. View election results at http://localhost:3001/api/v1/r/results

*The fictitious 20,000 voters in data/users.json are courtesy of https://github.com/benkeen/generatedata.*