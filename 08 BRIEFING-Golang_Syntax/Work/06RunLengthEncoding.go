package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func RLE_Encode(DataString string) string {

	//Exit if empty DataString
	if DataString == "" { return "" }

	PreviousChar := DataString[0:1]
	EncodeString := ""
	Count := 0

	for _, CurChar := range DataString {
		if string(CurChar) == PreviousChar {
			Count += 1
		} else {
			if Count > 1 {
				EncodeString = EncodeString + strconv.Itoa(Count) + string(PreviousChar)
			} else {
				EncodeString = EncodeString + string(PreviousChar)
			}
			//reset counter for new char
			Count = 1
			PreviousChar = string(CurChar)
		}
	}
	return EncodeString
}


func main() {
	var TextLine string
	fmt.Print("Enter string to encode: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TextLine = scanner.Text()
	
	fmt.Println("Encoded String:", RLE_Encode(TextLine))
}
