.ONESHELL:
env = GOOS=linux GOARCH=amd64

APP_NAME=api-editor
IGNORED_FOLDER=.ignore
COVERAGE_FILE=${IGNORED_FOLDER}/coverage.txt
BIN_FOLDER=bin

MODULE_NAME := $(shell go list -m)

.PHONY: all
all: install lint test build

.PHONY: install
install:
	go mod download

.PHONY: lint
lint:
	golint ./...

.PHONY: test
test:
	mkdir -p ${IGNORED_FOLDER}
	go test -race -coverprofile=${COVERAGE_FILE} -covermode=atomic ./...

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ${BIN_FOLDER}/app ${MODULE_NAME}/cmd/${APP_NAME}

.PHONY: clean
clean:
	if [ -f ${COVERAGE_FILE} ]; then
		rm -rf ${COVERAGE_FILE}
	fi

.PHONY: fclean
fclean: clean
	rm -rf ${BIN_FOLDER}

.PHONY: tools
tools:
	go get -u golang.org/x/lint/golint
	go get github.com/golang/mock/gomock
	go install github.com/golang/mock/mockgen
