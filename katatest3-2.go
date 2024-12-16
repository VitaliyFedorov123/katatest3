package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func romanToArabic(roman string) (int, error) {
	if roman == "I" {
		return 1, nil
	} else if roman == "II" {
		return 2, nil
	} else if roman == "III" {
		return 3, nil
	} else if roman == "IV" {
		return 4, nil
	} else if roman == "V" {
		return 5, nil
	} else if roman == "VI" {
		return 6, nil
	} else if roman == "VII" {
		return 7, nil
	} else if roman == "VIII" {
		return 8, nil
	} else if roman == "IX" {
		return 9, nil
	} else if roman == "X" {
		return 10, nil
	}
	return 0, errors.New("некорректное римское число")
}

func arabicToRoman(number int) (string, error) {
	if number == 1 {
		return "I", nil
	} else if number == 2 {
		return "II", nil
	} else if number == 3 {
		return "III", nil
	} else if number == 4 {
		return "IV", nil
	} else if number == 5 {
		return "V", nil
	} else if number == 6 {
		return "VI", nil
	} else if number == 7 {
		return "VII", nil
	} else if number == 8 {
		return "VIII", nil
	} else if number == 9 {
		return "IX", nil
	} else if number == 10 {
		return "X", nil
	} else if number == 100 {
		return "C", nil
	}
	return "", errors.New("Выдача паники, так как в римской системе нет отрицательных чисел.")
}

func calculate(num1 int, num2 int, operator string) (int, error) {
	if operator == "+" {
		return num1 + num2, nil
	} else if operator == "-" {
		return num1 - num2, nil
	} else if operator == "*" {
		return num1 * num2, nil
	} else if operator == "/" {
		if num2 == 0 {
			return 0, errors.New("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		}
		return num1 / num2, nil
	}
	return 0, errors.New("Выдача паники, так как строка не является математической операцией.")
}

func main() {
	fmt.Println("Input")

	var input string
	fmt.Scanln(&input)

	if input == "exit" {
		return
	}

	var operator string
	if strings.Contains(input, "+") {
		operator = "+"
	} else if strings.Contains(input, "-") {
		operator = "-"
	} else if strings.Contains(input, "*") {
		operator = "*"
	} else if strings.Contains(input, "/") {
		operator = "/"
	} else {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		return
	}

	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		fmt.Println("Выдача паники, так как строка не является математической операцией.")
		return
	}

	num1Str := strings.TrimSpace(parts[0])
	num2Str := strings.TrimSpace(parts[1])

	var num1, num2 int
	var err error
	isRoman1 := false
	isRoman2 := false

	if strings.ContainsAny(num1Str, "IVXLCDM") {
		num1, err = romanToArabic(num1Str)
		if err != nil {
			fmt.Println("Выдача паники, так как в римской системе нет отрицательных чисел:", err)
			return
		}
		isRoman1 = true
	} else {
		num1, err = strconv.Atoi(num1Str)
		if err != nil || num1 < 1 || num1 > 10 {
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			return
		}
	}

	if strings.ContainsAny(num2Str, "IVXLCDM") {
		num2, err = romanToArabic(num2Str)
		if err != nil {
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *):", err)
			return
		}
		isRoman2 = true
	} else {
		num2, err = strconv.Atoi(num2Str)
		if err != nil || num2 < 1 || num2 > 10 {
			fmt.Println("ОВыдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			return
		}
	}

	if isRoman1 != isRoman2 {
		fmt.Println("Выдача паники, так как используются одновременно разные системы счисления.")
		return
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).", err)
		return
	}

	if isRoman1 {
		romanResult, err := arabicToRoman(result)
		if err != nil {
			fmt.Println("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).", err)
			return
		}
		fmt.Println("Output:", romanResult)
	} else {
		fmt.Println("Output:", result)
	}
}
