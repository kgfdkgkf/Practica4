package main

import (
	"fmt"
	"time"
)

func search(name string, d time.Duration, ch chan<- string) {
	time.Sleep(d)
	ch <- fmt.Sprintf("результат из: %s", name)
}

func main() {
	resChan := make(chan string, 3)
	go search("медленная БД", 800*time.Millisecond, resChan)
	go search("быстрая БД", 150*time.Millisecond, resChan)
	go search("облачное хранилище", 400*time.Millisecond, resChan)

	firstR := <-resChan
	fmt.Println("самый быстрый ответ:", firstR)
}