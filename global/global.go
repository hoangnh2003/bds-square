package global

import (
	"bds-square-backend/pkg/logger"
	"bds-square-backend/pkg/setting"
	"database/sql"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
	Mdbc   *sql.DB
)
