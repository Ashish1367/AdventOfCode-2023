package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Second() {

	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {

		re, _ := regexp.Compile("blue|red|green|[0-9]+")

		line := scanner.Text()

		match := re.FindAllString(line, -1)

		red, blue, green := 0, 0, 0

		for i := range match {
			if match[i] == "blue" {
				num, _ := strconv.Atoi(match[i-1])
				if num > blue {
					blue = num
				}

			}
			if match[i] == "red" {
				num, _ := strconv.Atoi(match[i-1])
				if num > red {
					red = num
				}
			}
			if match[i] == "green" {
				num, _ := strconv.Atoi(match[i-1])
				if num > green {
					green = num
				}
			}
		}
		total += red * green * blue
	}
	fmt.Println(total)

}
