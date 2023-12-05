package day4

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func First() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		re, _ := regexp.Compile("[0-9]+")
		line := scanner.Text()
		match := re.FindAllString(line, -1)

		occ := make(map[int]int)
		for ig, i := range match {
			if ig == 0 {
				continue
			}
			num, _ := strconv.Atoi(i)
			occ[num]++
		}
		count := 0
		for _, c := range occ {

			if c > 1 {
				if count == 0 {
					count++
					continue
				}
				count *= 2
			}
		}
		sum += count
	}

	println(sum)
}
