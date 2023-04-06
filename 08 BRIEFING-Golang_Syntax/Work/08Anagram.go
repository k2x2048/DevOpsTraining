package main

import (
	"fmt"
)

func IsAnagram(String1 string, String2 string) bool {
	if len(String1) != len(String2) {
		return false
	}

	SourceMap := make(map[rune]int)
	TargetMap := make(map[rune]int)

	for _, Letter := range String1 {
		SourceMap[Letter]++
	}

	for _, Letter := range String2 {
		TargetMap[Letter]++
	}

	for Letter, SourceCount := range SourceMap {
		if targetCount, ok := TargetMap[Letter]; !ok || SourceCount != targetCount {
			return false
		}
	}
	
	return true
}


func main() {
	var String1 string = "abc"
	var StringList []string = []string{"bca","fgh","hfhg","tty","bbb","abc","nkk","jjj","ddd"}

	for _, Word := range StringList {
		if IsAnagram(String1, Word) {
			fmt.Println(Word)
		}
	}
}
