package main

import (
	"fmt"
)

func no_fallthrough(score int) {
	var grade string
	switch {
	case score >= 95:
		grade = "A+"
	case score >= 90:
		grade = "A"
	case score >= 85:
		grade = "B+"
	case score >= 80:
		grade = "B"
	case score >= 75:
		grade = "C+"
	case score >= 70:
		grade = "C"
	case score >= 65:
		grade = "D+"
	case score >= 60:
		grade = "D"

	default:
		grade = "F"
	}

	fmt.Println("당신의 학점은", grade, "입니다.")
}
func main() {
	fmt.Println("교수님 제 성적이 어떻게 되나요?")
	score := 93
	no_fallthrough(score)
}
