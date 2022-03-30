package user

import (
	"context"
	dbSQL "database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/Vulpecula1660/fiber-natours/model/postgresql"
)

// GetInput :
type GetInput struct {
	Account string
	UserID  string
}

// GetCount : 取得 user
// Transaction 為選填
func GetCount(ctx context.Context, input *GetInput, tx *dbSQL.Tx) (ret int64, err error) {
	var dbS *dbSQL.DB

	if tx == nil {
		dbS = postgresql.GetConn(os.Getenv("DATABASE_Name"))
	}

	sql := " SELECT "
	sql += "    COUNT(*) "
	sql += " FROM `user` "
	sql += " WHERE "

	var params []interface{}
	var wheres []string

	// 帳號
	if input.Account != "" {
		wheres = append(wheres, " `account` = ? ")
		params = append(params, input.Account)
	}

	// 沒有條件時回傳錯誤
	if len(wheres) == 0 {
		return 0, fmt.Errorf("sql 語法錯誤")
	}

	sql += strings.Join(wheres, " AND ")

	var rows *dbSQL.Rows

	if tx == nil {
		rows, err = dbS.QueryContext(ctx, sql, params...)
	} else {
		rows, err = tx.QueryContext(ctx, sql, params...)
	}
	if err != nil {
		return 0, fmt.Errorf("count 錯誤: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&ret,
		); err != nil {
			return 0, fmt.Errorf("scan 錯誤: %v", err)
		}
	}

	return ret, nil
}
