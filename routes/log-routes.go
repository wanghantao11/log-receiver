package routes

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/wanghantao11/log-receiver/internal/pkg/log"
)

func SetupLogRoutes(e *echo.Group, service *log.Service) {
	// Insert logs from external
	e.POST("", func(c echo.Context) error {
		msg := []log.Log{}
		err := c.Bind(&msg)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		err = service.AddLogs(msg)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	})
}
