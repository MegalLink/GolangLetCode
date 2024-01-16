package begginer_problems

import (
	"fmt"
)

/**

 Roman to Integer
Easy
Topics
Companies
Hint
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, 2 is written as II in Roman numeral, just two ones added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.



Example 1:

Input: s = "III"
Output: 3
Explanation: III = 3.
Example 2:

Input: s = "LVIII"
Output: 58
Explanation: L = 50, V= 5, III = 3.
Example 3:

Input: s = "MCMXCIV"
Output: 1994
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.

Constraints:

1 <= s.length <= 15
s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
It is guaranteed that s is a valid roman numeral in the range [1, 3999].
*/
func romanToInt(s string) int {
	result := 0
	if len(s) < 1 || len(s) > 15 {
		return result
	}

	numberToRomanMapper := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for index := 0; index < len(s); index++ {
		actualValue := numberToRomanMapper[s[index]]
		if actualValue == 0 {
			fmt.Println("Not recognized character")
			return 0 // return 0 instead of actualValue
		}
		if index+1 < len(s) {
			nextValue := numberToRomanMapper[s[index+1]]
			if actualValue < nextValue {
				result = result + nextValue - actualValue
				index++ // Skip the next value since it has already been processed
				continue
			}
		}
		result = result + actualValue
	}

	return result
}

func romanToIntFastest(s string) int {
	sum := 0
	r := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	for i, el := range s {
		sum += r[string(el)]
		if i != 0 {
			if r[string(s[i-1])] < r[string(el)] {
				sum -= 2 * r[string(s[i-1])]
			}
		}
	}

	return sum
}
