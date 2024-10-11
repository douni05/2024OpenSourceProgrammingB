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
	var realscore float64
	fmt.Print("점수 입력 : ")
	in := bufio.NewReader(os.Stdin)
	score, err := in.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	score = strings.TrimSpace(score)             // 줄바꿈, 띄어쓰기, 탭 줄 제거 (python, strip 과 유사)
	realscore, _ = strconv.ParseFloat(score, 64) // 실수형 64비트 다입으로 형변환
	if realscore >= 90 {
		fmt.Println("A")
	} else {
		fmt.Println("BCDE")
	}
}
