package main

import "fmt"

var a1 = [...]int{1,3,5,7,8}

func main()  {
	var b int
	for _,value := range a1{
		b += value
	}
	fmt.Println(b)
}