package main
import "fmt"
func main() {
	var b int = 13
	var a int
	number := [6]int{1,2,3,4}
	for a=0; a<10; a++ {
		fmt.Printf("value of a is %d\n",a)
	}
	for a<b {
		a++
		fmt.Printf("value of a is %d\n",a)
	}
	for i,x := range number {
		fmt.Printf("value of a is %d at i is %d\n",x,i)
	}
}