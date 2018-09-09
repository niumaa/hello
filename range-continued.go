//可以将下标或值赋予_来忽略它
//若你只需要索引，去掉value的部分即可。
package main
import "fmt"
func main() {
	pow :=make([]int, 10)
	for i  :=range pow {
		pow[i ] = 1 << uint(i) //==2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
