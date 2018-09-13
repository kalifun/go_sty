package main

import (
	"fmt"
	"math"
)

func main()  {
	fun := func(x,y float64) float64 {
		return math.Sqrt(x*x+y*y)    //25的平方根
	}
	fmt.Println(fun(3,4))
}
