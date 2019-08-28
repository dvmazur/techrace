package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	number, _ := reader.ReadString('\n')

	num, err := strconv.Atoi(strings.TrimSpace(number))
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Printf("The number of divisors of %d is %d\n", num, numberOfDivisors(num))
}

func numberOfDivisors(n int) int {
	pfs := primeFactorization(n)

	num := 1
	for _, exponents := range pfs {
		num *= (exponents + 1)
	}

	return num
}

func primeFactorization(n int) (pfs map[int]int) {
	pfs = make(map[int]int)
	for n%2 == 0 {
		if _, ok := pfs[2]; ok {
			pfs[2] += 1
		} else {
			pfs[2] = 1
		}
		n = n / 2
	}
	for i := 3; i*i <= n; i = i + 2 {
		for n%i == 0 {
			if _, ok := pfs[i]; ok {
				pfs[i] += 1
			} else {
				pfs[i] = 1
			}
			n = n / i
		}
	}
	if n > 2 {
		pfs[n] = 1
	}
	return
}
