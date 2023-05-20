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
	return &LoggingService{next}
}

func (s *LoggingService) GetCatFact(ctx context.Context) (fact *CatFact, err error) {
	defer func(tm time.Time) {
		fmt.Printf("fact=%v error=%v time=%v", fact, err.Error(), time.Since(tm))
	}(time.Now())

	return s.next.GetCatFact(ctx)
}
