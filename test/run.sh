#! /bin/bash 
set -e -x

cd ..
go install
sheetdrop db migrate -d
sheetdrop seed

cd controllers
go test -v ./... -coverprofile=cover.out
go tool cover -html=cover.out