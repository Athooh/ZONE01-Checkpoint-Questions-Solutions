// Seth  Solution
package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	if Atoi(args) <= 0 {
		return
	}
	
	word := ""
	if IsPowerOf2(Atoi(args)) {
		word = "true"
	} else {
		word = "false"
	}

	for _, c := range word {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}

func IsPowerOf2(num int) bool {
	return num > 0 && (num&(num-1)) == 0
}

func Atoi(s string) int {
	var number int
	sign := 1

	for idx, char := range s {
		if char == '-' && idx == 0 {
			sign = -1
		} else if char == '+' && idx == 0 {
			sign = 1
		} else if char >= '0' && char <= '9' {
			number = number * 10 + int(char -'0')
		} else {
			return 0
		}
	}
	return number * sign
}
