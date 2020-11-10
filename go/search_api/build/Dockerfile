FROM golang:1.12.9

WORKDIR /go-test-task

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go-test-task/cmd/api

RUN go build

ENTRYPOINT ["./api"]
