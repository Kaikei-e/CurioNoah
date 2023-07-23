package adaptor

import (
	"github.com/labstack/echo/v4"
	"insightstream/dependencyInversion"
	"insightstream/domain/baseFeeds"
	"insightstream/ent"
	"net/http"
	"strconv"
)

func InfiniteFetching(c echo.Context, cl *ent.Client) error {
	if c == nil {
		return http.ErrBodyNotAllowed
	}

	qpStr := c.QueryParam("page")
	qp, err := strconv.Atoi(qpStr)

	diFC := dependencyInversion.NewFeedCollection()
	fds, hadExceeded, err := diFC.FetchInfinite(qp, cl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if hadExceeded {
		err = c.JSON(200, response{fds, hadExceeded})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	} else {
		err = c.JSON(200, response{fds, false})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err)
		}
	}

	return nil
}

type response struct {
	Feeds       []baseFeeds.EachFeed `json:"baseFeeds"`
	HadExceeded bool                 `json:"hadExceeded"`
}
