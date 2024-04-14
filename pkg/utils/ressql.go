package utils

import (
	"fmt"
	"strings"
)

func ResSql(sql string, args ...any) string {
	for _, arg := range args {
		argStr := fmt.Sprintf("'%v'", arg)
		sql = strings.Replace(sql, "?", argStr, 1)
	}
	return sql
}
