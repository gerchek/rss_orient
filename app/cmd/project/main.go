package main

import (
	"fmt"
	"rss/internal/app"
	"rss/internal/setup/constructor"
	"rss/internal/setup/routes"
	"rss/pkg/gormclient"
	"rss/pkg/logging"
)

func main() {
	logger := logging.Log()
	client, err := gormclient.NewClient()
	if err != nil {
		logger.Warn(err)
	}

	logger.Info("setting up all repository, service, controller...")
	constructor.SetConstructor(client, logger)

	logger.Info("initializing a new app...")
	fmt.Println("hi")
	var r routes.MyRoute
	go r.Routes()

	app.NewApp(logger)

}
