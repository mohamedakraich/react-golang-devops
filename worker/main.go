package main

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// ----- to subscribe to a channel

	// There is no error because go-redis automatically reconnects on error.
	pubsub := rdb.Subscribe(ctx, "channel")
	// Close the subscription when we are done.
	defer pubsub.Close()

	// ----- to receive a message
	ch := pubsub.Channel()

	for msg := range ch {
		// fmt.Println(msg.Channel, msg.Payload)

		fibIndex, err := strconv.Atoi(msg.Payload)
		if err != nil {
			panic(err)
		}

		rdb.HSet(ctx, "values", msg.Payload, fib(fibIndex))
	}
}

// the fibonacci function
func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
