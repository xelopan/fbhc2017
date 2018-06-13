package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var numCases int

func main() {
	inp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(inp)
	if scanner.Scan() {
		numCases, _ = strconv.Atoi(scanner.Text())
	}
	for day, ans := 1, 0; day <= numCases; day++ {
		if scanner.Scan() {
			numGoods, _ := strconv.Atoi(scanner.Text())
			goods := make([]int, numGoods)
			for i := 0; i < numGoods; i++ {
				if scanner.Scan() {
					goods[i], _ = strconv.Atoi(scanner.Text())
				}
			}
			sort.Ints(goods)
			fmt.Printf("Case #%d: %d %d\n", day, goods, ans)
		}
	}
}

func solveDay(items []int) int {
	return 1
}
