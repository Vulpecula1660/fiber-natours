package user

import (
	"context"
	dbSQL "database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Vulpecula1660/fiber-natours/model/dto"
	"github.com/Vulpecula1660/fiber-natours/model/postgresql"
)

// Upd : 更新用戶 upd user
// Transaction 為選填
func Upd(ctx context.Context, tx *dbSQL.Tx, input *dto.User) (err error) {
	if input == nil {
		return fmt.Errorf("參數錯誤")
	}

	var dbM *dbSQL.DB

	if tx == nil {
		dbM = postgresql.GetConn(os.Getenv("DATABASE_Name"))
	}

	sql := " UPDATE users SET "
	sql += "    password = ?"
	sql += " WHERE "

	var params []interface{}
	var wheres []string

	params = append(params, input.Password)

	// 會員ID
	if input.ID != "" {
		wheres = append(wheres, " `id` = ? ")
		params = append(params, input.ID)
	}

	// 沒有條件時回傳錯誤
	if len(wheres) == 0 {
		return fmt.Errorf("sql 語法錯誤")
	}

	sql += strings.Join(wheres, " AND ")

	// 執行sql

	if tx == nil {
		_, err = dbM.ExecContext(ctx, sql, params...)
	} else {
		_, err = tx.ExecContext(ctx, sql, params...)
	}
	if err != nil {
		return fmt.Errorf("upd錯誤 error: %v, sql: %v, params: %v ", err, sql, params)
	}

	return err
}
