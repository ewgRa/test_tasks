FROM golang:1.15.5

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /app/cmd/api

RUN go build

ENTRYPOINT ["./api"]
