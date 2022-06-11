#!/bin/bash
go test ./... -coverprofile cover.out.tmp -v
cat cover.out.tmp | grep -v "_mock.go" > cover.out
go tool cover -html=cover.out