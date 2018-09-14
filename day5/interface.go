package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {

}

func (nokiaphone NokiaPhone) call() {
	fmt.Println("I am NokiaPhone")

}

type IPhone struct {
	
}

func (iphone IPhone) call() {
	fmt.Println("I am IPhone")
}

func main()  {
	var phone Phone

	phone  =  new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}


