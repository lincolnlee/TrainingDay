package main

import (
	"fmt"
)

//初始化和主函数写在这个文件中
var utility Utility = Utility{}

func main() {
	fmt.Println("Start...")

	appConfig := AppConfig{}
	appConfig.GetCfg()

	machine := Machine{AppConfig: &appConfig}

	machine.StartHTTPService()
	machine.StartTCPService()
	machine.DealSendCmd()
}
