run:
	go build -o space-race
	./space-race

build:
	go build -o space-race

all:
	go build -o space-race
	GOOS=linux GOARCH=amd64 go build -o space-race_linux
	GOOS=windows GOARCH=386 go build -o space-race.exe
