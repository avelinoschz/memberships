FROM golang:1.14-alpine3.11

LABEL maintainer="Avelino Sanchez <avelino.schz@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY internal internal

COPY cmd cmd

RUN go build -o /memberships-api cmd/memberships-api/main.go 

EXPOSE 8000

CMD ["./memberships-api"]