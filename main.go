package main

import (
	"fmt"
)

func main() {
	var s, words_ string
	fmt.Scan(&s, &words_)
	// words := make([]string, 0)
	// temp := make([]byte, 0)
	// for _, v := range words_ {
	// 	if v == rune(',') {
	// 		words = append(words, string(temp))
	// 		temp = []byte{}
	// 	}
	// 	temp = append(temp, byte(v))
	// }
	words := []string{"acd"}
	count := 0
	for _, word := range words {
		sIdx, wIdx := 0, 0
		for ; sIdx < len(s) && wIdx < len(word); sIdx++ {
			if s[sIdx] == word[wIdx] {
				wIdx++
			}
		}
		if wIdx == len(word) {
			fmt.Println(word)
			count++
		}
	}
	fmt.Println(count)
}
