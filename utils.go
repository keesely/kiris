/*************************************************************************
   > File Name: utils.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2019.10.25
************************************************************************/
package kiris

func Ternary(cond bool, Tval, Fval interface{}) interface{} {
	if cond {
		return Tval
	}
	return Fval
}
