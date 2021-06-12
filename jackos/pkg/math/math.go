package math

import "fmt"

// Multiply 掛け算
// +と-だけで実装する
func Multiply(x int16, y int16) int16 {

	// check
	if x == 0 || y == 0 {
		return 0
	}

	var sum int16
	checkBit := int16(1)
	shiftedX := x

	for i := 0; i < 16; i++ {

		fmt.Println(i, ":回目")

		fmt.Printf("y is %b\n", y)
		fmt.Printf("checkBit is %b\n", checkBit)
		fmt.Printf("shiftedX is %b\n", shiftedX)
		fmt.Printf("y & checkBit = %b\n", y&checkBit)

		if (y & checkBit) != 0 {
			fmt.Printf("=== sum ===\n")
			fmt.Printf("%d = %d + %d\n", sum+shiftedX, sum, shiftedX)

			sum = sum + shiftedX
		}

		shiftedX = shiftedX << 1
		checkBit = checkBit << 1
	}

	return sum
}

// // Multiply 掛け算
// // +と-だけで実装する
// func Multiply(x int16, y int16) int16 {

// 	// check
// 	if x == 0 || y == 0 {
// 		return 0
// 	}

// 	var sum int16
// 	shiftedX := x

// 	for i := 0; i < 16; i++ {

// 		fmt.Println(i, ":回目")

// 		fmt.Println("y & shiftedX = ", y&shiftedX)

// 		if (y & shiftedX) == 1 {
// 			sum = sum + x
// 		}

// 		shiftedX = shiftedX << 1
// 		fmt.Println("shiftedX is ", shiftedX)
// 	}

// 	return sum
// }
