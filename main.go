package main

import (
	"bufio"
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

func calculation(strSplit []string, simbol string) (calc int, calcRoman string) {
	roman := false
	a, err := strconv.Atoi(strSplit[0])
	if err != nil {
		a = romanToInt(strSplit[0])
		roman = true
	}
	b, err := strconv.Atoi(strSplit[1])
	if err != nil {
		b = romanToInt(strSplit[1])
		roman = true
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
		calcRoman = intToRoman(calc)
	} else {
		calcRoman = ""
	}
	return calc, calcRoman
}

func calculationString(text string) (result int, resultRoman string) {

	if strings.Contains(text, "+") {
		result, resultRoman = calculation(strings.Split(text, "+"), "+")
	}
	if strings.Contains(text, "-") {
		result, resultRoman = calculation(strings.Split(text, "-"), "-")
	}
	if strings.Contains(text, "*") {
		result, resultRoman = calculation(strings.Split(text, "*"), "*")
	}
	if strings.Contains(text, "/") {
		result, resultRoman = calculation(strings.Split(text, "/"), "/")
	}
	return result, resultRoman
}

func main() {

	fmt.Println("Введите задачу")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка: ", err)
	}
	text = strings.TrimSpace(text)
	//fmt.Println(text)

	resultInt, resultRom := calculationString(text)

	if resultRom == "" {
		fmt.Println(resultInt)
	} else {
		fmt.Println(resultRom)
	}

}
