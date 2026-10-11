package main

import (
	"crypto/md5"
	"fmt"
	"sync"
)

func worker(id int, files chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for file := range files {
		hash := md5.Sum([]byte(file))
		fmt.Printf("Файл: %s -> MD5: %x\n", file, hash)
	}
}

func main() {
	filesp := []string{
		"document.txt",
		"image.png",
		"video.mp4",
		"archive.zip",
		"notes.md",
	}
	filesc := make(chan string, len(filesp))
	num := 3
	var wg sync.WaitGroup
	for w := 1; w <= num; w++ {
		wg.Add(1)
		go worker(w, filesc, &wg)
	}
	for _, file := range filesp {
		filesc <- file
	}
	close(filesc)
	wg.Wait()
}