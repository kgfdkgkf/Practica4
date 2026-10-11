package main
import "fmt"

func worker(i int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Обработка задачи", i, j)
		results <- j * j
		fmt.Println("Завершение обработки задачи", i, j)
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for x := 1; x <= 3; x++ {
		go worker(x, jobs, results)
	}

	for j := 1; j<=10; j++ {
		jobs <- j
	}
	close(jobs)

	for y := 1; y <= 10; y++ {
		fmt.Println("результат: ", <-results)
	}
}