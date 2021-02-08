package main

import "fmt"

func main()  {
	fmt.Println("hello World")
	for i:=0;i<10;i++{
		fmt.Print(i," ")
	}
	fmt.Println()
	for i:=0;i<10;i++{
		if i%2==0 {
			fmt.Print(i," ")
		}
	}
	fmt.Println()

	fmt.Println("2*2 is ", 2*2)

}