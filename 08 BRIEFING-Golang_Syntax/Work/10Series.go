package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func All(n int, s string) (res []string) {
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
	}
	return
}

func main() {

	fmt.Print("Enter some digits: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	StringDigits := scanner.Text()
	StringDigitsLen := len(StringDigits)

	if StringDigitsLen < 1 {
		fmt.Println("Invalid entry.")
		os.Exit(0)
	}

	fmt.Print("Enter lenght of substrings (between 1 and ", StringDigitsLen, "): ")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TextLine := scanner.Text()
	var SubstringLengh, err = strconv.Atoi(TextLine)
	if err != nil {
		fmt.Println("Wrong entry:", TextLine)
		os.Exit(0)
	}

	if StringDigitsLen < SubstringLengh || SubstringLengh < 1 {
		fmt.Println("Invalid SubstringLengh:", TextLine)
		os.Exit(0)
	}

	var RStrings []string = All(SubstringLengh, StringDigits)

	for i, String := range RStrings {
		fmt.Println("String ", i+1, ": ", String)
	}

}
