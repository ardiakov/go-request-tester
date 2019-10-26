package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Url           string
	CountRequests int
	Concurrency   int
}

func main() {
	config := Config{"https://openinnovations.ru/api/map", 100, 4}

	counter := 0

	for i := 1; i <= config.Concurrency; i++ {
		go sendRequest(i, config.CountRequests, config.Url, &counter)
	}

	time.Sleep(time.Second * 10)

	fmt.Printf("Всего запросов: %d", counter)
}

func sendRequest(Concurrency int, CountRequests int, Url string, counter *int) {
	for i := 1; i <= CountRequests; i++ {
		*counter++

		response, _ := http.Get(Url)

		defer response.Body.Close()

		if response.StatusCode != 200 {
			log.Panic("Fuck")
		}

		fmt.Printf("Concurency: %d | Request: %d\n", Concurrency, i)
	}
}
