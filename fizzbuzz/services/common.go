package services

import "fmt"

/*
- Accepts five parameters : three integers int1, int2 and limit, and two strings str1 and str2.
- Returns an array of strings with numbers from 1 to limit, where: all multiples of int1 are replaced by str1,
all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
 */
func CommonFunc(int1, int2, limit int, str1, str2 string) []string {

	var result = make([]string,limit,limit)

	for i := 1; i <= limit; i++ {

		var str_i string

		if i%int1 == 0 {
			// this number is multiple of int1
			str_i = str1
		}
		if i%int2 == 0 {
			// this number is multiple of int2

			if i%int1 == 0 {
				// this number is multiple of int1 and int2
				str_i = fmt.Sprintf("%s%s", str_i, str2)
			} else{
				str_i = str2
			}

		}

		if i%int1 != 0 && i%int2 != 0 {
			// in this case, get the number itself
			str_i = fmt.Sprintf("%d", i)
		}

		result[i-1] = str_i
	}

	return result

}
/*
Call CommonFunc with int1=3, int2=5, limit=100, str1=fizz, str2=buzz
 */
func fizzbuzz() []string {

	var result = CommonFunc(3, 5, 100, "fizz", "buzz")

	return result

}

/*
Print an Array
 */
func printArray(array []string) {
	for i := 1; i <= len(array); i++ {
		fmt.Printf(array[i-1])
		fmt.Printf("\n")
	}
}
