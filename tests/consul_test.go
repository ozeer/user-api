package tests

import (
	"testing"
	"user-api/middleware"
)

func TestRegisterService(t *testing.T) {
	err := middleware.RegisterService("127.0.0.1", 8021, "user-api", []string{"shop", "api"}, "user-api")

	if err != nil {
		panic(err)
	}
}

func TestAllService(t *testing.T) {
	middleware.AllService()
}

func TestFilterService(t *testing.T) {
	middleware.FilterService()
}
