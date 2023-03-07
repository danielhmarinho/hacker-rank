package main

import (
	"fmt"
)

func main() {
	// fmt.Println(IsOneAway("pale", "ple"))
	// fmt.Println(IsOneAway("pales", "pale"))
	// fmt.Println(IsOneAway("pale", "bale"))
	// fmt.Println(IsOneAway("pale", "bake"))
	fmt.Println(IsOneAway("pale", "palle"))
}

func IsOneAway(s1 string, s2 string) bool {
	s1len := len(s1)
	s2len := len(s2)
	diff := s1len - s2len

	if diff > 0 {
		return isOneInsertAway(s2, s1)
	}

	if diff < 0 {
		return isOneInsertAway(s1, s2)
	}

	return isOneReplceAway(s1, s2)
}

// isOneInsertAway returns true if given strings are one insert away
// s1 should be the sortest string always
func isOneInsertAway(s1 string, s2 string) bool {
	edits := 0

	for i, c := range s1 {
		if c != rune(s2[i+edits]) {
			edits++
		}

		if edits > 1 {
			return false
		}
	}

	return true
}

// isOneReplceAway returns true if given strings are one replace away
func isOneReplceAway(s1 string, s2 string) bool {
	edits := 0

	for i, c := range s1 {
		if c != rune(s2[i]) {
			edits++
		}

		if edits > 1 {
			return false
		}
	}

	return true
}
