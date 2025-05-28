package initialize

import (
	"bds-square-backend/global"
	"bds-square-backend/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}