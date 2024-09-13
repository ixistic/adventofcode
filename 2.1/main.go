package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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

	maxCubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	for {
		data, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		re := regexp.MustCompile("[0-9]+")
		ids := strings.Split(string(data), ":")
		fmt.Println(ids)
		sets := strings.Split(ids[1], ";")
		correct := true
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				numbers := re.FindAllString(cube, -1)
				number, _ := strconv.Atoi(numbers[0])
				for k, v := range maxCubes {
					if strings.Contains(cube, k) && number > v {
						correct = false
					}
				}
				if !correct {
					break
				}
			}
			if !correct {
				break
			}
		}
		if correct {
			idValues := re.FindAllString(ids[0], -1)
			idValue, _ := strconv.Atoi(idValues[0])
			fmt.Println(idValue)
			sum += idValue
		}
	}
	fmt.Println(sum)
}
