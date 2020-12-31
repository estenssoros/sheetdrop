#! /bin/bash 
set -e -x

cd ..
go install
sheetdrop db migrate -d
sheetdrop seed

go test -v ./controllers -coverprofile=cover.out
go tool cover -html=cover.out