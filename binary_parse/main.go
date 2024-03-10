package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	binary_to_decimal(1101011)
	decimal_to_binary(39)
	binary_to_octal(1101011)
}

func binary_to_decimal(binary int) {
	string_value := fmt.Sprint(binary)
	count := 6
	result := 0.0
	for _, item := range string_value {
		num, _ := strconv.Atoi(string(item))
		result += float64(num) * (math.Pow(2.0, float64(count)))
		count -= 1
	}
	fmt.Println("The devcimal value is:", result)
}

func decimal_to_binary(num float64) {
	final_val := []float64{}
	for math.Floor(num) > 0 {
		final_val = append(final_val, math.Floor(math.Mod(num, 2)))
		num /= 2
	}

	for i, j := 0, len(final_val)-1; i < j; i, j = i+1, j-1 {
		final_val[i], final_val[j] = final_val[j], final_val[i]
	}
	for _, item := range final_val {
		fmt.Println(item)
	}
}

func binary_to_octal(binary int) {
	type Section struct {
		s1 int
		s2 int
		s3 int
	}
	strBinary := strconv.Itoa(binary)
	var binary_sections []Section

	new_section := Section{0, 0, 0}
	i := 0
	for _, item := range strBinary {
		item, err := strconv.Atoi(string(item))
		if err != nil {
			fmt.Println("Error converting:", err)
		}

		switch i {
		case 0:
			new_section.s1 = item
			i += 1
		case 1:
			new_section.s2 = item
			i += 1
		case 2:
			new_section.s3 = item
			binary_sections = append(binary_sections, new_section)
			new_section = Section{0, 0, 0}
			i = 0
		default:
			fmt.Println("Out of bounds")
		}
		for _, item := range binary_sections {
			fmt.Println("Item:", item)
		}
	}
}
