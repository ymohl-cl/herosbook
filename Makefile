env = GOOS=linux GOARCH=amd64

all:
	 $(env) go build -o Bin/api cmd/api

.PHONY: all
