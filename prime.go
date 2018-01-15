package main
import "fmt"
func main() {
	var i,no,count int
	fmt.Printf("enter the last number till you wnt to print prime numbers \n")
	fmt.Scanf("%d",&no)
	fmt.Printf("the prime numbers are \n")
	for i=1;i<=no;i++ {
		count=0
		for j:=1;j<=i;j++ {
		if(i%j==0) {
			count++
		}
		}
		if(count==2) {
			fmt.Printf("%d\n",i)
		
		}
	}
}