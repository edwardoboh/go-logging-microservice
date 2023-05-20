package main

import (
	"context"
	"fmt"
	"time"
)

type LoggingService struct {
	next Service
}

func NewLoggingService(next Service) Service {
	return &LoggingService{next: next}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(tm time.Time) {
		fmt.Printf("fact=%+v error=%v time=%v \n", fact, err, time.Since(tm))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
