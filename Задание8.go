package main

import "fmt"


func Worker(ID int, input <-chan int) {
	for reqID := range input {
		fmt.Printf("балансировщик направил запрос %d\n", reqID)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go Worker(1, ch1)
	go Worker(2, ch2)
	go Worker(3, ch3)
	workerP := []chan int{ch1, ch2, ch3}
	for i := 1; i <= 9; i++ {
		Workert := workerP[(i-1)%len(workerP)]
		Workert <- i
	}
	close(ch1)
	close(ch2)
	close(ch3)

	fmt.Println("все запросы распределены")
}