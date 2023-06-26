package adaptor

import (
	"github.com/labstack/echo/v4"
	"insightstream/dependencyInversion"
	"insightstream/ent"
	"insightstream/models/feeds"
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
		return err
	}

	if hadExceeded {
		err = c.JSON(200, response{fds, hadExceeded})
		if err != nil {
			return err
		}
	} else {
		err = c.JSON(200, response{fds, false})
		if err != nil {
			return err
		}
	}

	err = c.JSON(http.StatusInternalServerError, "failed to fetch feeds")

	return err
}

type response struct {
	Feeds       []feeds.EachFeed `json:"feeds"`
	HadExceeded bool             `json:"hadExceeded"`
}
