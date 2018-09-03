package main
import "fmt"
func main() {
	defer fmt.Println("world")    //defer语句会将函数推迟到外层函数返回后执
	fmt.Println("hello")
}
