package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(200 * time.Millisecond)
	for i := 1; i <= 16; i++ {
		<-tick
		go func(id int) {
			fmt.Printf("Запрос %d обработан\n", id)
		}(i)
	}
}