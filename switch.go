package  main
import "fmt"
func main() {
	var a,b,c,ch int32
	var d float64
	fmt.Printf("1.Addition\n")
	fmt.Printf("2.Subtraction\n")
	fmt.Printf("3.Multiplication\n")
	fmt.Printf("4.Division\n")
	fmt.Printf("5.Modulus\n")
	ch=4
	fmt.Printf("enter the value of a and b\n")
	fmt.Scanf("%d %d",&a,&b)
	switch ch {
		case 1 : c=a+b
				 fmt.Printf("The addition of a and b is %d\n",c)
		case 2 : c=a-b
				 fmt.Printf("The subtraction of a and b is %d\n",c)
		case 3 : d=(float64)(a/b)
				 fmt.Printf("The division of a and b is %f\n",d)
		case 4 : c=a%b
				 fmt.Printf("The modulus of a and b is %d\n",c)
		default : fmt.Printf("Wrong Input\n")
	}
}