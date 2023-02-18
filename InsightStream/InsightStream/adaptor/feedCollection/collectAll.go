package feedCollection

import (
	"errors"
	"github.com/labstack/echo/v4"
)

func CollectAll(c echo.Context) error {
	// TODO

	if c == nil {
		return errors.New("context is nil")
	}

	err := c.JSON(200, "success")
	if err != nil {
		return err
	}

	return nil
}
