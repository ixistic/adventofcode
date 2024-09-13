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

	for {
		maxCubes := map[string]int{
			"red":   -1,
			"green": -1,
			"blue":  -1,
		}
		data, _, err := reader.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		re := regexp.MustCompile("[0-9]+")
		ids := strings.Split(string(data), ":")
		sets := strings.Split(ids[1], ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				numbers := re.FindAllString(cube, -1)
				number, _ := strconv.Atoi(numbers[0])
				for k, v := range maxCubes {
					if strings.Contains(cube, k) && number > v {
						maxCubes[k] = number
					}
				}
			}
		}
		if maxCubes["red"] != -1 && maxCubes["green"] != -1 && maxCubes["blue"] != -1 {
			fmt.Println(maxCubes["red"], maxCubes["green"], maxCubes["blue"], maxCubes["red"]*maxCubes["green"]*maxCubes["blue"])
			sum += maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
		}
	}
	fmt.Println(sum)
}
