package main
import "fmt"
func main(){
	fmt.Println("counting")
	for i :=0; i < 10; i++ {
		defer  fmt.Println(i) //推迟的函数调用会被压入一个zhan中，按从后进先出的顺序
	}
}
