package middleware

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func RegisterService(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	// 生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://192.168.1.4:8021/health",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	// 生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Address = address
	registration.Port = port
	registration.Name = name
	registration.Tags = tags
	registration.ID = id
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		panic(err)
	}

	return err
}

func AllService() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().Services()

	if err != nil {
		panic(err)
	}

	for k := range services {
		fmt.Println(k)
	}
}

func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	services, err := client.Agent().ServicesWithFilter(`Service == "shop"`)

	if err != nil {
		panic(err)
	}

	for k := range services {
		fmt.Println(k)
	}
}
