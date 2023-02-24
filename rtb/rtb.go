package main

import (
	"context"
	"fmt"
	"strings"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	url := "http://go.dev"
	bid := bidOne(ctx, url)
	fmt.Println(bid)
}

type Bid struct {
	AdURL string
	Price int
}

func bidOne(ctx context.Context, url string) Bid {
	ch := make(chan Bid)

	go func() {
		ch <- bestBid(url)
	}()

	select {
	case bid := <-ch:
		return bid
	case <-ctx.Done():
		return Bid{}
	}
}

func bestBid(url string) Bid {
	// simulate some work
	d := 100 * time.Millisecond
	if strings.HasPrefix(url, "https://") {
		d = 20 * time.Millisecond
	}
	time.Sleep(d)

	return Bid{
		AdURL: "http://adsRus.com/default",
		Price: 7,
	}
}
