package main

import (
	"flag"
	"os"

	"github.com/labstack/echo"

	"github.com/ymohl-cl/herosbook/cmd/api-nomenclature/handler"
	"github.com/ymohl-cl/herosbook/pkg/app"
)

var appName = flag.String("appName", "heroesbook_nomenclature", "application name")

func init() {
	flag.Parse()
	if appName == nil {
		flag.PrintDefaults()
		os.Exit(-1)
	}
}

func main() {
	var err error
	var server app.App
	var subRouter *echo.Group

	if server, err = app.New(*appName); err != nil {
		panic(err)
	}
	if subRouter, err = server.SubRouter("/api", true); err != nil {
		panic(err)
	}
	if err = handler.New(*appName, subRouter); err != nil {
		panic(err)
	}
	if err = server.Start(); err != nil {
		panic(err)
	}
	return
}
