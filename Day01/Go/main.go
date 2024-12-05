// TODO: Read the contents split and add the differences

// Read from file

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	left  []int
	right []int
)

func Read_file() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Could not open file : %v", err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		l, _ := strconv.Atoi(fields[0])
		r, _ := strconv.Atoi(fields[1])

		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	var sum int = 0

	for i := 0; i < len(left); i++ {
		a := left[i]
		b := right[i]

		diff := int(math.Abs(float64(a - b)))
		sum += diff
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occurred during scanning: %v", err)
	}
}

func main() {
	Read_file()
}
