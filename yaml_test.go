/*************************************************************************
   > File Name: conf_test.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2019.10.25
************************************************************************/
package kiris

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestConfGet(t *testing.T) {
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fmt.Println("Load Yaml Config")
	cnf := NewYamlLoad("./conf/test.yml")

	fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	fmt.Println(cnf.Get(""))

	cnf.Set("name", "test")
	cnf.Set("product.location", "PRC")
	fmt.Println("Get Name (=test): ", cnf.Get("name"))

	fmt.Println("Get Location(=UTC): ", cnf.Get("location", "UTC"))

	fmt.Println("Get Product.Location(!=UTC): ", cnf.Get("product.location", "UTC"))

	fmt.Println("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")

	cnf.Set("name", rand.Intn(100))
	fmt.Println("Get Name(updated !=test): ", cnf.Get("name"))

	cnf.Set("product.location", rand.Intn(100))
	fmt.Println("Get product.location(updated is Inteval): ", cnf.Get("product.location", "UTC"))

	cnf.Set("product.redis.host", "localhost")
	fmt.Println("Get product.redis.host(updated => localhost): ", cnf.Get("product.redis.host", "127.0.0.1"))
	fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")

	fmt.Println(cnf.Get(""))
	cnf.Set("product.port", 3000)
	datas := cnf.Get("product").(map[string]interface{})
	fmt.Println(datas)

	for k, v := range datas {
		fmt.Println(k, " => ", v)
	}
	fmt.Println("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")

	cnf.Set("stage", DeepCopy(cnf.Get("product")))
	cnf.Set("stage.debug", true)
	cnf.Set("stage.env", "stage")
	fmt.Println("Get Stage.debug: ", cnf.Get("stage.debug"))
	fmt.Println("Get Product.debug: ", cnf.Get("product.debug"))
	//
	cnf.SaveAs("./conf/aaa.yml")
	cnf.Save()

}
