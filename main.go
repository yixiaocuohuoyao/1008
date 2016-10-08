package main

import "fmt"

func main(){
	i:=1
	j:=&i
	fmt.Print(addONE(j))
}

func addONE(i *int)int{
	*i+=1
	return *i
}