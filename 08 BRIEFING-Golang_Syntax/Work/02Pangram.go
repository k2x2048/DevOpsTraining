package main

import (
	"fmt"
	"strings"
)

//const colorRed = "\033[0;31m"
//const colorGreen = "\033[0;31m"

const colorNone = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
//const colorYellow = "\033[33m"
//const colorBlue = "\033[34m"
//const colorPurple = "\033[35m"
//const colorCyan = "\033[36m"
//const colorWhite = "\033[37m"

func isPangram(Sentense string) bool {
	var SentenseLc = strings.ToLower(Sentense)
	for code := 97; code < 123 ; code++ {
		CurChar := string(code)
		//fmt.Println("'", CurChar, "' in '", SentenseLc, "'?")
		if !strings.Contains(SentenseLc, CurChar) {
			return false
		}
	}
	return true
}

/*
func isPangram(Sentense string) bool {
	var Alphabet = "abcdefghijklmnopqrstuvwxyz"
	var SentenseLc = strings.ToLower(Sentense)
	for pos := 0; pos < len(Alphabet) ; pos++ {
		CurChar := string(Alphabet[pos:pos+1])
		//fmt.Println("'", CurChar, "' in '", SentenseLc, "'?")
		if !strings.Contains(SentenseLc, CurChar) {
			return false
		}
	}
	return true
}
*/

func main() {
	var Sentense string
	
	fmt.Print("write a sentense: ")
	fmt.Scanln(&Sentense)
	
	if isPangram(Sentense) {
		fmt.Println(colorGreen, "'", Sentense, "' is a Pangram.", colorNone)
	} else {
		fmt.Println(colorRed, "'", Sentense, "' is not a Pangram.", colorNone)
	}

}
