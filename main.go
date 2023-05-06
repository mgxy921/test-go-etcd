package main

import (
	"fmt"
	"reflect"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

const (
	etcdAddr = "localhost:2379"
)

func main() {
	// 创建 ETCD Registry
	etcdRegistry := etcd.NewRegistry(
		registry.Addrs(etcdAddr),
	)

	// // 创建 Micro Service
	// MicroService := micro.NewService(
	// 	micro.Registry(etcdRegistry),
	// )

	// // 运行服务
	// if err := MicroService.Run(); err != nil {
	// 	fmt.Println(err)
	// }

	// 获取服务列表
	services, err := etcdRegistry.ListServices()
	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.TypeOf(services))

	for _, service := range services {
		fmt.Println(reflect.TypeOf(service))
		for _, node := range service.Nodes {
			fmt.Println(node.Address)
			fmt.Println(reflect.TypeOf(node))
		}
	}
}
