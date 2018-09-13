package main

import "fmt"

func main()  {
	s := []int{2,423,22,6,8,2,56,9,29}
	fmt.Println("s==",s)
	fmt.Println("s[1:4]==",s[1:4])
	//省略下标表示从0开始
	fmt.Println("s[:5]==",s[:5])
	//省略上标表示到len(s）结束
	fmt.Println("s[2:]",s[2:])
}