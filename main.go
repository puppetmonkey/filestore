package main

import (
	"filestore/dao/ipfs"
	"filestore/dao/mysql"
	"filestore/dao/redis"
	"filestore/logger"
	"filestore/pkg/snowflake"
	"filestore/router"
	"filestore/setting"

	"fmt"
)

// @title bluebell项目接口文档
// @version 1.0
// @description Go web开发进阶项目实战课程bluebell

// @contact.name liwenzhou
// @contact.url http://www.liwenzhou.com

// @host 127.0.0.1:8084
// @BasePath /api/v1
func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("need config file.eg: bluebell config.yaml")
	// 	return
	// }
	// 加载配置
	// err := setting.Init(os.Args[1])
	err := setting.Init("conf/config.yaml")
	if err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		return
	}
	if err := logger.Init(setting.Conf.LogConfig, setting.Conf.Mode); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	if err := mysql.Init(setting.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer mysql.Close() // 程序退出关闭数据库连接
	if err := redis.Init(setting.Conf.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer redis.Close()

	if err := ipfs.Init("localhost:5001"); err != nil {
		fmt.Printf("Failed to initialize IPFS: %v", err)
	}
	// if err := eth.Init("localhost:8545"); err != nil {
	// 	fmt.Printf("Failed to initialize eth: %v", err)
	// }
	if err := snowflake.Init(setting.Conf.StartTime, setting.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}
	// // 初始化gin框架内置的校验器使用的翻译器
	// if err := controller.InitTrans("zh"); err != nil {
	// 	fmt.Printf("init validator trans failed, err:%v\n", err)
	// 	return
	// }
	// 注册路由
	r := router.SetupRouter(setting.Conf.Mode)
	err = r.Run(fmt.Sprintf(":%d", setting.Conf.Port))

	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
