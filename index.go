package l4gd

import (
	"fmt"
	"github.com/lifeng1992/l4ge"
)

func CircleArea(r float64) float64 {
	result := l4ge.NumberPi() * r * r

	fmt.Println("====>circle_area", result)

	return result
}
