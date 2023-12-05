package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// func isSymbol(s string) bool {
// 	return s == "*" || s == "+" || s == "$" || s == "#"
// }

func First() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines [][]string

	for scanner.Scan() {
		line := scanner.Text()
		strSlice := strings.Split(line, " ")
		for _, str := range strSlice {
			lines = append(lines, []string{str})
		}

	}

	fmt.Println(lines)
	fmt.Println(lines[0][5])

}
