package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	fname := "D:\\go_sty\\dat5\\defer.txt"
	f, err := os.Open(fname)
	defer f.Close()
	if err != nil{
		os.Exit(1)
	}
	bread := bufio.NewReader(f)
	for {
		line, ok := bread.ReadString('\n')
		if ok != nil{
			break
		}
		fmt.Println(strings.Trim(line,"\r\n"))
	}

}