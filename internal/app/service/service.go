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
