package main
import "fmt"
func main() {
	var i,m,n,count int
	fmt.Printf("enter the m and n\n")
	fmt.Scanf("%d %d",&m,&n)
	fmt.Printf("the prime numbers are\n")
	for i=m;i<=n;i++ {
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
	