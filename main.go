package main

import (
	"fmt"
)

var aplhabet = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func input() string {
	var word string
	fmt.Print("Введите римское число: ")
	fmt.Scan(&word)
	return word
}

func checkInput(word string) bool {
	if word == "" {
		return false
	}

	for _, sym := range word {
		if _, ok := aplhabet[sym]; !ok {
			return false
		}
	}

	return true
}

func checkRepeat(runes []rune) bool {
	for i := 0; i < len(runes); i++ {
		current := runes[i]

		if current == 'I' || current == 'X' || current == 'C' {
			count := 1
			for j := i + 1; j < len(runes) && runes[j] == current; j++ {
				count++
			}
			if count > 3 {
				return false
			}
		}

		if current == 'V' || current == 'L' || current == 'D' {
			if i+1 < len(runes) && runes[i+1] == current {
				return false
			}
		}
	}
	return true
}

func checkSubtraction(runes []rune) bool {
	for i := 0; i < len(runes)-1; i++ {
		current := runes[i]
		next := runes[i+1]

		if sravnenie(current, next) == -1 {
			isValid := false
			switch current {
			case 'I':
				if next == 'V' || next == 'X' {
					isValid = true
				}
			case 'X':
				if next == 'L' || next == 'C' {
					isValid = true
				}
			case 'C':
				if next == 'D' || next == 'M' {
					isValid = true
				}
			}
			if !isValid {
				return false
			}

			if i > 0 && runes[i-1] == current {
				return false
			}

			if i+2 < len(runes) {
				next2 := runes[i+2]
				if sravnenie(current, next2) != 1 {
					return false
				}
			}
		}
	}
	return true
}
func checkValidRoman(word string) bool {
	runes := []rune(word)

	if !checkRepeat(runes) {
		return false
	}

	if !checkSubtraction(runes) {
		return false
	}

	return true
}

func sravnenie(sym1, sym2 rune) int {
	num1 := aplhabet[sym1]
	num2 := aplhabet[sym2]

	switch {
	case num1 > num2:
		return 1
	case num1 < num2:
		return -1
	default:
		return 0
	}
}

func convert(word string) int {
	var res int
	runes := []rune(word)

	for i, sym := range runes {
		if i == len(runes)-1 {
			res += aplhabet[sym]
		} else {
			next := runes[i+1]
			cmp := sravnenie(sym, next)

			if cmp == -1 {
				res -= aplhabet[sym]
			} else {
				res += aplhabet[sym]
			}
		}
	}

	return res
}

func main() {
	word := input()

	if !checkInput(word) {
		fmt.Println("Ошибка: введены недопустимые символы")
		return
	}

	if !checkValidRoman(word) {
		fmt.Println("Ошибка: неверное написание римского числа")
		return
	}

	result := convert(word)
	fmt.Print("Результат: ", result)
}
