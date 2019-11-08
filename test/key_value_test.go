package kiris

import (
	"fmt"
	"testing"
)

var KvMaps = NewK_VMaps()
var saveFile = "./data/dump.kvmap"

func TestSet(t *testing.T) {
	name := KvMaps.Set("Name", "Kee")
	sex := KvMaps.Set("Sex", 1)
	age := KvMaps.Set("Age", 28)
	addr := KvMaps.Set("Addr", "China, Fujian, Xiamen")
	name.SetAttr("next", sex)
	sex.SetAttr("next", age)
	age.SetAttr("next", addr)
	addr.SetAttr("next", nil)

	fmt.Println(name, "\n", sex, "\n", age, "\n", addr, "\n", sex.GetAttr("next"))
	fmt.Println(StrPad("", "->", 100, KIRIS_STR_PAD_RIGHT))
}

func TestPrint(t *testing.T) {
	KvMaps.Del("Age")
	KvMaps.Set("Name", "Marco")

	KvMaps.Print()

	fmt.Printf("Print Age Value (Defalut 22): %d \n", KvMaps.GetValue("Age", 22))
	KvMaps.Set("Age", 26)
	fmt.Printf("Print Age Value (= 26): %d \n", KvMaps.GetValue("Age"))
	e := KvMaps.Save(saveFile)
	fmt.Println("Key-Value Save: ", e)
}

func TestLoad(t *testing.T) {
	fmt.Println(StrPad("", "=", 100, KIRIS_STR_PAD_RIGHT))
	fmt.Printf("+%50s%-48s+\n", "TEST LOAD DB FILE", "")
	kMap := NewK_VMaps()
	kMap.Load(saveFile)
	kMap.Print()

	fmt.Println(kMap.List())
}
