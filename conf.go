package justproxy

import (
	"encoding/json"
	"io/ioutil"
)

type ProxyItem struct {
	Src  string `json:"src"`
	Dest string `json:"dest"`
}

type Config struct {
	Proxys []*ProxyItem `json:"proxys"`
}

func LoadConfig(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	conf := new(Config)
	err = json.Unmarshal(data, conf)
	return conf, err
}
