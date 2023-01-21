package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	result := 0
	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func intToRoman(s int) (roman string) {
	numbers := []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	romans := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	index := len(romans) - 1

	for s > 0 {
		for numbers[index] <= s {
			roman += romans[index]
			s -= numbers[index]
		}
		index--
	}

	return roman
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func calculation(strSplit []string, simbol string) (calc int, calcRoman string, err error) {
	roman := false

	if len(strSplit) > 2 {
		return len(strSplit), "", errors.New("Максимум одна операция")
	}

	a, err1 := strconv.Atoi(strSplit[0])
	if err1 != nil {
		a = romanToInt(strSplit[0])
		roman = true
	}
	b, err2 := strconv.Atoi(strSplit[1])
	if err2 != nil {
		b = romanToInt(strSplit[1])
		roman = true
	}

	if (err1 != nil && err2 == nil) || (err1 == nil && err2 != nil) {
		return -1, "", errors.New("Некорректное выражение")
	}

	if a > 10 || b > 10 {
		return -1, "", errors.New("Число больше 10")
	}

	if a < 1 || b < 1 {
		return -1, "", errors.New("Число меньше 1")
	}

	switch simbol {
	case "+":
		calc = add(a, b)
	case "-":
		calc = sub(a, b)
	case "*":
		calc = mul(a, b)
	case "/":
		calc = div(a, b)
	}
	if roman {
		if calc < 1 {
			return -1, "", errors.New("Римские числа. Результат меньше еденицы.")
		} else {
			calcRoman = intToRoman(calc)
		}
	} else {
		calcRoman = ""
	}
	return calc, calcRoman, err
}

func calculationString(text string) (result int, resultRoman string, err error) {

	if strings.Contains(text, "+") {
		result, resultRoman, err = calculation(strings.Split(text, "+"), "+")
	}
	if strings.Contains(text, "-") {
		result, resultRoman, err = calculation(strings.Split(text, "-"), "-")
	}
	if strings.Contains(text, "*") {
		result, resultRoman, err = calculation(strings.Split(text, "*"), "*")
	}
	if strings.Contains(text, "/") {
		result, resultRoman, err = calculation(strings.Split(text, "/"), "/")
	}
	return result, resultRoman, err
}

func main() {

	fmt.Println("Введите задачу")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")
	//	fmt.Println(text)

	resultInt, resultRom, err := calculationString(text)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		if resultRom == "" {
			fmt.Println(resultInt)
		} else {
			fmt.Println(resultRom)
		}
	}

}
