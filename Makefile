.PHONY:buildlinux buildmac
build: buildmac buildlinux

buildmac:
	env GOOS=darwin GOARCH=amd64 go build -o bin/ocs-mac .

buildlinux:
	env GOOS=linux GOARCH=amd64 go build -o bin/ocs-linux .
