package routes

import (
	"github.com/labstack/echo"

	"github.com/wanghantao11/log-receiver/internal/pkg/log"
)

func RoutesV1(e *echo.Echo, logService *log.Service) {
	eg := e.Group("/v1")
	SetupLogRoutes(eg.Group("/logs"), logService)
}
