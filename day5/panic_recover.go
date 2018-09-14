package main

import "fmt"

func main()  {
	fmt.Println("I am working")
	panic("i will down?")
	msg := recover()
	fmt.Println(msg)
}
