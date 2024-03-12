package main

import (
	"fmt"
	"math"
)

func divide(a, b int) int {
	if b == 0 || (a == math.MaxInt32 && b == -1) {
		return math.MaxInt32
	}
	sign := true
	if (a < 0) != (b < 0) {
		sign = false
	}

	if (!sign) && (a < 0) {
		a = -a
	} else if (!sign) && (b < 0) {
		b = -b
	}

	result := 0
	for a >= b {
		tmp, multiple := b, 1
		for a >= (tmp << 1) {
			tmp <<= 1
			multiple <<= 1
		}
		result += multiple
		a -= tmp
	}
	if sign {
		return result
	}

	return -result
}
func main() {
	var a, b int
	_, err := fmt.Scanf("a = %d, b = %d", &a, &b)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(divide(a, b))
}
