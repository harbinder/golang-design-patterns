package strings

import "fmt"

/*
We use a hash and three variables

Hash keeps track of the last index position of any character
longestSubstringLength – this holds the longest substring length without repeating character seen so far
currentSubstringLength – this holds the current substring length without repeating character
start – this indicates the start of the current substring without repeating character
We iterate over the string and check this hash for the current character. We simply increment the currentSubstringLength in the below two conditions

If the entry is not present for the current character then the current character has not been seen before.
If the entry is present and the current character has been seen before but it is not part of the current length.
Otherwise

We reset the start position and currentSubstringLength to include the current character in the current length.
Before resetting we check if currentSubstringLength is greater than longestSubstringLength.
If yes then we set longestSubstringLength to currentSubstringLength.
*/
func LongestSubstring() {
	str := "abbabcda"
	fmt.Println("Input string: ", str)

	charLastIndex := make(map[string]int)
	var currStrLen, LongestStrLen int
	strStartIndex := 0
	for index := range str {
		fmt.Printf("\nindex: %v, char: %v", index, str[index])

		lastIndex, ok := charLastIndex[string(str[index])]
		if !ok || lastIndex < index-currStrLen {
			currStrLen++
		} else {
			if currStrLen > LongestStrLen {
				LongestStrLen = currStrLen
			}
			strStartIndex = index + 1
			currStrLen = (index - strStartIndex) + 1
		}
		charLastIndex[string(str[index])] = index

		fmt.Println(strStartIndex)
	}
	if currStrLen > LongestStrLen {
		LongestStrLen = currStrLen
	}

	fmt.Println(charLastIndex)
	fmt.Println("Longest SubString Len: ", LongestStrLen)
	var substr string
	for char := range charLastIndex {
		substr = fmt.Sprintf("%s%s", substr, char)
	}

}
