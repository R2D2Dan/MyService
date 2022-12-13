package service

// Логика сервиса
import (
	"time"
)

const my_layout = "02.01.2006 15:04:05"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) CurrnetDate() string {
	date := time.Now()
	return date.Format(my_layout)
}

func (s *Service) Fibonacci(number int) int {
	if number == 0 {
		return 1
	}
	return number * s.Fibonacci(number-1)

}
