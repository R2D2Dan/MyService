package middleware

// Промежуточный проверки
import (
	"strings"

	"github.com/labstack/echo"
)

func ChekUser(next echo.HandlerFunc) echo.HandlerFunc {

	return func(ctx echo.Context) error {

		user := ctx.Request().Header.Get("User")
		if !strings.EqualFold(user, "admin") {
			ctx.Response().Writer.Write([]byte("not admin"))
			return nil
		}

		err := next(ctx)
		if err != nil {
			return err
		}
		return nil
	}

}
