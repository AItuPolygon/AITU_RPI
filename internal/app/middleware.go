package app

import (
	"context"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	TaskKey    = "task"
	ElementKey = "element"
)

func TaskElementExtractMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Path()
		parts := strings.SplitN(strings.Trim(path, "/"), "/", 3)
		fmt.Println(parts)
		// Ensure there are at least three parts for task and element
		if len(parts) >= 3 {
			task := parts[1]
			element := parts[2]

			// Store task and element in the context
			ctx := context.WithValue(c.Request().Context(), TaskKey, task)
			ctx = context.WithValue(ctx, ElementKey, element)
			c.SetRequest(c.Request().WithContext(ctx))
		}

		return next(c)
	}
}
