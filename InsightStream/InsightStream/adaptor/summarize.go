package adaptor

//
//import (
//	"fmt"
//	"github.com/labstack/echo/v4"
//	"insightstream/ent"
//	"log/slog"
//	"net/http"
//	"strconv"
//)
//
//func SummarizeEachFeedToday(c echo.Context, cl *ent.Client) interface{} {
//	if c == nil {
//		return http.ErrBodyNotAllowed
//	}
//
//	qpStr := c.QueryParam("page")
//	qp, err := strconv.Atoi(qpStr)
//	if err != nil {
//		slog.Error("failed to convert query param: %v", err)
//		return c.JSON(http.StatusInternalServerError, fmt.Errorf("failed to convert query param: %w", err))
//	}
//
//	return nil
//}
