// Package seeders 存放数据填充文件
package seeders

import "gohub/pkg/seed"

func Initialize() {

	// 触发记载本目录下其他文件中的init方法

	// 指定优先于目录下的其他文件运行
	seed.SetRunOrder([]string{
		"SeederUsersTable",
	})
}
