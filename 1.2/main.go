package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input/input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	sum := 0
	words := map[string]string{
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}

		first := ""
		firstIndex := -1
		last := ""
		lastIndex := -1
		for key, value := range words {
			i := strings.Index(string(data), key)

			if firstIndex == -1 && i != -1 {
				firstIndex = i
				first = value
			} else if i < firstIndex && i != -1 {
				firstIndex = i
				first = value
			}

			ii := strings.LastIndex(string(data), key)
			if lastIndex == -1 && i != -1 {
				lastIndex = ii
				last = value
			} else if ii > lastIndex && i != -1 {
				lastIndex = ii
				last = value
			}

			// fmt.Println(key, firstIndex, first, lastIndex, last)
		}

		value, _ := strconv.Atoi(first + last)
		fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)
}
