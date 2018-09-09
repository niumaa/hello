//for循环的range形式可遍历切片或映射
//当使用for循环遍历切片时， 每次迭代返回2个值，第一个值为当前元素的下标，第二个该下标对应元素的一个副本。
package main
import "fmt"
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
func main() {
	for i, v :=range pow{
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
