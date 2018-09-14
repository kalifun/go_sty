package main

import "fmt"

func set_value(x *int)  {
	*x = 200
}
func main()  {
	x := new(int)
	set_value(x)
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(*x)
}
