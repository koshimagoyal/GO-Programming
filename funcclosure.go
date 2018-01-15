package main
import "fmt"
func sequence() func() int{
	i := 0
	return func() int{
		i+=1
		return i
	}
}
func main() {
	number := sequence()
	fmt.Println(number())
	fmt.Println(number())
	fmt.Println(number())
	number1 := sequence()
	fmt.Println(number1())
	fmt.Println(number1())
}
