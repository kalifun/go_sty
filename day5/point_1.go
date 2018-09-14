package  main

import "fmt"

func change(x *int)  {
	*x = 200
}

func main()  {
	var x int = 100
	fmt.Println(x)
	change(&x)
	fmt.Println(x)

}