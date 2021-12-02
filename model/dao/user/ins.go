package user

import (
	"context"
	dbSQL "database/sql"
	"fmt"
	"os"

	"github.com/Vulpecula1660/fiber-natours/model/dto"
	"github.com/Vulpecula1660/fiber-natours/model/postgresql"
)

// Ins : 新增用戶 ins user
// Transaction 為選填
func Ins(ctx context.Context, tx *dbSQL.Tx, input *dto.User) (err error) {
	if input == nil {
		return fmt.Errorf("參數錯誤")
	}

	var dbM *dbSQL.DB

	if tx == nil {
		dbM = postgresql.GetConn(os.Getenv("DATABASE_Name"))
	}

	sql := " INSERT INTO user ("
	sql += "    id,"
	sql += "    account,"
	sql += "    password"
	sql += " )"
	sql += " VALUES "
	sql += " ( $1, $2, $3)"

	var params []interface{}

	params = append(params, input.ID)
	params = append(params, input.Account)
	params = append(params, input.Password)

	// 執行sql

	if tx == nil {
		_, err = dbM.ExecContext(ctx, sql, params...)
	} else {
		_, err = tx.ExecContext(ctx, sql, params...)
	}
	if err != nil {
		return fmt.Errorf("ins錯誤 error: %v, sql: %v, params: %v ", err, sql, params)
	}

	return err
}
