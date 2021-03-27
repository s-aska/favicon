package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/s-aska/favicon/webapp/web/appengine"
	stackdriverlog "github.com/yfuruyama/stackdriver-request-context-log"
)

// Bootstrap ...
func Bootstrap() {
	e := echo.New()

	config := stackdriverlog.NewConfig(appengine.ProjectID)
	config.RequestLogOut = os.Stderr              // request log to stderr
	config.ContextLogOut = os.Stdout              // context log to stdout
	config.Severity = stackdriverlog.SeverityInfo // only over INFO logs are logged

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		logger := stackdriverlog.RequestContextLogger(c.Request())
		logger.Errorf("[HTTPErrorHandler] error:%v", err)
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == http.StatusInternalServerError {
				c.JSON(he.Code, "Internal Server Error")
			} else {
				c.JSON(he.Code, he.Message)
			}
		} else {
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
		}
	}
	registerTemplates(e)
	registerRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Infof("Defaulting to port %s", port)
	}
	log.Infof("Listening on port %s", port)
	for _, e := range os.Environ() {
		log.Infof("%s", e)
	}

	handler := stackdriverlog.RequestLogging(config)(e)

	http.Handle("/", handler)

	if appengine.IsDevAppServer() {
		e.Static("/static", "app/static")
		e.Static("/favicon.ico", "app/static/favicon.ico")
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}
