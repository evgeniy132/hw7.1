package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateNumbers(ch chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		number := rand.Intn(100)
		ch <- number
	}
	close(ch)
}

func calculateAverage(input <-chan int, result chan<- float64) {
	sum := 0
	count := 0
	for number := range input {
		sum += number
		count++
	}
	average := float64(sum) / float64(count)
	result <- average
	close(result)
}

func printAverage(result <-chan float64) {
	average := <-result
	fmt.Println("Average", average)
}

func main() {
	numbersChannel := make(chan int)
	averageChannel := make(chan float64)

	go generateNumbers(numbersChannel)
	go calculateAverage(numbersChannel, averageChannel)
	go printAverage(averageChannel)

	time.Sleep(1 * time.Second)

}
