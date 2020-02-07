build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o bin/darwin/git-cloner .
