package main

import "fmt"

func main() {
	playString()
	playNumber()
	playBool()
}

func playBool() {
	var isMe bool = true
	isYou := true
	fmt.Println("bool", isMe, isYou)
}

func playNumber() {
	var ageOne int = 123456789
	var ageTwo int8 = 127
	var ageTree int16 = 256
	var ageFour int32 = 512
	fmt.Println("int", ageOne, ageTwo, ageTree, ageFour)

	var floatOne float32 = 25.98
	var floatTwo float64 = 50.55
	fmt.Println("float", floatOne, floatTwo)
}

func playString() {
	var text string = "Puck"
	textAutoType := "Auto"
	var textLateAssign string

	fmt.Println(text, textAutoType, textLateAssign)

	textAutoType = ""
	textLateAssign = "newTextLateAssign"

	fmt.Println("string", text, textAutoType, textLateAssign)

	fmt.Printf("you %v me %v \n", 35, "You")
	fmt.Printf("you %q me %v \n", 35, "You")
	fmt.Printf("you %q me %v \n", "Puck", "You")
}
