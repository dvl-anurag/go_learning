package basics

import "fmt"

func WithParam(a, b int) int {
	fmt.Println(a + b)
	c := a + b
	fmt.Println(Split(c))
	return c

}
func Swap(x, y int) (int, int) {
	return y, x
}

func Split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
