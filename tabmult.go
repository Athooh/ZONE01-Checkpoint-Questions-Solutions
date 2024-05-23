// tabmult
// Instructions
// Write a program that displays a number's multiplication table.

// The parameter will always be a strictly positive number that fits in an int. Said number multiplied by 9 will also fit in an int.
// Usage
// $ go run . 9
// 1 x 9 = 9
// 2 x 9 = 18
// 3 x 9 = 27
// 4 x 9 = 36
// 5 x 9 = 45
// 6 x 9 = 54
// 7 x 9 = 63
// 8 x 9 = 72
// 9 x 9 = 81
// $ go run . 19
// 1 x 19 = 19
// 2 x 19 = 38
// 3 x 19 = 57
// 4 x 19 = 76
// 5 x 19 = 95
// 6 x 19 = 114
// 7 x 19 = 133
// 8 x 19 = 152
// 9 x 19 = 171
// $ go run .

// $

// package main

// import (
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		return
// 	}

// 	args := os.Args[1]
// 	numb, err := strconv.Atoi(args)
// 	if err != nil || numb <= 0 {
// 		fmt.Println("Error! Please provide a valid number.")
// 		return
// 	}

// 	for i := 1; i <= 9; i++ {
// 		fmt.Printf("%d x %d = %d\n", i, numb, i*numb)
// 	}
// }

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	stInt := os.Args[1]
	var number int
	var outPut []string
	for i, v := range stInt {
		if i == 0 && v == '-' {
			return
		}
		n := int(v - 48)
		number = ((number * 10) + n)
	}

	for i := 1; i <= 9; i++ {
		newStr := string(i + 48)
		res := (i * number)
		resStr := ""

		for res > 0 {
			digit := res % 10
			resStr = string(digit+'0') + resStr
			res /= 10
		}

		outPut = append(outPut, newStr+" "+"x"+" "+stInt+" "+"="+" "+resStr)

	}
	for i := 0; i < len(outPut); i++ {
		for _, v := range outPut[i] {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}

}

/*
John Odhiambo's Solution

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(str string) int {
	runes := []rune(str)

	LenRune := 0
	result := 0

	for i := range runes {
		LenRune = i + 1
	}

	if LenRune == 0 {
		return 0
	}

	tens := 1

	for k := 0; k < LenRune-1; k++ {
		if runes[k] == '+' || runes[k] == '-' {
			continue
		}

		tens *= 10
	}

	for i := range runes {
		if (runes[0] == '-' || runes[0] == '+') && i == 0 {
			continue
		}

		if runes[i] < '0' || runes[i] > '9' {
			return 0
		}

		numb := 0

		for j := '0'; j < runes[i]; j++ {
			numb++
		}

		result += numb * tens
		tens /= 10
	}

	if runes[0] == '-' {
		result *= -1
	}

	return result
}

func Itoa(num int) string {
	if num == 0 {
		return "0"
	}

	isNegativeNum := false

	if num < 0 {
		isNegativeNum = true
		num = -num
	}

	var result []byte

	for num > 0 {
		digit := num % 10
		result = append([]byte{byte('0' + digit)}, result...)
		num /= 10
	}

	if isNegativeNum {
		result = append([]byte{'-'}, result...)
	}
	return string(result)
}

func tabMult() {
	argumentString := os.Args[1]

	num := Atoi(argumentString)

	for i := 1; i <= 9; i++ {
		z01.PrintRune(rune('0' + i))
		z01.PrintRune(' ')
		z01.PrintRune('*')
		z01.PrintRune(' ')

		for _, ch := range Itoa(num) {
			z01.PrintRune(ch)
		}

		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')

		for _, ch := range Itoa(i * num) {
			z01.PrintRune(ch)
		}
		z01.PrintRune('\n')
	}
}

func main() {
	tabMult()
}

*/