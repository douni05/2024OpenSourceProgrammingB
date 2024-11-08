package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// input
// 10
// 19

// output
// 11 13 17 19
func isPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	} else {
		i := 3
		for i*i <= n {
			if n%i == 0 {
				return false
			}
			//fmt.Printf("%d ", i)
			i = i + 2
		}
	}
	return true
}

func main() {
	fmt.Print("첫 번쨰 정수(시작 값) 입력 : ")
	in := bufio.NewReader(os.Stdin)
	number, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	number = strings.TrimSpace(number)
	n1, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("두 번쨰 정수(시작 값) 입력 : ")
	//in := bufio.NewReader(os.Stdin)
	b, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	b = strings.TrimSpace(b)
	n2, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}

	for i := n1; i <= n2; i++ {
		if isPrime(i) {
			fmt.Printf("%d", i)
		}
	}
}
