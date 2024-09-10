package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		re := regexp.MustCompile("[0-9]")
		numbers := re.FindAllString(string(data), -1)
		value := 0
		if len(numbers) == 1 {
			value, _ = strconv.Atoi(numbers[0] + numbers[0])
		} else {
			value, _ = strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		}
		fmt.Println(value)
		sum += value
	}
	fmt.Println(sum)
}
