//切片通过2个下标来界定， 即 一个上届和一个下界，a[low: high], 
package main
import "fmt"
func main() {
	primes :=[6]int{2, 3 , 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
}
