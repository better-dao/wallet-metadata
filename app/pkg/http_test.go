package pkg

import (
	"testing"
)

func TestHttpGet(t *testing.T) {

	HttpGet(bscTokenUrl)

	//HttpGet(ethTokenUrl)
}

func TestGetBscTokenMeta(t *testing.T) {
	GetBscTokenMeta()
}
