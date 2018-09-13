package main

import "fmt"

type test4 struct {
	lat,long float64
}

var m map[string]test4

func main()  {
	m = make(map[string]test4)
	m["FUN"] = test4{
		40.54412,-884.4445554,
	}
	fmt.Println(m["FUN"])
}
