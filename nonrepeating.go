package main

import "fmt"

func main() {

	fmt.Println(lengthOfNonRepeatingSubstr("abcabcbb"))

	fmt.Println(lengthOfNonRepeatingSubstr("bbbbb"))

	fmt.Println(lengthOfNonRepeatingSubstr("abbabab"))

}

func lengthOfNonRepeatingSubstr(s string) int {

	lastOccured := make(map[byte]int)
	start := 0
	maxLength := 0

	for i, ch := range []byte(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccured[ch] = i
	}
	return maxLength
}
