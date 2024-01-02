package adaptor

import (
	"errors"
	"fmt"
	"insightstream/dependencyInversion"
	"insightstream/ent"
	"log/slog"

	"github.com/labstack/echo/v4"
)

func CollectAll(c echo.Context, cl *ent.Client) error {

	if c == nil {
		slog.Error("context is nil")
		return fmt.Errorf("context is nil: %w", errors.New("context is nil"))
	}

	f := dependencyInversion.NewFeedCollection()
	err := f.FetchAll(cl)
	if err != nil {
		slog.Error("failed to fetch all: %v", err)
		return fmt.Errorf("failed to fetch all: %w", err)
	}

	err = c.JSON(200, "success")
	if err != nil {
		slog.Error("failed to return json: %v", err)
		return fmt.Errorf("failed to return json: %w", err)
	}

	return nil
}
