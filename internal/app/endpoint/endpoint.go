package endpoint

// Вывыод результата
import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type Service interface {
	CurrnetDate() string
	Fibonacci(number int) int
}

type Endpoint struct {
	sr Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		sr: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	return nil
}

func (e *Endpoint) TimeInMoscow(ctx echo.Context) error {

	date := e.sr.CurrnetDate()
	result := fmt.Sprintf("Current time in Moscow: %s", date)

	if err := ctx.String(http.StatusOK, result); err != nil {
		log.Println("/status return error...", err)
	}
	return nil

}

func (e *Endpoint) Fibonacci(ctx echo.Context) error {

	number, err := strconv.Atoi(ctx.Param("number"))
	if err != nil {
		ctx.String(202, "not number param...")
		return nil
	}

	if err := ctx.String(http.StatusOK, strconv.Itoa(e.sr.Fibonacci(number))); err != nil {
		log.Println("/status return error...", err)
	}
	return nil
}
