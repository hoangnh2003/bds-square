package initialize

import (
	"bds-square-backend/global"
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Run() *gin.Engine {
	LoadConfig()
	fmt.Println("Loading config mysql", global.Config.Mysql.Dbname)
	InitLogger()
	global.Logger.Info("config Log ok!!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	return r
}
