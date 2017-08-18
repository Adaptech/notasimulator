#build the go programs
FROM golang:1.8.3-alpine3.6 as builder
RUN apk update && \
        apk add git curl unzip && \
        curl -L -o dep.zip https://github.com/golang/dep/releases/download/v0.3.0/dep-linux-amd64.zip && \
        echo '96c191251164b1404332793fb7d1e5d8de2641706b128bf8d65772363758f364  dep.zip' | sha256sum -c - && \
        unzip -d /usr/bin dep.zip && rm dep.zip
WORKDIR /go/src/Adaptech/notasimulator/
COPY Gopkg.toml Gopkg.lock authenticateVoter.go castVote.go closePolls.go createElectionAdmin.go createOrganization.go createReferendum.go jsonReq.go main.go openPolls.go readElectionData.go readUserData.go registerVoter.go runElection.go sendError.go ./
RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 GOOS=linux go build

#create an image that performs deployments
FROM alpine:3.6
WORKDIR /app/tmp
RUN mkdir -p /app/bin && apk update
COPY data /app/tmp/data
COPY --from=builder /go/src/Adaptech/notasimulator/notasimulator /app/bin/
CMD [ "/app/bin/notasimulator" ]
