package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculator(a, b int, s string) int {
	switch s {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		if b == 0 {
			panic("Выдача паники, деление на 0 невозможно")
		}
		return a / b
	}
}

func check(textArr []string) []string {
	if len(textArr) < 3 { //проверка на корректность что введенной строке не меньше 3-х элементов.
		panic("Выдача паники, строка не является математической операцией.")
	} else if len(textArr) > 3 { //проверка на корректность что введенной строке не больше 3-х элементов.
		panic("Выдача паники, так как формат математической операции не " +
			"удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	operator := "+-*/"
	if !strings.Contains(operator, textArr[1]) { //проверка что орератор верный.
		panic("Выдача паники, так как введен некорректный оператор.")
	}
	return textArr
}

func arabOrRome(a, b string) (first, second int, arab, rome bool) {

	firstArab, err1 := strconv.Atoi(a)
	secondArab, err2 := strconv.Atoi(b)

	if err1 == nil && err2 == nil {
		if firstArab < 1 || firstArab > 10 || secondArab < 1 || secondArab > 10 {
			panic("Паника: введены некорректные числа. Числа должны быть от 1 до 10")
		}
		arab = true
		return firstArab, secondArab, true, false
	}

	firstRome, ok1 := romeToArab(a)
	secondRome, ok2 := romeToArab(b)

	if ok1 && ok2 {
		if firstRome < 1 || firstRome > 10 || secondRome < 1 || secondRome > 10 {
			panic("Паника: введены некорректные числа. Числа должны быть от 1 до 10")
		}
		rome = true
		return firstRome, secondRome, false, true
	}

	if (err1 == nil && ok2) || (ok1 && err2 == nil) {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	panic("Выдача паники, так как введен некорректный операнд.")

}

func romeToArab(rome string) (arab int, ok bool) {
	romeNum := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	for i := range rome {
		if val, found := romeNum[string(rome[i])]; found {
			if i+1 < len(rome) && romeNum[string(rome[i+1])] > val {
				arab -= val
			} else {
				arab += val
			}
		} else {
			return 0, false
		}
	}
	return arab, true
}

func arabToRome(num int) string {
	arabToRomeMap := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}
	keys := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	answerRome := ""
	for _, key := range keys {
		for num >= key {
			num -= key
			answerRome += arabToRomeMap[key]
		}
	}
	return answerRome
}

func safeExecute() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	textArr := strings.Split(text, " ")

	textArr = check(textArr)
	firstNum, secondNum, arab, rome := arabOrRome(textArr[0], textArr[2])

	result := calculator(firstNum, secondNum, textArr[1])

	if arab {
		fmt.Println(result)
	} else if rome {
		if result < 1 {
			panic("Выдача паники, так как в римской системе нет чисел меньше 1")
		}
		fmt.Println(arabToRome(result))
	}
}

func main() {
	safeExecute()
}
