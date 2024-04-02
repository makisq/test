package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanToArabic = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

var arabicToRoman = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

var arabicValues = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func toRoman(num int) string {
	var result strings.Builder
	for _, value := range arabicValues {
		for num >= value {
			result.WriteString(arabicToRoman[value])
			num -= value
		}
	}
	return result.String()
}

func fromRoman(roman string) (int, error) {
	result := 0
	lastDigit := 0
	for i := len(roman) - 1; i >= 0; i-- {
		digit, ok := romanToArabic[rune(roman[i])]
		if !ok {
			return 0, fmt.Errorf("invalid roman numeral")
		}
		if digit < lastDigit {
			result -= digit
		} else {
			result += digit
		}
		lastDigit = digit
	}
	return result, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 2 + 2 or II + II): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	parts := strings.Fields(input)
	if len(parts) != 3 {
		fmt.Println("Invalid expression format. Please enter the expression in the format 'a + b'.")
		os.Exit(1)
	}

	num1Str, operator, num2Str := parts[0], parts[1], parts[2]
	num1, num1err := strconv.Atoi(num1Str)
	num2, num2err := strconv.Atoi(num2Str)

	if num1err != nil || num2err != nil {
		num1, num1err = fromRoman(num1Str)
		num2, num2err = fromRoman(num2Str)
		if num1err != nil || num2err != nil {
			fmt.Println("Invalid number format. Please enter either arabic or roman numerals.")
			os.Exit(1)
		}
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			os.Exit(1)
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator:", operator)
		os.Exit(1)
	}

	if num1err == nil && num2err == nil {
		fmt.Println(result)
	} else {
		if result <= 0 {
			fmt.Println("Roman numbers must be positive.")
			os.Exit(1)
		}
		fmt.Println(toRoman(result))
	}
}
