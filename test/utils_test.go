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

func TestStrPad(t *testing.T) {
	s1 := StrPad("Hello World (Left)", "你好世界", 100, KIRIS_STR_PAD_LEFT)
	s2 := StrPad("你好世界 (Right)", "=", 100, KIRIS_STR_PAD_RIGHT)
	s3 := StrPad(" Hello World ", "Both", 100, KIRIS_STR_PAD_BOTH)
	fmt.Printf("strPad > %s, LEN: %d \n", s1, StrCount(s1))
	fmt.Printf("strPad > %s, LEN: %d \n", s2, StrCount(s2))
	fmt.Printf("strPad > %s, LEN: %d \n", s3, StrCount(s3))
}
