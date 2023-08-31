package initialize

import (
	"testing"
	"user-api/global"
)

func TestInitLogger(t *testing.T) {

	InitLogger()
	global.Log.Info("hello")
}
