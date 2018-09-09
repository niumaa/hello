//make函数会分配一个元素为零值的数组并返回一个引用了他的切片
//a := make([]int, 5) //len(a)=5
//要指定它的容量，需向make传入第三个参数
//b :=make([]int, 0, 5) //len(b)=0, cap(b)=5
//b = b[: cap(b)] //len(b)=5, cap(b)=5
//b = b[1:] //len()=4, cap(b)=4

package  main
import "fmt"
func main() {
	a := make([]int,  5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
	
	c := b[:2]
        printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)
}
func printSlice(s string, x  []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",s , len(x), cap(x), x)
}
