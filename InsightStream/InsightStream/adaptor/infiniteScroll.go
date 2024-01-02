package adaptor

import (
	"fmt"
	"insightstream/dependencyInversion"
	"insightstream/domain/baseFeeds"
	"insightstream/ent"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func InfiniteFetching(c echo.Context, cl *ent.Client) error {
	if c == nil {
		return http.ErrBodyNotAllowed
	}

	qpStr := c.QueryParam("page")
	qp, err := strconv.Atoi(qpStr)
	if err != nil {
		slog.Error("failed to convert query param: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to convert query param: %w", err))
	}

	diFC := dependencyInversion.NewFeedCollection()
	fds, hadExceeded, err := diFC.FetchInfinite(qp, cl)
	if err != nil {
		slog.Error("failed to load infinite feeds: %v", err)
		return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to load infinite feeds: %w", err))
	}

	if hadExceeded {
		err = c.JSON(200, response{fds, hadExceeded})
		if err != nil {
			slog.Error("failed to return json: %v", err)
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to return json"))
		}
	} else {
		err = c.JSON(200, response{fds, false})
		if err != nil {
			slog.Error("failed to return json: %v", err)
			return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to return json"))
		}
	}

	return nil
}

type response struct {
	Feeds       []baseFeeds.EachFeed `json:"feeds"`
	HadExceeded bool                 `json:"hadExceeded"`
}
