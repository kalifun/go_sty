package main

import "fmt"

func main()  {
	q := make(map[string]int)

	q["fun"] = 45
	fmt.Println(q["fun"])

	q["fun"] = 41
	fmt.Println(q["fun"])

	delete(q,"fun")
	fmt.Println("the value:",q["fun"])

	v, ok := q["fun"]
	fmt.Println("the value:",v,"really",ok)


}
