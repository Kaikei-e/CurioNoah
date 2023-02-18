package adaptor

import (
	"errors"
	"github.com/labstack/echo/v4"
	"insightstream/dependencyInversion"
	"insightstream/ent"
)

func CollectAll(c echo.Context, cl *ent.Client) error {
	// TODO

	if c == nil {
		return errors.New("context is nil")
	}

	f := dependencyInversion.NewFeedCollection()
	err := f.FetchAll(cl)

	err = c.JSON(200, "success")
	if err != nil {
		return err
	}

	return nil
}
