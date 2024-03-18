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
	// for i := 0; i < len(FirstNumber); i++ {
	// 	fmt.Println(FirstNumber[i])
	// }
	//Определение второго числа
	LastNumber0 := text[2]
	LastNumber := strings.Split(LastNumber0, "")
	// Проверка на тип данных
	// fmt.Printf("%T", text)
	// a1, _ := strconv.Atoi(FirstNumber[0])
	// if a1 <= 10 && a1 >= 0 {
	// 	fmt.Println(check(text))
	if FirstNumber[0] == "1" || FirstNumber[0] == "2" || FirstNumber[0] == "3" || FirstNumber[0] == "4" || FirstNumber[0] == "5" || FirstNumber[0] == "6" || FirstNumber[0] == "7" || FirstNumber[0] == "8" || FirstNumber[0] == "9" || FirstNumber[0] == "10" {
		fmt.Println(check(text))
	} else if FirstNumber[0] == "I" || FirstNumber[0] == "V" || FirstNumber[0] == "X" {
		// } else {
		// for i := 0; i < len(LastNumber); i++ {
		// 	fmt.Println(LastNumber[i])
		// }
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
		//
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
			} else {
			}
			fmt.Println(FirstNumber[i])
			fmt.Println(FirstRome[0])
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
				// fmt.Println(SecondRome)
			}
			result = 0
			if i == len(FirstRome)-1 {
				for i := len(FirstRome) - 1; i >= 0; i-- {
					result = result + FirstRome[i]
				}
			}
		}

		fmt.Println("Result 1 = ", result)
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
			// fmt.Println(LastNumber[i])
		}
		//Перевод второго числа в арабское
		var result2 int = 0
		for i := 1; i < len(SecondRome); i++ {
			if len(SecondRome) == 1 {
				result2 = SecondRome[0]
			} else if SecondRome[i] > SecondRome[i-1] {
				SecondRome[i] -= SecondRome[i-1]
				SecondRome[i-1] = 0
				// fmt.Println(SecondRome)
			}
			if i == len(SecondRome)-1 {
				for i := len(SecondRome) - 1; i >= 0; i-- {
					result2 = result2 + SecondRome[i]
				}
			}
		}

		// fmt.Println(FirstRome)
		// fmt.Println(SecondRome)
		fmt.Println("Result 2 = ", result2)
		text3 := []string{strconv.Itoa(result), text[1], strconv.Itoa(result2)}
		fmt.Println(check(text3))
	}
}
