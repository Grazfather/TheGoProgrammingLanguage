package main

import "fmt"

func isAnagram(s1, s2 string) bool {
	// 0. Shortcut: Since we count spaces and punctuation, any length mismatch means NO
	if len(s1) != len(s2) {
		return false
	}
	// 1. Make a counter map of each string
	m1 := make(map[rune]int)
	m2 := make(map[rune]int)
	for _, c := range s1 {
		m1[c]++
	}
	for _, c := range s2 {
		m2[c]++
	}
	// 2. Compare to make sure the count for each key matches
	for c, count := range m1 {
		if c2, ok := m2[c]; count != c2 || !ok {
			return false
		}
	}
	return true
}

func reportAnagram(s1, s2 string) {
	if isAnagram(s1, s2) {
		fmt.Printf("'%s' and '%s' ARE anagrams of eachother\n", s1, s2)
	} else {
		fmt.Printf("'%s' and '%s' are NOT anagrams of eachother\n", s1, s2)
	}
}

func main() {
	reportAnagram("apple", "orange")
	reportAnagram("rats", "star")
	reportAnagram("words and spaces", "spaces and words")
	reportAnagram("different length", " different length ")
}
