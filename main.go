package main

import "fmt"

func main() {
	var x int
	var d, c, v bool
	fmt.Scan(&x)
	d = (((x%1000)/100)*((x%100)/10))%2 == 0
	c = ((x/1000)+((x%100)/10))%(x%10) == 0
	v = (d || c) && !(d && c)
	fmt.Println(d, c, v)
}
