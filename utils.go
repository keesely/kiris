/*************************************************************************
> File Name: utils.go
> Author: Kee
> Mail: chinboy2012@gmail.com
> Created Time: 2019.10.25
************************************************************************/
package kiris

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func Ternary(cond bool, Tval, Fval interface{}) interface{} {
	if cond {
		return Tval
	}
	return Fval
}

func DeepCopy(value interface{}) interface{} {
	if valueMap, ok := value.(map[string]interface{}); ok {
		newMap := make(map[string]interface{})
		for k, v := range valueMap {
			newMap[k] = DeepCopy(v)
		}

		return newMap
	} else if valueSlice, ok := value.([]interface{}); ok {
		newSlice := make([]interface{}, len(valueSlice))
		for k, v := range valueSlice {
			newSlice[k] = DeepCopy(v)
		}

		return newSlice
	}

	return value
}

func GetEnv(key string, def ...interface{}) interface{} {
	var _def interface{}
	if def != nil && def[0] != nil {
		_def = def[0]
	}

	val := os.Getenv(key)
	return Ternary(val != "", val, _def)
}

func Typeof(value interface{}) string {
	return fmt.Sprintf("%T", value)
}

func Rand(x, y int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(y-x+1) + x
}
