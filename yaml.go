package kiris

import (
	"errors"
	"fmt"
	yaml "gopkg.in/yaml.v2"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

type Yaml struct {
	originData yaml.MapSlice
	data       map[string]interface{}
	sync.RWMutex
}

var saveFile string

func NewYamlLoad(filename string) *Yaml {
	saveFile = RealPath(filename)
	f, e := FileGetContents(saveFile)
	if e != nil {
		log.Fatal("Get Yaml Error: ", e)
	}

	return NewYaml(f)
}

func NewYaml(yc []byte) *Yaml {
	cnf := make(map[string]interface{})
	e := yaml.Unmarshal(yc, cnf)
	data := &yaml.MapSlice{}
	e = yaml.Unmarshal(yc, data)
	if e != nil {
		log.Fatal("Unmarshal Yaml: ", e)
	}

	cnf = FormatValueMaps(cnf)
	return &Yaml{data: cnf, originData: *data}
}

func FormatValueMaps(m map[string]interface{}) map[string]interface{} {
	// 获取keys
	for k, v := range m {
		switch value := v.(type) {
		case string:
			m[k] = ExpandValueEnv(value)
		case map[string]interface{}:
			m[k] = FormatValueMaps(value)
		case map[interface{}]interface{}:
			_value := make(map[string]interface{})
			for _k, _v := range m[k].(map[interface{}]interface{}) {
				_value[_k.(string)] = _v
			}
			m[k] = _value
			m[k] = FormatValueMaps(_value)
		case map[string]string:
			for k2, v2 := range value {
				value[k2] = ExpandValueEnv(v2)
			}
			m[k] = value
		}
	}
	return m
}

// Convert `$(ENV)` || `$(ENV||defaultValue)` || `$(ENV||)`
// Return the env value || if env is nil return defaultValue || env is nil return ""
func ExpandValueEnv(value string) string {
	rVal := value

	dVal := ""
	regx := regexp.MustCompile(`(?U)\$\{.+\}`)

	if x := regx.FindAllString(rVal, -1); len(x) > 0 {
		for _, v := range x {
			vL := len(v)
			if vL < 3 {
				continue
			}
			if key := v[2 : vL-1]; len(key) > 0 {
				dValal := ""
				dValalIndex := strings.Index(v, "||")
				if dValalIndex > 0 {
					key = v[2:dValalIndex]
					dValal = v[dValalIndex+2 : vL-1]
				}

				eVal := GetEnv(key, dValal).(string)
				rVal = strings.Replace(rVal, v, eVal, -1)
			}
		}
	}

	if rVal == "" {
		rVal = dVal
	}

	return Ternary(rVal == "", dVal, rVal).(string)
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

	_data := DeepCopy(this.data).(map[string]interface{})
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
				this.originData = getSaveData(this.originData, this.data)
			} else {
				vv := make(map[string]interface{})
				_data[k] = vv
				_data = _data[k].(map[string]interface{})
			}
		}
	}
	return nil
}

func (this *Yaml) Del(key string) error {
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
					if i == len(keys)-1 {
						delete(_data, k)
					} else {
						_data = v.(map[string]interface{})
					}
				}
			default:
				{
					delete(_data, k)
				}
			}
		}
	}
	return nil
}

func (this *Yaml) Save() error {
	return this.SaveAs(saveFile)
}

func (this *Yaml) SaveAs(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	ye := yaml.NewEncoder(f)
	//err = ye.Encode(this.data)
	this.originData = getSaveData(this.originData, this.data)
	err = ye.Encode(this.originData)
	return err
}

func (this *Yaml) SaveToString() ([]byte, error) {
	this.originData = getSaveData(this.originData, this.data)
	return yaml.Marshal(this.originData)
}

func getSaveData(originData yaml.MapSlice, data map[string]interface{}) yaml.MapSlice {
	//var _data = yaml.MapSlice{}
	for k, v := range data {
		var flag bool
		for i := 0; i < len(originData); i++ {
			m := originData[i]
			if _, ok := data[m.Key.(string)]; !ok {
				originData = append(originData[:i], originData[i+1:]...)
				continue
			}
			if m.Key == k {
				flag = true
				switch value := v.(type) {
				case map[string]interface{}:
					originData[i].Value = getSaveData(originData[i].Value.(yaml.MapSlice), value)
				default:
					originData[i].Value = value
				}
			}
		}

		if true != flag {
			if value, ok := yaml.Marshal(v); ok == nil {
				m := &yaml.MapSlice{}
				_ = yaml.Unmarshal(value, m)
				originData = append(originData, yaml.MapItem{Key: k, Value: *m})
			}
		}
	}
	return originData
}
