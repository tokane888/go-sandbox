package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func generator(ctx context.Context, num int) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()
		defer close(out)
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("generator() done")
				fmt.Println(context.Cause(ctx))
				if err := ctx.Err(); errors.Is(err, context.Canceled) {
					fmt.Println("canceled")
				} else if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("deadline exceeded")
				}
				return
			case <-ticker.C:
				select {
				case <-ctx.Done():
					fmt.Println("generator() done")
					return
				case out <- num:
				}
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithTimeoutCause(context.Background(), 1500*time.Millisecond, errors.New("timeout"))
	wg.Add(1)
	out := generator(ctx, 1)
LOOP:
	for range 3 {
		result, ok := <-out
		if ok {
			fmt.Println(result)
		} else {
			fmt.Println("timeout in main()")
			break LOOP
		}
	}

	cancel()
	wg.Wait()
}
