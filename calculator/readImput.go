package calculator

import (
	"fmt"
)

func ReadInput() (int, int) {
	var num1, num2 int
	_, err := fmt.Scanln(&num1, &num2)
	if err != nil {
		return 0, 0
	}

	return num1, num2
}
