package pkg

import (
	"testing"
)

func TestReadJsonFile(t *testing.T) {
	// from:
	from := "../../data/test1.json"

	// result:
	dist := make(map[string]interface{})

	// parse:
	err := ReadJsonFile(from, &dist)

	t.Logf("result: %v,  %+v\n", err, dist)

	// iter
	for k, v := range dist {
		t.Logf("\tkv: k= %v, v= %+v\n", k, v)
	}

}
