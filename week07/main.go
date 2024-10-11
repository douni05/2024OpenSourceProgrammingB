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
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	score = strings.TrimSpace(score)                // 줄바꿈, 띄어쓰기, 탭 줄 제거 (python, strip 과 유사)
	realscore, _ := strconv.ParseInt(score, 16, 32) // 정수형 32비트 다입으로 형변환
	if realscore >= 60 {
		fmt.Println("A")
		fmt.Printf("%d\n", realscore)
	} else {
		fmt.Println("BCDE")
		fmt.Printf("%d\n", realscore)
	}
}
