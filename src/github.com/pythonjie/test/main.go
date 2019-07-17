package main

// 流程控制之跳出for循环

// 当i=5时就跳出for循环

//func main() {
//	for i:=0;i<10;i++{
//		if i==5 {
//			break
//		}
//		fmt.Println(i)
//	}
//
//}

// 当i=5时，跳过此次for循环（不执行for循环内部的打印语句），继续下一次循环

//func main (){
//	for i:=0;i<10;i++{
//		if i==5{
//			continue
//		}
//		fmt.Println(i)
//	}
//}


// switch
// 简化大量的判断（一个变量和具体的值作比较）

//func main (){
//	finger:=5
//	if finger==1 {
//		fmt.Println("大拇指")
//	}else if finger==2 {
//		fmt.Println("食指")
//	}else {
//		fmt.Println("你猜是哪个指头")
//	}
//}
//
//func main(){
//	finger :=5
//	switch finger{
//	case 1:
//		fmt.Println("大拇指")
//	case 2:
//		fmt.Println("食指")
//	default:
//		fmt.Println("你猜")
//
//
//	}
//
//
//}

// 跳出多层for循环

//var flag=true
//func main (){
//	for i:=0 ;i<4 ;i++{
//		for j:='A';j<'C';j++{
//			if j=='B'{
//				flag=true
//				break
//			}
//			fmt.Printf("%d-%c\n",i,j)
//		}
//		if flag {
//			break
//		}
//	}
//}


// goto+label实现跳出多层for循环

//func main (){
//	for i:=0 ;i<4 ;i++{
//		for j:='A';j<'C';j++{
//			if j=='B' {
//				goto flag
//			}
//			fmt.Printf("%d-%c\n",i,j)
//		}
//
//	}
//flag: fmt.Println("game over")
//
//
//}


// 运算符
//func main (){
//	var a=5
//	var b=7
//	fmt.Println(a+b,a-b,a*b,a/b,a%b)
//	fmt.Println(ab,a-b,a*b,a/b,a%b)
//

// func f1(x int, y int ) (ret int ){
// 	ret=x+y
// 	return
// }
//
//
//
//func f2() (ret int ){
//	ret=5
//	return
//}
//
//
//func ff(x func(x int,y int) (ret int)) (func() (ret int)){
//	return f2
//}
//
//
//func  main(){
//
//	 f7:=ff(f1)
//	fmt.Printf("%T\n",f7)
//
//	fmt.Printf("%T",f2)
//}
//

//func adder(x int) func(int) int{
//	return func(y int) int{
//	x+=y
//	return x
//	}
//}
//
//
//
//
//
//func main(){
//	ret:=adder(100)
//	ret2:=ret(200)
//	fmt.Println(ret2)
//}





