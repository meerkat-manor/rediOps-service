package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"

	"merebox.com/rediops/api"
)

func main() {
	var port = flag.Int("port", 8075, "Port for HTTP server micro service")
	var dataFolder = flag.String("data", "./data", "Data folder")
	var staticFolder = flag.String("static", "./static", "Static folder")
	var configFolder = flag.String("configuration", "./config", "Configuration folder")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	configFilename := *configFolder + "/rediops.yaml"
	var ro = api.NewRediops(*configFolder, configFilename, *dataFolder)

	e := echo.New()
	// Log all requests
	e.Use(echomiddleware.Logger())
	// Use our validation middleware to check all requests against the
	// OpenAPI schema.
	var options middleware.Options
	// Skip static assets such as HTML, Images, CSS, etc
	options.Skipper = func(c echo.Context) bool {
		if (strings.HasSuffix(c.Request().URL.Path, ".html") || strings.HasSuffix(c.Request().URL.Path, ".json")) {
		  return true
		}
		if (strings.HasSuffix(c.Request().URL.Path, ".ico")) {
			return true
        }
		if (c.Request().URL.Path =="/") {
			return true
  		}
		  if strings.HasPrefix(c.Request().URL.Path, "/assets/") {
			return true
		}
		
		return false
	}
	e.Use(middleware.OapiRequestValidatorWithOptions(swagger, &options))

	api.RegisterHandlers(e, ro)

	e.File("/rediops.json", (*configFolder + "/rediops.json"))
	e.Static("/", *staticFolder)

	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
