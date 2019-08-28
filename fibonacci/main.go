package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const num = 10

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	number, _ := reader.ReadString('\n')

	num, err := strconv.Atoi(strings.TrimSpace(number))
	if err != nil {
		log.Fatal(err)
		return
	}

	if num < 1 {
		log.Fatalf("n=%d too small\n", num)
		return
	}

	a := 0
	b := 1

	for c := 0; c < num; c++ {
		a, b = b, a+b
	}

	fmt.Println(a)
}
