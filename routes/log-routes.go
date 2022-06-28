package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"

	"github.com/wanghantao11/log-receiver/internal/pkg/log"
)

func SetupLogRoutes(e *echo.Group, service *log.Service) {
	// swagger:operation POST /v1/logs LogMsg
	//
	// Create logs.
	//
	// ---
	// tags: [LogMsg]
	// parameters:
	//   - in: body
	//     name: logs
	//     description: logs payload to be created.
	//     schema:
	//       type: object
	//       required:
	//         - id
	//         - l
	//         - t
	//       properties:
	//         id:
	//           description: Log ID.
	//           type: string
	//         l:
	//           description: Log Level.
	//           type: string
	//         m:
	//           description: Log Message.
	//           type: string
	//         t:
	//           description: Log Timestamp.
	//           type: string
	// security:
	//   - JWTAuth: []
	// responses:
	//   '204':
	//     description: Logs are added successfully.
	//   '400':
	//     description: >
	//       Error codes:
	//         - `ErrBadRequest`: Bad request.
	//   '500':
	//     description: Internal server error
	e.POST("", func(c echo.Context) error {
		msg := []log.LogMsg{}
		err := c.Bind(&msg)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		logs, err := log.ToLogs(msg)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		logs, err = service.CreateLogs(logs)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		fmt.Printf("%v", logs)

		return c.NoContent(http.StatusNoContent)
	})

	// swagger:operation GET /v1/logs
	//
	// Fetch all logs based on given time period.
	//
	// Returns an array of logs for the given time period.
	//
	// ---
	// tags: [logs]
	// security:
	//   - JWTAuth: []
	// responses:
	//   '200':
	//     description: The success response.
	//   '400':
	//     description: >
	//       Error codes:
	//         - `ErrBadRequest`: Bad request.
	//   '500':
	//     description: Internal server error
	e.GET("", func(c echo.Context) error {
		// Parse optional query params "from" and "to"
		fromTimeString := c.QueryParam("from")
		toTimeString := c.QueryParam("to")
		var fromTime, toTime time.Time

		if fromTimeString != "" {
			fromTimestamp, err := strconv.ParseInt(fromTimeString, 10, 64)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			fromTime = time.Unix(fromTimestamp, 0)
		}

		if toTimeString != "" {
			toTimestamp, err := strconv.ParseInt(toTimeString, 10, 64)
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
			toTime = time.Unix(toTimestamp, 0)
		}

		result, err := service.GetLogs(fromTime, toTime)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	})
}
