package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)


/*
eggs (1)
peanuts (2)
shellfish (4)
strawberries (8)
tomatoes (16)
chocolate (32)
pollen (64)
cats (128)

func GetAllergies(Number int) []string {

}
*/


var substances = []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}
var allergy = map[string]int{
	"eggs":1,
	"peanuts":2,
	"shellfish":4,
	"strawberries":8,
	"tomatoes":16,
	"chocolate":32,
	"pollen":64,
	"cats":128,
}

//Allergies return list of allergies for a given score
func Allergies(score int) []string {
	var result []string
	for _, substance := range substances {
		if IsAllergicTo(score, substance) {
			result = append(result, substance)
		}
	}
	return result
}

func IsAllergicTo(score int, substance string) bool {
	return score&allergy[substance] == allergy[substance]
}

func main() {
	var TextLine string
	fmt.Print("Type allergy score: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	TextLine = scanner.Text()

	var Score, err = strconv.Atoi(TextLine)
	if err != nil || Score < 0 || Score > 255 {
		fmt.Println("Wrong code:", Score)
		os.Exit(0)
	}

	//Print Words occurences
	if Score > 0 {
		for _, Substance := range Allergies(Score) {
			fmt.Println("Allergic to ", Substance)
		}
	} else {
		fmt.Println("No allergies")
	}
}
