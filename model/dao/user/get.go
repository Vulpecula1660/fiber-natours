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

// Get : 取得 user
func Get(ctx context.Context, input *GetInput, tx *dbSQL.Tx) (ret []*dto.User, err error) {
	var dbS *dbSQL.DB

	if tx == nil {
		dbS = postgresql.GetConn(os.Getenv("DATABASE_Name"))
	}

	sql := " SELECT "
	sql += "    id,"
	sql += "    account,"
	sql += "    password"

	sql += " FROM user "
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
		return nil, fmt.Errorf("sql 語法錯誤")
	}

	sql += strings.Join(wheres, " AND ")

	var rows *dbSQL.Rows

	if tx == nil {
		rows, err = dbS.QueryContext(ctx, sql, params...)
	} else {
		rows, err = tx.QueryContext(ctx, sql, params...)
	}

	if err != nil {
		return nil, fmt.Errorf("select 錯誤: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		data := &dto.User{}
		if err := rows.Scan(
			&data.ID,
			&data.Account,
			&data.Password,
		); err != nil {
			return nil, fmt.Errorf("scan 錯誤: %v", err)
		}
		ret = append(ret, data)
	}

	return ret, err
}
