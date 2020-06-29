#!/bin/bash
# exports all the neccesary environment variables
export $(cat configs/example.env | xargs)

# prepare all dependencies needed
deploy/docker-compose up

# run the main program
# TODO change it to run via docker
go run ./cmd/memberships-api/main.go