package feedCollection

import (
	"errors"
	"github.com/labstack/echo/v4"
	"insightstream/dependencyInversion"
)

func CollectAll(c echo.Context) error {
	// TODO

	if c == nil {
		return errors.New("context is nil")
	}

	f := dependencyInversion.NewFeedCollection()
	err := f.FetchAll()

	err = c.JSON(200, "success")
	if err != nil {
		return err
	}

	return nil
}
