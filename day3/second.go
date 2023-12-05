package day3

import (
	"bufio"
	"log"
	"os"
)

func Second() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

	}

}
