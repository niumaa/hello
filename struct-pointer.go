//指向结构体的指针P，通过（*p).X来访问字段X.允许使用隐形间接引用，直接p.X
package main 
import "fmt"
type Vertex struct {
	X int
	Y int
}
func main() {
	v :=Vertex{1, 2}
	p :=&v
	p.X = 1e9
	fmt.Println(v)
}
