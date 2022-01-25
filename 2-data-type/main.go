package main

import (
	"fmt"
	"strings"
)

func main() {
	playString()
	playNumber()
	playBool()
	playArray()
}

func playArray() {
	var days [4]int = [4]int{1, 2, 3, 4}
	var names = [2]string{"Puck", "Muck"}

	fmt.Println("Array", days, names)
	fmt.Println("Array len", len(days), len(names))

	fmt.Println("Array slice", days[1:3], days[1:], days[:2])

	names[0] = "Suck"
	fmt.Println("Array append", append(names[:1], "Duck"), names)
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
	fmt.Println("strings.Contains", strings.Contains(textLateAssign, "new"))
	fmt.Println("strings.ReplaceAll", strings.ReplaceAll(textLateAssign, "new", "ex"))

	fmt.Printf("Printf you %v me %v \n", 35, "You")
	fmt.Printf("Printf you %q me %v \n", 35, "You")
	fmt.Printf("Printf you %q me %v \n", "Puck", "You")
}
