package main

import (
	"context"
	"fmt"
	"time"
)

func main() { 
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 3)
	// ctx, cancel := context.WithTimeout(ctx, time.Second * 3)
	defer cancel()
	bookhotel(ctx)
}

func bookhotel(ctx context.Context) {
	select {
	case <- ctx.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached.")
		return
	case <- time.After(5 * time.Second):
		fmt.Println("Hotel booked")
	}	
}