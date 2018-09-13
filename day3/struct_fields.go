package main

import "fmt"

type test1 struct {
	x int
	y int
}

func main()  {
	v := test1{2,3}
	v.x =  10
	fmt.Println(v.x)
}