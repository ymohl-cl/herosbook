env = GOOS=linux GOARCH=amd64

all:
	 $(env) go build -o Bin/herosbook cmd/api

.PHONY: all
