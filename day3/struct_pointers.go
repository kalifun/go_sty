package main

import "fmt"

type test2 struct {
	x int
	y int
}

func main()  {
	v := test2{2,3}
	p := &v
	p.x = 9
	fmt.Println(v)
}
