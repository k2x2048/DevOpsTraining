package main

import (
	"fmt"
	"bufio"
	"os"
	"time"
)

//+-31.688 years
const Gigasecond = time.Second * 1e9

func main() {
	var TextLine string
	fmt.Print("What is your birthdate (YYYY-MM-DD HH:MM:SS) ?: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TextLine = scanner.Text()
	
	//TextLine = "1975-10-15 14:30:27"
	layout := "2006-01-02 15:04:05"
	
	BirthDate, err := time.Parse(layout, TextLine)
	if err != nil {
		fmt.Println("Could not parse date:", err)
	} else {
		var GigaBirthDay time.Time = BirthDate.Add(Gigasecond)
		fmt.Println("Date of your GigaBirthDay:", GigaBirthDay.Format(layout))
	}
	//fmt.Println("Seconds : ", time.Duration.Seconds(Gigasecond))
	//fmt.Println("Hours : ", time.Duration.Hours(Gigasecond))
	//fmt.Println("Years : ",Gigasecond/31207680)
}
