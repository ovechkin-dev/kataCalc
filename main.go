package main

import (
	"fmt"
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

func main() {
	// parse three numbers mentioned in task description
	r := "V"
	w := "VII"
	v := romanToInt(r)
	x := romanToInt(w)

	fmt.Println(r, "==", v)
	fmt.Println(w, "==", x)
	fmt.Println(intToRoman(add(v, x)))
	fmt.Println(add(v, x))
	fmt.Println(sub(v, x))
	fmt.Println(mul(v, x))
	fmt.Println(div(v, x))

}
