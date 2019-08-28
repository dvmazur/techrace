package main

import (
	"fmt"
)

func main() {
	for message := range techRace() {
		fmt.Println(message)
	}
}

func techRace() <-chan string {
	out := make(chan string)
	go func() {
		for i := 1; true; i++ {
			result := ""
			if i%3 == 0 {
				result += "Tech"
			}
			if i%5 == 0 {
				result += "Race"
			}
			if result == "" {
				result = fmt.Sprintf("%v", i)
			}
			out <- result
		}
		close(out)
	}()

	return out
}
