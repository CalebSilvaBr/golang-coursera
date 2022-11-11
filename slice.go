package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	const alpha = "abcdefghijklmnopqrstuvwyz"
	intList := make([]int, 0, 3)

	var scanValue string
	for {
		fmt.Print("Enter an integer", ": ")
		_, err := fmt.Scan(&scanValue)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		exitCode := strings.ToLower(scanValue)

		if strings.ContainsAny(alpha, exitCode) {
			fmt.Println("Please, enter an integer!")
		} else if exitCode == "x" {
			fmt.Println("Exiting...")
			os.Exit(0)
		} else {
			parseIntScanValue, _ := strconv.Atoi(scanValue)
			intList = append(intList, parseIntScanValue)
			sort.Ints(intList)
			fmt.Println(intList)
		}
	}
}
