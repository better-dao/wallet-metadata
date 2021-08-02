package pkg

import (
	"testing"
)

func TestHttpGet(t *testing.T) {

	HttpGet(bscTokenMainNetUrl)

	//HttpGet(ethTokenMainNetUrl)
}

func TestGetBscTokenMeta(t *testing.T) {
	GetBscTokenMeta()
}
