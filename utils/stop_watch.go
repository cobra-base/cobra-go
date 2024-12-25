package utils

import "time"

type StopWatch struct {
    start time.Time
    end   time.Time
}

func NewStopWatch() *StopWatch {
    sw := &StopWatch{
        start: time.Now(),
    }
    return sw
}

func (s *StopWatch) Start() {
    s.start = time.Now()
}

func (s *StopWatch) Stop() {
    s.end = time.Now()
}

func (s *StopWatch) Duration() int64 {
    return s.end.Sub(s.start).Milliseconds()
}
