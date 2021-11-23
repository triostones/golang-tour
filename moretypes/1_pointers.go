package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21
	fmt.Println(i) // see the new value of i
	p = &j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
