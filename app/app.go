package app

import (
	"calc/variables"
	"fmt"
	"strconv"
	"strings"
)

func Base(s string) {
	var operator string
	var stringsFound int
	numbers := make([]int, 0)
	romans := make([]string, 0)
	romansToInt := make([]int, 0)
	for idx := range variables.Operators {
		for _, val := range s {
			if idx == string(val) {
				operator += idx
				variables.Data = strings.Split(s, operator)
			}
		}
	}
	switch {
	case len(operator) > 1:
		panic(variables.Long)
	case len(operator) < 1:
		panic(variables.NoOperation)
	}
	for _, elem := range variables.Data {
		num, err := strconv.Atoi(elem)
		if err != nil {
			stringsFound++
			romans = append(romans, elem)
		} else {
			numbers = append(numbers, num)
		}
	}

	switch stringsFound {
	case 1:
		panic(variables.Different)
	case 0:
		errCheck := numbers[0] > 0 && numbers[0] < 11 &&
			numbers[1] > 0 && numbers[1] < 11
		if val, ok := variables.Operators[operator]; ok && errCheck == true {
			variables.AIn, variables.BIn = &numbers[0], &numbers[1]
			fmt.Println(val())
		} else {
			panic(variables.NoValid)
		}
	case 2:
		for _, elem := range romans {
			if val, ok := variables.Roman[elem]; ok && val > 0 && val < 11 {
				romansToInt = append(romansToInt, val)
			} else {
				panic(variables.NoValid)
			}
		}
		if val, ok := variables.Operators[operator]; ok {
			variables.AIn, variables.BIn = &romansToInt[0], &romansToInt[1]
			intToRoman(val())
		}
	}
}
func intToRoman(romanResult int) {
	var romanNum string
	if romanResult == 0 {
		panic(variables.Zero)
	} else if romanResult < 0 {
		panic(variables.Negative)
	}
	for romanResult > 0 {
		for _, elem := range variables.ConvIntToRoman {
			for i := elem; i <= romanResult; {
				for index, value := range variables.Roman {
					if value == elem {
						romanNum += index
						romanResult -= elem
					}
				}
			}
		}
	}
	fmt.Println(romanNum)
}
