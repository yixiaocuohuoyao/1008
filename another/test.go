package main

import "fmt"

func main() {
  	s:=fmt.Printf
	x:=[]byte("a s d a s d a s d ")
	ss:=[]uint8(x)
	s("%x\n", ss)
}
