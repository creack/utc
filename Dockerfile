FROM    golang:1.6

ADD     . $GOPATH/src/github.com/creack/utc

RUN     go install github.com/creack/utc

ENTRYPOINT ["utc"]
