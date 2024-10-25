package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answer := rand.Intn(6) + 1 // dice 1 ~ 6
	fmt.Println(answer)

	fmt.Print("숫자 입력 : ")
	in := bufio.NewReader(os.Stdin)
	input, err := in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input) // 줄바꿈, 띄어쓰기, 탭 줄 제거 (python, strip 과 유사)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(guess)
}
