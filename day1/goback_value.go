package main

import "fmt"

func split(sum int) (x,y int) {
	x = sum*4
	y = sum-100
	return  x,y
}
func main()  {
	fmt.Println(split(1000))
}
