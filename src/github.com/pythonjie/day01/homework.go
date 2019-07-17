//--1.常量声明及练习题回去自己写一下，特别是iota定义数量级时（1024）
//package main
//
//import "fmt"
//
//const (
//	a    = 1
//	b    = 1 << (10 * iota)
//	c, d = 1 << (2 * iota), 1 << (3 * iota)
//)
//
//func main() {
//	fmt.Println(a, b, c, d)
//}

//--2.编写代码分别定义一个整型、浮点型、布尔型、字符串型变量，使用`fmt.Printf()`搭配`%T`分别打印出上述变量的值和类型
/*
package main

import "fmt"

var num int8
var num2 = 7
var (
	s1   string
	isOk bool
	num1 int
	fl   float32
)

func main() {
	fmt.Printf("num的类型为%T，s1的类型为%T,isOk的类型为%T,f1的类型为%T", num1, s1, isOk, fl)

}
 */

//--3.编写代码统计出字符串`"hello沙河小王子"`中汉字的数量。(自己查资料)
//package main
//
//import "fmt"
//
//var i = 0
//
//func main() {
//	s := "Hello沙河小王子"
//	for _, v := range s {
//		l := len(string(v))
//
//		if l > 1 {
//			i = i + 1
//		}
//	}
//	fmt.Println(i)
//}

//--4.打印九九乘法表
package main

import "fmt"

func main() {
	for i := 9; i >=1; i-- {
		fmt.Println("")
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
	}

}
