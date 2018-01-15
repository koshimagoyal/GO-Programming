package main
import "fmt"
func main() {
	var a,b,c int32
	fmt.Printf("\n enter value of a and b")
	fmt.Scanf("%d %d",&a,&b)
	c=a+b
	fmt.Printf("\n Addition of a and b is %d",c)
	c=a-b
	fmt.Printf("\n Subtraction of a and b is %d",c)
}