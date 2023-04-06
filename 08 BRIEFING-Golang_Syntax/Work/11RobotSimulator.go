package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"bufio"
)

type direction int

const (
	North direction = iota //+y
	East                   //+x
	South                  //-y
	West                   //-x
)

func DirectionToString(PDir direction) string {
	switch PDir {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	}
	return "?"
}

type action int

const (
	Forward action = iota
	Right
	Backward
	Left
)

type robot struct {
	name     string
	position struct {
		x         int
		y         int
		direction direction
	}
}

func (e *robot) move(MoveAction action) {
	switch MoveAction {
	case Forward:
		switch e.position.direction {
		case North:
			e.position.y += 1
		case East:
			e.position.x += 1
		case South:
			e.position.y -= 1
		case West:
			e.position.x -= 1
		}
	case Right:
		switch e.position.direction {
		case North:
			e.position.direction = East
		case East:
			e.position.direction = South
		case South:
			e.position.direction = West
		case West:
			e.position.direction = North
		}
	case Backward:
		switch e.position.direction {
		case North:
			e.position.y -= 1
		case East:
			e.position.x -= 1
		case South:
			e.position.y += 1
		case West:
			e.position.x += 1
		}
	case Left:
		switch e.position.direction {
		case North:
			e.position.direction = West
		case East:
			e.position.direction = South
		case South:
			e.position.direction = East
		case West:
			e.position.direction = North
		}
	}
}

func (e *robot) getposition() string {
	return "{" + strconv.Itoa(e.position.x) + ", " + strconv.Itoa(e.position.y) + "} facing " + DirectionToString(e.position.direction)
}

func main() {

	MyRobot := new(robot)
	//MyRobot := &robot{}
	MyRobot.name = "iCar"
	MyRobot.position.direction = North
	MyRobot.position.x = 0
	MyRobot.position.y = 0

	//argsWithProg := os.Args
	var RobotCommands []rune
	if len(os.Args) > 1 {
		RobotCommands = []rune(strings.Join(os.Args[1:], ""))
	} else {
		fmt.Print("Enter commands (R=right, F=forward, L=left, B=backwards):")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		RobotCommands = []rune(scanner.Text())
		if len(RobotCommands) == 0 {
			fmt.Println("No commands. No move. Bye bye.")
			os.Exit(0)
		}
	}
	fmt.Println("Initial position:", MyRobot.getposition())

	for _, CurChar := range RobotCommands {
		//fmt.Printf("CurChar: %c\n", CurChar)
		switch CurChar {
		case 'R', 'r': //Right
			fmt.Println("turn Right")
			MyRobot.move(Right)
		case 'F', 'f', 'A', 'a': //Forward
			fmt.Println("move Forward")
			MyRobot.move(Forward)
		case 'L', 'l': //Left
			fmt.Println("turn Left")
			MyRobot.move(Left)
		case 'B', 'b': //Backward
			fmt.Println("move Backwards")
			MyRobot.move(Backward)
		}
	}
	fmt.Println("final position:", MyRobot.getposition())
}
