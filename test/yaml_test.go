package kiris

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestConfGet(t *testing.T) {
	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
	fmt.Println("Load Yaml Config")
	yml := NewYaml([]byte(`
default:
  admin:
    name: admin
    host: 10.1.1.1
    port: 22
    user: admin
    password: admin
    key: ""
    options: {}
  bb:
    name: bb
    host: 10.2.2.2
    port: 22
    user: ad
    password: ad
    key: ""
    options: {}
  cc:
    dd:
      ee:
        name: cc.dd
        host: 11.11.11.11
        port: 22
        user: a
        password: b
        key: ""
        options: {}
`))
	yml.Set("test.Name", "kee")
	yml.Set("test.Location", "PRC")
	yml.SaveAs("./data/test.yml")

	str, _ := yml.SaveToString()
	fmt.Println("Save To String: ", string(str))

	cnf := NewYamlLoad("./data/test.yml")

	fmt.Println("BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB")
	fmt.Println(cnf.Get(""))

	cnf.Set("name", "test")
	cnf.Set("product.location", "PRC")
	fmt.Println("Get Name (=test): ", cnf.Get("name"))

	fmt.Println("Get Location(=UTC): ", cnf.Get("location", "UTC").(string))

	fmt.Println("Get Product.Location(!=UTC): ", cnf.Get("product.location", "UTC"))

	fmt.Println("CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC")

	cnf.Set("name", rand.Intn(100))
	fmt.Println("Get Name(updated !=test): ", cnf.Get("name").(int))

	cnf.Set("product.location", rand.Intn(100))
	fmt.Println("Get product.location(updated is Inteval): ", cnf.Get("product.location", "UTC"))

	cnf.Set("product.redis.host", "localhost")
	fmt.Println("Get product.redis.host(updated => localhost): ", cnf.Get("product.redis.host", "127.0.0.1"))
	fmt.Println("DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD")

	fmt.Println(cnf.Get(""))
	cnf.Set("product.port", 3000)
	datas := cnf.Get("product") //.(map[string]interface{})
	fmt.Println("DEEPCOPY: ", datas)

	//for k, v := range datas {
	//fmt.Println(k, " => ", v)
	//}
	fmt.Println("EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE")

	cnf.Set("stage", DeepCopy(cnf.Get("product")))
	cnf.Set("stage.debug", true)
	cnf.Set("stage.env", "_stage_")
	cnf.Set("stage.port", 3000)
	cnf.Set("disk.io", 65539)
	networks := []string{"lo", "wlan", "eth-0"}
	cnf.Set("networks", networks)
	fmt.Println("Get Stage.debug: ", cnf.Get("stage.debug"))
	fmt.Println("Get Product.debug: ", cnf.Get("product.debug"))

	fmt.Println("DELDELDELDELDELDELDELDELDELDELDELDELDELDELDELDELDELDELDELDELDELDEL")
	cnf.Del("product.location")
	cnf.Del("stage.port")
	cnf.Del("default.cc.dd.ee")

	cnf.SaveAs("./data/test2.yml")
	//cnf.Save()

}
