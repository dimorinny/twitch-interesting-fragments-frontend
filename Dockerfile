FROM golang:1.8

WORKDIR /go/src/github.com/dimorinny/twitch-interesting-fragments-frontend/

ADD . .

RUN go get && go install

CMD /go/bin/twitch-interesting-fragments-frontend