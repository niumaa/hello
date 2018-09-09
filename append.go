//向切片追加元素，内建的append函数。
//func append(s []T,vs ...T) []T
//append的第一个参数s是一个元素类型是T的切片，其余类型为T的值会追加到切片的尾部.
//当s的底层数组太小，不足以容纳所有给定的值时，会分配一个更大的数组。返回的切片会指向这个新分配的数组
package main
import "fmt"
func main() {
	var s []int
	printSlice(s)
	//append works on nil slices
	s =append(s, 0)
	printSlice(s)
	//The slice grows as needed
	s =  append(s, 1)
	printSlice(s)
	//we can add more than one element at a time
	s = append(s, 2, 3,  4)
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
