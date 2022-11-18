package conf

import (
	"context"
	"github.com/pelletier/go-toml/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/ioutil"
	"os"
)

type Conf struct {
	Version     string
	LangVersion string
	SCPath      string
}

var Config Conf

func InitConf() {
	bs, err := ioutil.ReadFile("./conf.toml")
	if err != nil {
		panic(err)
	}
	if err := toml.Unmarshal(bs, &Config); err != nil {
		panic(err)
	}
}

func Save() {
	data, err := toml.Marshal(Config)
	if err != nil {
		runtime.LogPrint(context.Background(), err.Error())
		return
	}
	if err := ioutil.WriteFile("conf.toml", data, os.ModePerm); err != nil {
		runtime.LogPrint(context.Background(), err.Error())
		return
	}
	return
}
