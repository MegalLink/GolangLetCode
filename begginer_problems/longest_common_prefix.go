package begginer_problems

import "unicode/utf8"

/**
14. Longest Common Prefix
Easy
Topics
Companies
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lowercase English letters.*/

// my solution
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i, charPrefix := range strs[0] {
		for _, charWord := range strs[1:] {
			// si longitud de word es mayor la longitud de word
			//prefix=ab word=a => prefix = a  prefix=[:i]
			// para i = 1 si chardWord=a charPrefix=1 no hacer nada pero si chardWord =b y charPrefix=a => prefix:= prefix=[:i]
			if i >= len(charWord) || charWord[i] != byte(charPrefix) {
				return prefix[:i]
			}
		}
	}

	return prefix
}

// better solution
func longestCommonPrefixFaster(strs []string) string {
	var runes [][]rune

	// transform strs input to rune and get the largest word with his length so it dont transform in every loop
	longestLength := -1
	longestIndex := -1
	for index, str := range strs {
		runes = append(runes, []rune(str))
		runeCount := utf8.RuneCountInString(str)
		if longestLength < runeCount {
			longestIndex = index
			longestLength = runeCount
		}
	}
	if longestLength == 0 {
		return ""
	}

	matchedLength := 0

	// this is a labeled loop feature from goland it breaks the two loops when break is called
matchLoop:
	// recorre el prefijo mas largo
	for runeIndex, aRune := range runes[longestIndex] {
		for allRunesIndex, aRunes := range runes {
			//dont iterate over longest prefix
			if allRunesIndex == longestIndex {
				continue
			}
			//salir del loop
			if len(aRunes) <= runeIndex || aRune != aRunes[runeIndex] {
				break matchLoop
			}
		}
		matchedLength++
	}

	return string(runes[longestIndex][0:matchedLength])
}
