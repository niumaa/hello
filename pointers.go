//指针用于间接引用和重定向,go没有指针运算
package main
import "fmt"
func main() {
	i, j :=42, 2701
	p :=&i //pointer指向i地址
	fmt.Println(*p)
	*p = 21 //地址的值进行重新改写
	fmt.Println(i)
	p =&j
	*p = *p /37
	fmt.Println(j)
}
