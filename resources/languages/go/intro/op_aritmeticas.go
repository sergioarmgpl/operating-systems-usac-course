/*
+	addition
-	subtraction
*	multiplication
/	division
%	remainder
*/
package main

import "fmt"

func main() {
	a := 1
	b := 2
	c1 := a + b
	c2 := a - b
	c3 := a * b
	c4 := a / b
	fmt.Printf("%d + %d = %d\n", a, b, c1)
	fmt.Printf("%d - %d = %d\n", a, b, c2)
	fmt.Printf("%d * %d = %d\n", a, b, c3)
	fmt.Printf("%d / %d = %d\n", a, b, c4)
}
