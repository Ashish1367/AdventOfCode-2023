package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func Second() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	count := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		var fulldigit string
		for j := range line {
			if unicode.IsDigit(rune(line[j])) {
				fulldigit += string(line[j])

			}
			for k := j + 1; k <= len(line); k++ {
				sub := line[j:k]
				if val, found := count[sub]; found {
					fulldigit += val

					break
				}
			}
		}
		num := fmt.Sprintf("%s%s", string(fulldigit[0]), string(fulldigit[len(fulldigit)-1]))

		result, _ := strconv.Atoi(num)

		sum += result
	}
	fmt.Println(sum)

}
