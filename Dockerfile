FROM golang:1.14

RUN mkdir /go/src/talkni
WORKDIR /go/src/talkni

COPY go.mod /go/src/talkni/
COPY main.go /go/src/talkni/
COPY index.html /go/src/talkni/

CMD ["go", "run", "main.go"]

EXPOSE 8080
