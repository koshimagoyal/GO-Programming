package main
import "fmt"
func main() {
	var a,b,c=3,4.2,"foo"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("a is type of %T\n",a)
	fmt.Printf("b is type of %T\n",b)
	fmt.Printf("c is type of %T\n",c)
}