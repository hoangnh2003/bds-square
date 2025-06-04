package repo

import (
	"bds-square-backend/internal/database"
	"context"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Ctx context.Context
	Dao *database.Queries
	Db  *sql.DB
)

func init() {
	Ctx = context.Background()

	Db, Dao = connectDb()

	if Db == nil || Dao == nil {
		log.Fatal("Không thể khởi tạo DB hoặc DAO")
	}
}

func connectDb() (*sql.DB, *database.Queries) {
	mdb, err := sql.Open("mysql", "root:root1234@tcp(mysql:3306)/shopdevgo")
	if err != nil {
		log.Fatal("Kết nối DB thất bại:", err)
	}

	if err := mdb.Ping(); err != nil {
		log.Fatal("Không thể kết nối tới MySQL:", err)
	}

	return mdb, database.New(mdb)
}
