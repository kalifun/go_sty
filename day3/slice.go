package main

import "fmt"

func main()  {
	p := []int{2,3,4,5,6,7,8,9}
	fmt.Println("p ==",p)
	for i := 0;i < len(p) ;i++  {
		fmt.Printf("p[%d]==%d\n",i,p[i])
	}
}