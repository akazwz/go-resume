package i18n

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"strconv"
)

// Trans
// params:
// key
// lang
// path
// type

type Handler interface {
	do()
	SetNext(handler Handler) Handler
	Run()
}

type Next struct {
	nextHandler Handler
}

func (n *Next) SetNext(handler Handler) Handler {
	n.nextHandler = handler
	return handler
}

func (n *Next) Run() {
	if n.nextHandler != nil {
		n.nextHandler.do()
		n.nextHandler.Run()
	}
}

type RootHandler struct {
	Next
}

func (h RootHandler) do() {

}

type OneHandler struct {
	Next
}

func (o OneHandler) do() {
	fmt.Println("handle one do something")
}

type TwoHandle struct {
	Next
}

func (t TwoHandle) do() {
	fmt.Println("handle two do something")
}

type I18n struct {
	TransFilePath string
	Value         interface{}
}

func (i *I18n) SetTransFilePath(transFilePath string) *I18n {
	i.TransFilePath = transFilePath
	return i
}

func (i *I18n) Trans(arg ...string) *I18n {
	config := viper.New()
	configPath := "./config/i18n"
	configName := "en"
	configType := "yaml"
	switch len(arg) {
	case 1:
		break
	case 2:
		configName = arg[1]
		break
	case 3:
		configName = arg[1]
		configPath = arg[2]
		break
	case 4:
		configName = arg[1]
		configPath = arg[2]
		configType = arg[3]
		break
	}
	config.SetConfigName(configName)
	config.AddConfigPath(configPath)
	config.SetConfigType(configType)
	// 读取配置
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	value := config.Get(arg[0])
	i.Value = value
	return i
}

func (i *I18n) ToStr() string {
	value := i.Value
	var key string
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
