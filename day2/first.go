package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func First() {

	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	valid := 0

	for scanner.Scan() {

		re, _ := regexp.Compile("blue|red|green|[0-9]+")

		line := scanner.Text()

		match := re.FindAllString(line, -1)
		num, _ := strconv.Atoi(match[0])
		valid += num

		for i := range match {
			if match[i] == "blue" {
				blue, _ := strconv.Atoi(match[i-1])
				if blue > 14 {
					valid -= num
					break
				}

			}
			if match[i] == "red" {
				red, _ := strconv.Atoi(match[i-1])
				if red > 12 {
					valid -= num
					break
				}
			}
			if match[i] == "green" {
				green, _ := strconv.Atoi(match[i-1])
				if green > 13 {
					valid -= num
					break
				}
			}
		}

	}
	fmt.Println(valid)

}
