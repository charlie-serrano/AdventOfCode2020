package day1

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func ReadFileToArray(filename string) []string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data),"\n")
}

func ProductOfTwoNumbersWhichSumTo(sumTo int, array []string) int {
	for _, iValue := range array {
		for _, jValue := range array {
			iInt, _ := strconv.Atoi(iValue)
			jInt, _ := strconv.Atoi(jValue)

			if iInt + jInt == sumTo {
				return iInt * jInt
			}

		}
	}

	return -1
}

func ProductOfThreeNumbersWhichSumTo(sumTo int, array []string) int {
	for _, iValue := range array {
		for _, jValue := range array {
			for _, kValue := range array {
				iInt, _ := strconv.Atoi(iValue)
				jInt, _ := strconv.Atoi(jValue)
				kInt, _ := strconv.Atoi(kValue)

				if iInt + jInt + kInt == sumTo {
					return iInt * jInt * kInt
				}
			}

		}
	}
	return -1
}
