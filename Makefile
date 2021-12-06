.PHONY:buildlinux buildmac buildarm
build: buildmac buildlinux buildarm

buildmac:
	env GOOS=darwin GOARCH=amd64 go build -o bin/ocs_mac_amd64 .

buildlinux:
	env GOOS=linux GOARCH=amd64 go build -o bin/ocs_linux_amd64 .

buildarm:
	env GOOS=linux GOARCH=arm go build -o bin/ocs_linux_arm .
