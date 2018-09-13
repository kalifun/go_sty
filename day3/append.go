package main

import "fmt"

func main()  {
	var z []int
	printSlice("z",z)

	z = append(z,0)
	printSlice("z",z)

	z = append(z,1)
	printSlice("z",z)

	z = append(z,2,3,4,5)
	printSlice("z",z)
}
func printSlice(s string,x []int)  {
	fmt.Printf("%s len=%d cap=%d %v\n",s,len(x),cap(x),x)
}
