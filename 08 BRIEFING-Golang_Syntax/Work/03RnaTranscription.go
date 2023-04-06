package main

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

//const colorRed = "\033[0;31m"
//const colorGreen = "\033[0;31m"

const colorNone = "\033[0m"
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorYellow = "\033[33m"
//const colorBlue = "\033[34m"
//const colorPurple = "\033[35m"
//const colorCyan = "\033[36m"
//const colorWhite = "\033[37m"

//The bases in DNA are Adenine (‘A’), Thymine (‘T’), Guanine (‘G’) and Cytosine (‘C’).
//RNA shares Adenine (‘A’), Guanine (‘G’) and Cytosine (‘C’) with DNA, but contains Uracil (‘U’) rather than Thymine.

//A -> U
//C -> G
//T -> A
//G -> C


func DNAtoRNA(DNA string) string {
	var RNA string

//func Replace(s, old, new string, n int) string
//Here, s is the original or given string, old is the string that you want to replace. new is the string which replaces the old, and n is the number of times the old replaced.
//Note: If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string. If n < 0, there is no limit on the number of replacements.

	for pos := 0; pos < len(DNA) ; pos++ {
		Base := string(DNA[pos:pos+1])
		var NewBase string
		switch Base {
		case "a", "A":
			NewBase = "U"
		case "c", "C":
			NewBase = "G"
		case "t", "T":
			NewBase = "A"
		case "g", "G":
			NewBase = "C"
		default:
			NewBase = Base
		}
		RNA = RNA + NewBase
	}
	return RNA
}

func main() {
	var Sentense string
	
	fmt.Print("Enter DNA string: ")
	
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	Sentense = scanner.Text()
	
	fmt.Println("DNA string: ",colorGreen , strings.ToUpper(Sentense), colorNone)
	fmt.Println("RNA string: ",colorYellow,  strings.ToUpper(DNAtoRNA(Sentense)), colorNone)
	
}
