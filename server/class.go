package server

import (
	"fmt"
	"math"
)

func Mater() {
	floatValue1 := 0.1
	floatValue2 := 0.1
	p := 0.00001
	// 判断 floatValue1 与 floatValue2 是否相等
	if math.Dim(float64(floatValue1), floatValue2) < p {
		fmt.Println("floatValue1 和 floatValue2 相等")
	}
}
