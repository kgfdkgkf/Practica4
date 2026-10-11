package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"https://google.com",
		"https://github.com",
		"https://ya.ru",
	}

	semofor := make(chan struct{}, 2)
	var w sync.WaitGroup
	for _, url := range urls{
		w.Add(1)
		semofor <- struct{}{}
		go func(u string){
			defer w.Done()
			r, err := http.Get(u)
			if err != nil{
				fmt.Println("Ошибка ", err)
				return
				
			}
			defer r.Body.Close()
			fmt.Println("url: %s -> статус: %s\n", u, r.Status)
		}(url)
	}
	w.Wait()
}