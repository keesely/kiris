/*************************************************************************
   > File Name: conf.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2019.10.27
************************************************************************/
package kiris

import (
	"log"
)

type AppConf interface {
	Load(filename string)
	Get(key string, _def ...interface{}) interface{}
	Set(key string, value interface{}) error
	Save() error
	SaveAs(filename string) error
}

var adapters = make(map[string]AppConf)

func Register(cnfType string, adapter AppConf) {
	if nil == adapter {
		panic("Config Register Error: adapter is nil")
	}

	if _, ok := adapters[cnfType]; ok {
		panic("Config Register Error: adapter (" + cnfType + ") is registed")
	}

	adapters[cnfType] = adapter
}

func LoadConfig(cnfType string, filename string) AppConf {
	if adapter, ok := adapters[cnfType]; ok {
		return adapter
	}
	log.Fatal("Config Adapter Error: unknown adapter ", cnfType)
	return nil
}
