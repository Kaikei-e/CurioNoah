package feedCollection

import "github.com/labstack/echo/v4"

func CollectAll(c echo.Context) error {
	// TODO

	err := c.JSON(200, "success")
	if err != nil {
		return err
	}

	return nil
}
