package main

import (
	"fmt"
	"log"
	"net/http"
)

type Config struct {
	Url           string
	CountRequests int
	Concurrency   int
}

func main() {
	config := Config{"https://openinnovations.ru/api/map", 100, 4}

	channel := make(chan int)

	for i := 1; i <= config.Concurrency; i++ {
		go sendRequest(i, config.CountRequests, config.Url, channel)
	}

	fmt.Printf("Всего запросов: %d", <-channel)
}

func sendRequest(Concurrency int, CountRequests int, Url string, channel chan int) {
	for i := 1; i <= CountRequests; i++ {

		response, _ := http.Get(Url)

		defer response.Body.Close()

		if response.StatusCode != 200 {
			log.Panic("Fuck")
		}
		fmt.Printf("Concurency: %d | Request: %d\n", Concurrency, i)
	}

	data := <-channel
	channel <- data + CountRequests
}
