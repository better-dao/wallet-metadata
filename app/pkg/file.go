package pkg

import (
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
)

// 解析 json:
func ReadJsonFile(filename string, dist interface{}) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read data.
	input, _ := ioutil.ReadAll(f)

	// parse:
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	return json.Unmarshal(input, &dist)
}
