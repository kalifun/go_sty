package main

import "fmt"

func main() {
	i, j := 42, 27

	p := &i         //指向i
	fmt.Println(*p) //通过指针读取i
	*p = 31         //通过指针设置i
	fmt.Println(i)

	p = &j //指向j
	*p = *p / 5
	fmt.Println(j)
}
