package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/wanghantao11/log-receiver/config"
	logservice "github.com/wanghantao11/log-receiver/internal/pkg/log"
	"github.com/wanghantao11/log-receiver/internal/pkg/log/storer"
	"github.com/wanghantao11/log-receiver/routes"
)

const serviceName = "log-receiver"

func main() {
	config.Init(serviceName)

	// Setup DB
	logStore := storer.NewMemory()

	e := echo.New()

	// Logging
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)

	// Configuration
	apiIP := config.Get(config.APIIP)
	apiPort := config.Get(config.APIPORT)

	// Initialize log service
	logService := logservice.New(logStore)

	// Setup routes v1
	routes.RoutesV1(e, logService)

	e.Logger.Debug(e.Start(apiIP + ":" + apiPort))
}
