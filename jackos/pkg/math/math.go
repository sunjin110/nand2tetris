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
	checkBit := int16(1) // 1 -> 10 -> 100 -> 1000となっていく
	shiftedX := x

	for i := 0; i < 16; i++ {
		if (y & checkBit) != 0 {
			sum += shiftedX
		}

		shiftedX = shiftedX << 1
		checkBit = checkBit << 1
	}
	return sum
}

// Devide 割り算
func Devide(x int16, y int16) int16 {

	if y > x {
		return 0
	}

	// q := Devide(x, 2*y)
	q := Devide(x, y<<1)

	//  (x - (2 * q * y)) < y
	if (x - (Multiply(q, y) << 1)) < y {
		// q * 2
		return q << 1
	}
	// 2*q + 1
	return (q << 1) + 1
}

// Sqrt y = √xの整数部分を計算する
func Sqrt(x int16) int16 {

	var y int16

	for i := (16/2-1); i > 0; i-- {

		if (y + )

	}


	return 0
}

// MultiplyVerbose 掛け算
// +と-だけで実装する
func MultiplyVerbose(x int16, y int16) int16 {

	// check
	if x == 0 || y == 0 {
		return 0
	}

	var sum int16
	checkBit := int16(1) // 1 -> 10 -> 100 -> 1000となっていく
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
