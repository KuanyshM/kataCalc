package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	var input string
	fmt.Print("Input: ")
	fmt.Scanln(&input)

	result, err := calculate(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Output:", result)
	}
}

func calculate(input string) (interface{}, error) {

	pattern := regexp.MustCompile(`^(\d+|[IVXLC]+)\s*([-+*/])\s*(\d+|[IVXLC]+)$`)

	if !pattern.MatchString(input) {
		return nil, fmt.Errorf("Ввод не соответствует формату математической операции")
	}

	matches := pattern.FindStringSubmatch(input)
	arg1, isArabic1, err := convertToArabic(matches[1])
	if err != nil {
		return nil, err
	}

	operator := matches[2]

	arg2, isArabic2, err := convertToArabic(matches[3])
	if err != nil {
		return nil, err
	}

	result := 0
	switch operator {
	case "+":
		result = arg1 + arg2
	case "-":
		result = arg1 - arg2
	case "*":
		result = arg1 * arg2
	case "/":
		if arg2 == 0 {
			return nil, fmt.Errorf("Деление на ноль")
		}
		result = arg1 / arg2
	}

	if isArabic1 != isArabic2 {
		return nil, fmt.Errorf("Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}

	if (isArabic1 || isArabic2) && result < 1 {
		return nil, fmt.Errorf("Результат операции с римскими числами не может быть меншье 1")
	}
	if isArabic1 {
		return result, nil

	} else {
		return intToRoman(result), nil
	}
}

func convertToArabic(input string) (int, bool, error) {
	isArabic := regexp.MustCompile(`^\d+$`).MatchString(input)
	result := 0

	if isArabic {
		num, err := strconv.Atoi(input)
		if err != nil {
			return 0, isArabic, fmt.Errorf("Ошибка преобразования в число: %v", err)
		}
		result = num
	} else {

		isRoman := regexp.MustCompile(`^[IVX]+$`).MatchString(input)
		if !isRoman {
			return 0, isArabic, fmt.Errorf("Строка не является числом")
		}

		romanNumerals := map[rune]int{'I': 1, 'V': 5, 'X': 10}
		prevValue := 0

		for _, char := range input {
			value := romanNumerals[char]

			if value > prevValue {
				result += value - 2*prevValue
			} else {
				result += value
			}

			prevValue = value
		}

	}

	if result < 1 || result > 10 {
		return 0, isArabic, fmt.Errorf("Число должно быть от 1 до 10 включительно")
	}
	return result, isArabic, nil

}

func intToRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	rom := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result string

	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			result += rom[i]
			num -= val[i]
		}
	}

	return result
}
