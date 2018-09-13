package main

import "fmt"

type test5 struct {
	x,y float64
}

//var m  = map[string]test5{
//	"kali" : test5{
//		123.456,456.789,
//	},
//	"fun" : test5{
//		789.123,456.123,
//	},
//}
var m  = map[string]test5{
	"kali" : {123.456,456.789},
	"fun" : {789.123,456.123},
}

func main()  {
	fmt.Println(m)
}
