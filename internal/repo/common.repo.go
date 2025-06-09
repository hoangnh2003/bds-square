package repo

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Ctx context.Context
)

func init() {
	Ctx = context.Background()
}
