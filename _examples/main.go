package main

import (
	"fmt"
	"log"
	"time"

	"../fluent"
)

func main() {
	logger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "127.0.0.1"})
	if err != nil {
		fmt.Println(err)
	}
	defer logger.Close()
	tag := "myapp.access"
	var data = map[string]string{
		"foo":  "bar",
		"hoge": "hoge"}
	i := 0
	for i < 100 {
		e := logger.Post(tag, data)
		if e != nil {
			log.Println("Error while posting log: ", e)
		} else {
			log.Println("Success to post log")
		}
		i = i + 1
		time.Sleep(1000 * time.Millisecond)
	}
}
