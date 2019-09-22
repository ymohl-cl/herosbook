# Server

[![Source](https://img.shields.io/badge/git-source-orange.svg?style=flat-square)](https://github.com/ymohl-cl/gopkg/tree/server-release/server)
[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/ymohl-cl/gopkg/server)
[![Build Status](https://travis-ci.org/ymohl-cl/gopkg.svg?branch=server-release&style=flat-square)](https://travis-ci.org/ymohl-cl/gopkg)
[![codecov](https://codecov.io/gh/ymohl-cl/gopkg/branch/server-release/graph/badge.svg?style=flat-square)](https://codecov.io/gh/ymohl-cl/gopkg)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/ymohl-cl/gopkg/server-release/LICENSE)

server package provide a common server http.

## Requirements

Golang 1.12.4 or higher

download the package:

``` bash
go get -u github.com/ymohl-cl/gopkg/server
```

environment variables:

``` bash
export MYAPP_SSL_ENABLE="true"
export MYAPP_SSL_CERTIFICATE="path_to_certificats/cert.pem"
export MYAPP_SSL_KEY="path_to_key/key.pem"
export MYAPP_PORT="4242"
```

## Usage

``` Golang
import "github.com/ymohl-cl/gopkg/server"

func main() {
    s, err := server.New("myapp")
    if err != nil {
        panic(err)
    }
    if err = s.Start(); err != nil {
        panic(err)
    }
}
```

## Changelog

### v1.1.0

Account management

- server provide authentication with jwt token
- implement user register with postgres (postgres driver is temporary in server package)
- transaction manage users behaviors

### v1.0.0

Initial commit

- server http with tls in option
- implement ping default route
- logs activity is on by default
- tests and documentation
