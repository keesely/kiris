/*************************************************************************
   > File Name: conf.go
   > Author: Kee
   > Created Time: 2019.10.25
************************************************************************/
package kiris

import (
	"github.com/astaxie/beego/config"
	"log"
)

import _ "github.com/astaxie/beego/config/yaml"

type AppConf struct {
	Conf config.Configer
}

func LoadConfig(_type string, _conf string) *AppConf {
	cnf, err := config.NewConfig(_type, _conf)
	if err != nil {
		log.Fatal("Get Config Error: ", err)
		return nil
	}
	return &AppConf{cnf}
}

func (this *AppConf) Get(key string, _def ...interface{}) interface{} {
	var def interface{}
	if nil == _def || nil == _def[0] {
		def = nil
	} else {
		def = _def[0]
	}

	val, err := this.Conf.DIY(key)
	if err != nil {
		return def
	}
	return Ternary(nil == val, def, val)
}
