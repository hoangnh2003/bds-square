package main

import (
	"bds-square-backend/global"
	"bds-square-backend/internal/initialize"
	"bds-square-backend/internal/model"
	"bds-square-backend/internal/po"
	"fmt"
)

func main() {
	initialize.LoadConfig()
	initialize.InitLogger()
	initialize.InitMysql()

	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
		&model.GoCrmUser{},
	)

	if err != nil {
		fmt.Println("Migrate fail:", err)
	} else {
		fmt.Println("Migrate successful!")
	}
}
