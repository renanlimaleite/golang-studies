## para buildar normal
1. go build main.go

## para buildar para windowns
1. GOOS=windows go build main.go -> gera um main.exe

## para buildar para linux
1. GOOS=linux go build main.go -> gera um main

## para buildar para mac
1. GOOS=darwin go build main.go -> gera um main

## para buildar para mac goarch=amd64
1. GOOS=darwin GOARCH=amd64 go build main.go -> gera um main

## para buildar para windows goarch=amd64
1. GOOS=windows GOARCH=amd64 go build main.go -> gera um main.exe