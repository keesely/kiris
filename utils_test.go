/*************************************************************************
   > File Name: utils_test.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2019.10.25
************************************************************************/
package kiris

import (
	"fmt"
	"testing"
)

func TestTernary(t *testing.T) {
	a := 1
	b := 2
	fmt.Println(Ternary(a > b, "A > B", "B > A"))

	valA := "value A is not Empty"
	var valB interface{}
	fmt.Println(Ternary(valA != "", valA, "Value A is NULL"))

	valB = nil
	fmt.Println(Ternary(valB != nil, valB, "Value B is NULL"))

	valC := false
	fmt.Println(Ternary(valC, valA, Ternary(valA != valB, "Hello world", "Hi Kee")))
	valC = true
	fmt.Println(Ternary(valC, valA, Ternary(valA != valB, "Hello world", "Hi Kee")))
}
