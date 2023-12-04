package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("./data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum_1 := 0
	sum_2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		a_1, b_1 := parse_word([]rune(line))
		a_2, b_2 := get_numbers([]rune(line))
		fmt.Println(fmt.Sprintf("%v, %v in: %v", a_2, b_2, line))
		sum_1 += a_1 * 10 + b_1
		sum_2 += a_2 * 10 + b_2
	}
	fmt.Println(fmt.Sprintf("Star 1: %v", sum_1))
	fmt.Println(fmt.Sprintf("Star 2: %v", sum_2))
}


func get_numbers(runes []rune) (int, int) {

	a := -1
	b := -1

	string_numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	//inc := Max(Map(string_numbers, func(s string) int { return len(s) }))

	// Find first valid number
	for i := 0; i < len(runes); i++ {
		// Check if any of the runes are digits
		if unicode.IsDigit(runes[i]) && a == -1 {
			a = int(runes[i]) - 48
		}
			

		// If we are here, the first runes are not digits, so might be a number
		if a == -1 {
			for j, s := range string_numbers {
				if len(s) < i {
					continue
				}

				pos := strings.Index(string(runes[:i]), s)

				if pos != -1 {
					a = j + 1
					break
				}
			}
		}

		// Check if any of the runes are digits
		if unicode.IsDigit(runes[len(runes) - i - 1]) && b == -1 {
			b = int(runes[len(runes) - i - 1]) - 48
		}
			

		// If we are here, the first runes are not digits, so might be a number
		if b == -1 {
			for j, s := range string_numbers {
				pos := strings.Index(string(runes[i:]), s)

				if pos != -1 {
					b = j + 1
					break
				}
			}
		}

		if a != -1 && b != -1 {
			return a, b
		}
	}

	if b == -1 {
		return a, a
	}
	return a, b
}

func parse_word(line []rune) (int, int) {
	a := -1
	b := -1
	for i, r := range line {

		l := line[len(line) - i - 1]

		if a == -1 && unicode.IsDigit(r) {
			a = int(r) - 48
		}

		if b == -1 && unicode.IsDigit(l) {
			b = int(l) - 48
		}

		if a != -1 && b != -1 {
			return a, b
		}
	}

	return a, b
}



func Map[A any, B any](slice []A, fn func(A) B) []B {
    result := make([]B, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

func Max(slice []int) int {
    if len(slice) == 0 {
        return 0 // or panic, or return an error depending on your use case
    }
    max := slice[0]
    for _, v := range slice {
        if v > max {
            max = v
        }
    }
    return max
}
