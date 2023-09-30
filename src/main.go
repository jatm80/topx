package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Welcome to topx")

	var topOnes []int

	if len(os.Args) < 3 {
		fmt.Printf("Usage:  \n ./topx <file-path> <X> \n\n")
		log.Fatal("Not enough arguments")
	}

	filepath := os.Args[1]
	number := os.Args[2]

	x, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err.Error())
	}

	i := 0
	for i < x {

		file := openFile(filepath)
		defer file.Close()

		fileScan := bufio.NewScanner(file)
		y := 0

		for fileScan.Scan() {
			line, err := strconv.Atoi(fileScan.Text())
			if err != nil {
				log.Println(err.Error())
			}
			if line > y {
				if !inSlice(topOnes, line) {
					y = line
				}
			}
		}

		topOnes = append(topOnes, y)

		if fileScan.Err() != nil {
			log.Fatal(err.Error())
		}
		i++
	}

	for _, v := range topOnes {
		fmt.Println(v)
	}
}

func inSlice(topOnes []int, line int) bool {
	for _, v := range topOnes {
		if v == line {
			return true
		}
	}
	return false
}

func openFile(filepath string) *os.File {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal("Error: ", err.Error())
	}

	return file

}
