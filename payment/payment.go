package main

import (
	"fmt"
	"sync"
	"time"
)

type Payment struct {
	From   string
	To     string
	Amount float64 // USD

	once sync.Once
}

func (p *Payment) Process() {
	t := time.Now()
	p.once.Do(func() {
		p.process(t)
	})
}

func (p *Payment) process(t time.Time) {
	ts := t.Format(time.RFC3339)
	fmt.Printf("[%s] -> $%.2f -> %s\n", ts, p.Amount, p.To)
}

func main() {
	p := Payment{
		From:   "Wile. E. Coyote",
		To:     "ACME",
		Amount: 321.123,
	}
	p.Process()
	p.Process()
}
