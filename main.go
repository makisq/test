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

func isRomanNumeral(input string) bool {
	for _, r := range input {
		if _, ok := romanToArabic[r]; !ok {
			return false
		}
	}
	return true
}

func isOperator(token string) bool {
	operators := []string{"+", "-", "*", "/"}
	for _, op := range operators {
		if token == op {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter the expression (e.g., 2 + 2 or II + II): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Fields(input)
		if len(parts) != 3 {
			panic("Invalid expression format. Please enter the expression in the format 'a + b'.")
		}

		num1Str, operator, num2Str := parts[0], parts[1], parts[2]
		num1IsRoman, num2IsRoman := isRomanNumeral(num1Str), isRomanNumeral(num2Str)

		if !isOperator(operator) {
			panic("Invalid operator: " + operator)
		}

		if (num1IsRoman && !num2IsRoman) || (!num1IsRoman && num2IsRoman) {
			panic("Both numbers must be in the same numeral system.")
		}

		var num1, num2 int
		var err error

		if num1IsRoman {
			num1, err = fromRoman(num1Str)
			if err != nil {
				panic(err)
			}
			num2, err = fromRoman(num2Str)
			if err != nil {
				panic(err)
			}
			if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
				panic("Roman numbers must be in the range from I to X (1 to 10).")
			}
		} else {
			num1, err = strconv.Atoi(num1Str)
			if err != nil {
				panic(err)
			}
			num2, err = strconv.Atoi(num2Str)
			if err != nil {
				panic(err)
			}
			if num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
				panic("Arabic numbers must be in the range from 1 to 10.")
			}
		}

		var result int
		switch operator {
		case "+":
			result = num1 + num2
		case "-":
			if num1 < num2 {
				panic("Roman numbers must be positive.")
			}
			result = num1 - num2
		case "*":
			result = num1 * num2
		case "/":
			if num2 == 0 {
				panic("Error: Division by zero is not allowed.")
			}
			result = num1 / num2
		default:
			panic("Invalid operator: " + operator)
		}

		if num1IsRoman && result <= 0 {
			panic("Roman numbers must be positive.")
		}

		if num1IsRoman {
			fmt.Println(toRoman(result))
		} else {
			fmt.Println(result)
		}
	}
}
