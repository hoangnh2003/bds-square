package main

import (
	"bds-square-backend/internal/database"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mdb, err := sql.Open("mysql", "root:root1234@tcp(127.0.0.1:33306)/shopdevgo")
	if err != nil {
		panic(err)
	}
	defer mdb.Close()

	// execution
	dao := database.New(mdb)

	// get List
	ctx := context.Background()

	// insert shop
	err = dao.CreateShop(ctx, "shop 2")
	if err != nil {
		log.Fatal(err)
	}
	shopList, err := dao.GetShops(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(shopList)
}
