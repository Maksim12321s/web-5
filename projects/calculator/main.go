package main

import (
	"fmt"
	"time"
)

func calc(firstChan <-chan int, secondChan <-chan int, stopChan <-chan struct{}) <-chan int {

	ch := make(chan int)
	go func(ch chan int) chan int {
		defer close(ch)

		select {
		case x := <-firstChan:
			ch <- x * x
		case x := <-secondChan:
			ch <- x * 3
		case <-stopChan:
			return ch

		}
		return ch

	}(ch)
	return ch

}
func main() {

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan struct{}, 1)
	defer close(ch1)
	defer close(ch2)
	defer close(ch3)
	// var a struct{}
	// ch3 <- a
	var a int
	var num int
	fmt.Println("Введите число")
	fmt.Scan(&a)
	fmt.Println("Введите номер канала")
	fmt.Scan(&num)

	switch {
	case num == 1:
		ch1 <- a
		fmt.Println(<-calc(ch1, ch2, ch3))
	case num == 2:
		ch2 <- a
		fmt.Println(<-calc(ch1, ch2, ch3))
	case num == 3:
		ch3 <- struct{}{}
		fmt.Println(<-calc(ch1, ch2, ch3))
	default:
		fmt.Println("Введите корректный номер канала")
	}

	time.Sleep(time.Second)
}
