package main

import "fmt"

var a = []int{2,35,6,125,78,9,356}

func main()  {
	for x,y := range a{
		fmt.Printf("2**%d=%d\n",x,y)
	}
}
