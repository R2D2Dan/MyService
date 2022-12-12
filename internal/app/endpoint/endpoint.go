package endpoint

// Вывыод результата
import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Service interface {
	CurrnetDate() string
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

	date := e.sr.CurrnetDate()
	result := fmt.Sprintf("Current time in Moscow: %s", date)

	if err := ctx.String(http.StatusOK, result); err != nil {
		log.Println("/status return error...", err)
	}
	return nil

}
