package main

import (
	"fmt"
	"strconv"
)

var Numerals = [100]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}

func main() {
	var a, operator, b, c string
	fmt.Print("Введите данные: ")
	fmt.Scanln(&a, &operator, &b, &c)

	if len(c) > 0 {
		fmt.Println("Ошибка, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		return
	}

	if len(a) == 0 || len(b) == 0 || len(operator) == 0 {
		fmt.Println("Ошибка, так как строка не является математической операцией")
		return
	}
	searchRomenumerals(a, operator, b)
}

func searchRomenumerals(a, operator, b string) {
	var num1, num2 int

	for i, n := range Numerals {
		if a == n {
			num1 = i + 1
		}
		if b == n {
			num2 = i + 1
		}

		if num1 > 10 || num2 > 10 {
			fmt.Println("Калькулятор принимает на вход только числа от 1 до 10 включительно")
			return
		}

	}

	if num1 == 0 || num2 == 0 {
		arabCalc(a, operator, b)
		return
	}

	if num1 < num2 {
		fmt.Println("Ошибка, так как в римской системе нет отрицательных чисел")
		return
	}

	romeCalc(num1, operator, num2)
	return
}

func romeCalc(num1 int, operator string, num2 int) {
	if operator == "+" {
		fmt.Println(Numerals[num1+num2-1])
	}

	if operator == "-" {
		fmt.Println(Numerals[num1-num2-1])
	}

	if operator == "*" {
		fmt.Println(Numerals[num1*num2-1])
	}

	if operator == "/" {
		fmt.Println(Numerals[num1/num2-1])
	}
}

func arabCalc(a, operator, b string) {
	num1, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println("Ошибка, так как используются одновременно разные системы счисления.")
		return
	}

	num2, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("Ошибка, так как используются одновременно разные системы счисления.")
		return
	}

	if num1 > 10 || num2 > 10 {
		fmt.Println("Калькулятор принимает на вход только числа от 1 до 10 включительно")
		return
	}

	if operator == "+" {
		fmt.Println(num1 + num2)
		return
	}

	if operator == "-" {
		fmt.Println(num1 - num2)
		return
	}

	if operator == "*" {
		fmt.Println(num1 * num2)
		return
	}

	if operator == "/" {
		fmt.Println(num1 / num2)
		return
	} else {
		fmt.Println("Калькулятор принимает на вход только операторы + - * /")
	}
}

// 	fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
// }
