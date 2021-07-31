package pkg

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

type JsonContent interface {
	toBytes() []byte
}

// 解析 json:
func ReadJsonFile(filename string, dist interface{}) error {

	// read file:
	input, _ := ioutil.ReadFile(filename)

	// parse:
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(input, &dist)
}

//
func ConvertToJsonBytes(input interface{}) ([]byte, error) {
	// convert:
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Marshal(input)
}

// write:
func WriteJsonFile(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}
