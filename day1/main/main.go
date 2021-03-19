package main

import (
	"AdventOfCode2020/day1"
	"fmt"
)

func main() {
	data := day1.ReadFileToArray("input.txt")

	result1 := day1.ProductOfTwoNumbersWhichSumTo(2020, data)
	result2 := day1.ProductOfThreeNumbersWhichSumTo(2020, data)

	fmt.Printf("Part1: %d\n", result1)
	fmt.Printf("Part1: %d\n", result2)





}
