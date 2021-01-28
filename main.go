package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/kuka0len/go-cache/cacher"
)

func main() {
	messageCacher := cacher.New("message", 0)
	message := "Lying under the giant tree!"
	id := "666"

	val, err := messageCacher.Set(id, message)
	if err != redis.Nil {
		fmt.Println(val)
	} else {
		fmt.Println(err)
	}

	val, err = messageCacher.Get(id)
	if err != redis.Nil {
		fmt.Println("Message is: " + val)
	} else {
		fmt.Println("No message with the id " + id)
	}

	id = "667"
	val, err = messageCacher.Get(id)
	if err != redis.Nil {
		fmt.Println("Message is: " + val)
	} else {
		fmt.Println("No message with the id " + id)
	}
}
