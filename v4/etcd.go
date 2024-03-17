/*
 * @Author: mgxy921 mgxy921@gmail.com
 * @Date: 2024-03-17 15:53:37
 * @LastEditors: mgxy921 mgxy921@gmail.com
 * @LastEditTime: 2024-03-17 16:04:10
 * @FilePath: \go-etcd\v4\etcd.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v4

import (
	"fmt"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4/registry"
)

var EtcdReg registry.Registry

func init() {
	EtcdReg = etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", "localhost", "2379")),
	)
}

func ShowEtcdInfo() {
	EtcdReg = etcd.NewRegistry(
		registry.Addrs(fmt.Sprintf("%s:%s", "localhost", "2379")),
	)
	// 获取服务列表
	services, err := EtcdReg.ListServices()
	if err != nil {
		panic(err)
	}
	for _, service := range services {
		fmt.Println("service name:", service.Name)
		fmt.Println("node:")
		for _, node := range service.Nodes {
			fmt.Println("================================")
			// fmt.Println("    ", node.Metadata)
			fmt.Println("    ", "node.Address:", node.Address)

		}
		fmt.Println("****************************************************************")
	}
}
