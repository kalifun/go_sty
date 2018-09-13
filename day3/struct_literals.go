package main

import "fmt"

type test3 struct {
	x,y int
}

var (
	v1 = test3{1,2}
	v2 = test3{x:3}
	v3 = test3{y:4}
	p = &test3{5,6}
)
func main()  {
	fmt.Println(v1,v2,v3,p)
}