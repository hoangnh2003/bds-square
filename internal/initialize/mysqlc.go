package initialize

import (
	"bds-square-backend/global"
	"database/sql"
	"fmt"
	"time"

	"go.uber.org/zap"
)

func InitMysqlC() {
	m := global.Config.Mysql

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanicC(err, "InitMysql initialization error")

	global.Logger.Info("Initializing MySQL Successfully")
	global.Mdbc = db

	SetPoolC()

}

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func SetPoolC() {
	m := global.Config.Mysql
	sqlDb := global.Mdbc
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Second * time.Duration(m.ConnMaxLifetime))
}
