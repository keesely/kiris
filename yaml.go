/*************************************************************************
   > File Name: yaml.go
   > Author: Kee
   > Mail: chinboy2012@gmail.com
   > Created Time: 2019.10.25
************************************************************************/
package kiris

import (
	"errors"
	"fmt"
	beegoYaml "github.com/astaxie/beego/config/yaml"
	yaml "gopkg.in/yaml.v2"
	"log"
	"os"
	"strings"
	"sync"
)

type Yaml struct {
	data     map[string]interface{}
	filename string
	sync.RWMutex
}

func NewYamlLoad(filename string) *Yaml {
	cnf, err := beegoYaml.ReadYmlReader(filename)
	if err != nil {
		log.Fatal("Get Yaml Error: ", err)
	}
	return &Yaml{data: cnf, filename: filename}
	return nil
}

func (this *Yaml) Get(key string, _def ...interface{}) interface{} {
	var def interface{}
	if nil == _def || nil == _def[0] {
		def = nil
	} else {
		def = _def[0]
	}

	val, err := this.getData(key)
	if err != nil {
		return def
	}
	return Ternary(nil == val, def, val)
}

func (this *Yaml) getData(key string) (interface{}, error) {
	if len(key) == 0 {
		return this.data, nil
	}

	this.RLock()
	defer this.RUnlock()

	keys := strings.Split(key, ".")

	_data := this.data
	for idx, k := range keys {
		if v, ok := _data[k]; ok {
			switch v.(type) {
			case map[string]interface{}:
				{
					_data = v.(map[string]interface{})
					if idx == len(keys)-1 {
						return _data, nil
					}
				}
			default:
				{
					return v, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("not exists key %q", key)
}

func (this *Yaml) Set(key string, value interface{}) error {
	if len(key) == 0 {
		return errors.New("key is empty")
	}

	this.Lock()
	defer this.Unlock()

	keys := strings.Split(key, ".")
	var _data = this.data
	for i, k := range keys {
		if v, ok := _data[k]; ok {
			switch v.(type) {
			case map[string]interface{}:
				{
					_data[k] = v
					if i == len(keys)-1 {
						v = value
						_data[k] = v
					}
					_data = v.(map[string]interface{})
				}
			default:
				{
					v = value
					_data[k] = v
				}
			}
		} else {
			_data[k] = make(map[string]interface{})
			if i == len(keys)-1 {
				_data[k] = value
			} else {
				vv := make(map[string]interface{})
				_data[k] = vv
				_data = _data[k].(map[string]interface{})
			}
		}
	}
	return nil
}

func (this *Yaml) Save() error {
	return this.SaveAs(this.filename)
}

func (this *Yaml) SaveAs(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	ye := yaml.NewEncoder(f)
	err = ye.Encode(this.data)
	return err
}
