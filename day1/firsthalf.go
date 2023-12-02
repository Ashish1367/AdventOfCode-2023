package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func First() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var firstnum, secondnum rune
		for _, i := range line {
			if unicode.IsDigit(i) {
				if firstnum == 0 {
					firstnum = i
				}
				secondnum = i
			}
		}
		firstnumint, _ := strconv.Atoi(string(firstnum))
		secondnumint, _ := strconv.Atoi(string(secondnum))

		result := firstnumint*10 + secondnumint

		sum += result

	}
	fmt.Println(sum)
}
