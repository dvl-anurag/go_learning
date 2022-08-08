package basics

import "fmt"

const Pi = 3.14

func Consta() {
	const World = "welcome to go"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
