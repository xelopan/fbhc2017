package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fmt.Println("Reading from ", os.Args[1])

	outputFile, err := os.OpenFile("D:/ab/fbhc2017outputs/cals/progresspie.txt", os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outputFile)
	i := 0
	// var output string
	for scanner.Scan() {
		if i == 0 {
			i++
			continue
		}
		inputText := scanner.Text()
		input := strings.Split(inputText, " ")
		out := fmt.Sprintf("Case #%d: ", i)

		p, err := strconv.Atoi(input[0])
		x, err := strconv.Atoi(input[1])
		y, err := strconv.Atoi(input[2])

		if err != nil {
			log.Fatal(err)
		}
		if solve(p, x, y) {
			out += "black"
		} else {
			out += "white"
		}
		// output += out + "\n"
		writer.WriteString(out + "\n")
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Print(output)
	if err := writer.Flush(); err != nil {
		log.Fatal(err)
	}
}

/*
	Check if the point (x, y) is in the progress of p%
*/
func solve(p, x, y int) bool {
	if p == 0 {
		return false
	}
	// calculate distance to the centre
	dx := float64(x - 50)
	dy := float64(y - 50)
	d := math.Sqrt(dx*dx + dy*dy)
	if d > 50 {
		return false
	}
	if dx == 0 && dy == 0 {
		return true
	}
	// calculate the progress percentage of the point (x,y)
	theta := math.Atan2(dy, dx)
	if theta < 0 {
		theta += 2 * math.Pi
	}
	myp := (2*math.Pi - theta) / (2 * math.Pi)
	myp += .25
	if myp >= 1 {
		myp--
	}

	return myp <= (float64(p) / 100)
}
