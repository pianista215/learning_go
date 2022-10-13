package main

import "fmt"

func vals() (int,int){
	return 3,7
}

func complex_vals() (string,int){
	return "hola",6
}

func main(){
	a,b := vals()
	fmt.Println(a)
	fmt.Println(b)

	c,d := complex_vals()
	fmt.Println(c)
	fmt.Println(d)

	_,e := vals()
	fmt.Println(e)
}