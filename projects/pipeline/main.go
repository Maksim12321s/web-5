package main

import (
	"fmt"
	"time"
)

func removeDuplicates(in, out chan string) {
	defer close(out)
	var temp string
	for val := range in {
		if (temp) != val {
			out <- val

		}
		temp = val
		time.Sleep(time.Millisecond * 20)
	}

}
func main() {
	input := make(chan string)
	output := make(chan string)
	var a string
	fmt.Scan(&a)
	go func() {
		for _, r := range a {
			input <- string(r)
		}
		defer close(input)
	}()
	go removeDuplicates(input, output)
	go func() {
		for val := range output {
			fmt.Print(val)
		}
	}()
	time.Sleep(time.Second)
	fmt.Println()
}
