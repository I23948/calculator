package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция дял вывода результата арифметического действия
func check(text []string) int {
	a := string(text[0])
	a1, _ := strconv.Atoi(a)

	b := strings.TrimSpace((text[2]))
	b1, _ := strconv.Atoi(b)

	if text[1] == ("+") {
		return (a1 + b1)
	} else if text[1] == ("-") {
		return (a1 - b1)
	} else if text[1] == ("*") {
		return (a1 * b1)
	} else if text[1] == ("/") {
		return (a1 / b1)
	} else {
		return 0
	}
}
func main() {
	fmt.Println("Введите выражение")
	reader := bufio.NewReader(os.Stdin)
	text1, _ := reader.ReadString('\n')
	text := strings.Split(text1, " ")
	//Определение первого числа
	FirstNumber0 := text[0]
	FirstNumber := strings.Split(FirstNumber0, "")
	//Определение второго числа
	LastNumber0 := text[2]
	LastNumber := strings.Split(LastNumber0, "")
	// Проверка на тип данных
	if FirstNumber[0] == "1" || FirstNumber[0] == "2" || FirstNumber[0] == "3" || FirstNumber[0] == "4" || FirstNumber[0] == "5" || FirstNumber[0] == "6" || FirstNumber[0] == "7" || FirstNumber[0] == "8" || FirstNumber[0] == "9" || FirstNumber[0] == "10" {
		fmt.Println(check(text))
	} else if FirstNumber[0] == "I" || FirstNumber[0] == "V" || FirstNumber[0] == "X" {
		//Определение Первого римского числа
		FirstRome := []int{}
		for _, i := range FirstNumber {
			j, _ := strconv.Atoi(i)
			FirstRome = append(FirstRome, j)
		}
		//Определение Второго римского числа
		SecondRome := []int{}
		for _, i := range LastNumber {
			j, _ := strconv.Atoi(i)
			SecondRome = append(SecondRome, j)
		}
		//  Расшифровка Первого РИмского Числа
		for i := 0; i < len(FirstNumber); i++ {
			if FirstNumber[i] == "C" {
				FirstRome[i] = 100
			} else if FirstNumber[i] == "L" {
				FirstRome[i] = 50
			} else if FirstNumber[i] == "X" {
				FirstRome[i] = 10
			} else if FirstNumber[i] == "V" {
				FirstRome[i] = 5
			} else if FirstNumber[i] == "I" {
				FirstRome[i] = 1
			}
		}
		//Перевод первого римского числа в арабское
		var result int
		result = FirstRome[0]
		for i := 1; i < len(FirstRome); i++ {
			if len(FirstRome) == 1 {
				break

			} else if FirstRome[i] > FirstRome[i-1] {
				FirstRome[i] -= FirstRome[i-1]
				FirstRome[i-1] = 0
			}
			result = 0
			if i == len(FirstRome)-1 {
				for i := len(FirstRome) - 1; i >= 0; i-- {
					result = result + FirstRome[i]
				}
			}
		}

		//Расшифровка Второго РИмского числа
		for i := 0; i < len(LastNumber); i++ {
			if LastNumber[i] == "C" {
				SecondRome[i] = 100
			} else if LastNumber[i] == "L" {
				SecondRome[i] = 50
			} else if LastNumber[i] == "X" {
				SecondRome[i] = 10
			} else if LastNumber[i] == "V" {
				SecondRome[i] = 5
			} else if LastNumber[i] == "I" {
				SecondRome[i] = 1
			} else {
			}
		}
		//Перевод второго числа в арабское
		var result2 int = 0
		for i := 1; i < len(SecondRome); i++ {
			if len(SecondRome) == 1 {
				result2 = SecondRome[0]
			} else if SecondRome[i] > SecondRome[i-1] {
				SecondRome[i] -= SecondRome[i-1]
				SecondRome[i-1] = 0
			}
			if i == len(SecondRome)-1 {
				for i := len(SecondRome) - 1; i >= 0; i-- {
					result2 = result2 + SecondRome[i]
				}
			}
		}

		text3 := []string{strconv.Itoa(result), text[1], strconv.Itoa(result2)}

		var res int = check(text3)
		Ator(res)

	}
}

func repeat(symbol string, number int) string {
	var str string
	for i := 0; i < number; i++ {
		str += symbol
	}
	return str
}
func Ator(input int) {
	var c, l, x, v, i string = "C", "L", "X", "V", "I"
	var c1, l1, x1, v1, i1 int
	c1 = input / 100
	input = input % 100
	l1 = input / 50
	input = input % 50
	x1 = input / 10
	input = input % 10
	v1 = input / 5
	input = input % 5
	i1 = input
	var numbers = []int{c1, l1, x1, v1, i1}
	var literals = []string{c, l, x, v, i}
	var result string
	if x1 == 4 && l1 == 0 {
		result += "XL"
		numbers[2] = 0
	} else if x1 == 4 {
		result += "XC"
		numbers[2] = 0
		numbers[1] = 0
	}
	if i1 == 4 && v1 == 0 {
		numbers[4] = 0
	} else if i1 == 4 {
		numbers[4] = 0
		numbers[3] = 0
	}
	for i := 0; i < 5; i++ {
		result += repeat(literals[i], numbers[i])
	}
	if i1 == 4 && v1 == 0 {
		result += "IV"
	} else if i1 == 4 {
		result += "IX"
	}
	fmt.Println(result)

}
