package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func romanToArabic(roman string) (int, error) {
	romanMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}
	var result, prevValue int
	for i := len(roman) - 1; i >= 0; i-- {
		char := rune(roman[i])
		value, exists := romanMap[char]
		if !exists {
			return 0, errors.New("некорректное римское число")
		}
		if value < prevValue {
			result -= value
		} else {
			result += value
		}
		prevValue = value
	}
	if result < 1 || result > 10 {
		return 0, errors.New("римские числа должны быть в диапазоне от 1 до 10")
	}
	return result, nil
}

func arabicToRoman(arabic int) (string, error) {
	if arabic < 1 || arabic > 10 {
		return "", errors.New("арабские числа должны быть в диапазоне от 1 до 10")
	}
	values := []int{10, 9, 5, 4, 1}
	symbols := []string{"X", "IX", "V", "IV", "I"}
	var result strings.Builder
	for i := 0; i < len(values); i++ {
		for arabic >= values[i] {
			result.WriteString(symbols[i])
			arabic -= values[i]
		}
	}
	return result.String(), nil
}

func calculate(a, b int, operator string) (int, error) {
	switch operator {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("деление на ноль невозможно")
		}
		return a / b, nil
	default:
		return 0, errors.New("некорректный оператор")
	}
}

func main() {
	fmt.Println("Введите выражение (числа от 1 до 10)):")

	var input string
	fmt.Scanln(&input)
	input = strings.ReplaceAll(input, " ", "")

	var num1Str, num2Str, operator string
	for _, op := range "+-*/" {
		if parts := strings.Split(input, string(op)); len(parts) == 2 {
			num1Str, num2Str = parts[0], parts[1]
			operator = string(op)
			break
		}
	}

	num1, err1 := strconv.Atoi(num1Str)
	num2, err2 := strconv.Atoi(num2Str)
	isArabic := err1 == nil && err2 == nil

	if !isArabic {
		num1, err1 = romanToArabic(num1Str)
		num2, err2 = romanToArabic(num2Str)
	}

	if err1 != nil || err2 != nil || num1 < 1 || num1 > 10 || num2 < 1 || num2 > 10 {
		fmt.Println("Ошибка: числа должны быть от 1 до 10")
		return
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if isArabic {
		fmt.Println("Результат:", result)
	} else {
		romanResult, _ := arabicToRoman(result)
		fmt.Println("Результат:", romanResult)
	}
}
