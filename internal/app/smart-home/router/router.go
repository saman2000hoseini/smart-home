package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/saman2000hoseini/smart-home/internal/app/smart-home/config"
	"github.com/sirupsen/logrus"
)

// New creates a new application router.
func New(cfg config.Config) *echo.Echo {
	e := echo.New()

	debug := logrus.IsLevelEnabled(logrus.DebugLevel)

	e.Debug = debug

	e.HideBanner = true

	if !debug {
		e.HidePort = true
	}

	e.Server.ReadTimeout = cfg.HTTPServer.ReadTimeout
	e.Server.WriteTimeout = cfg.HTTPServer.WriteTimeout

	recoverConfig := middleware.DefaultRecoverConfig
	recoverConfig.DisablePrintStack = !debug
	e.Use(middleware.RecoverWithConfig(recoverConfig))

	e.Use(middleware.CORS())

	return e
}
