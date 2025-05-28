package initialize

import (
	"bds-square-backend/global"
	"fmt"

	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Println("Loading config mysql", global.Config.Mysql.Dbname)
	InitLogger()
	global.Logger.Info("config Log ok!!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
