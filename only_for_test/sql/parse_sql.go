/**
 * @Author: xueyanghan
 * @File: parse_sql.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/11/17 10:33
 */

package sql

import (
	"fmt"
	"github.com/xwb1989/sqlparser"
	"log"
	"testing"
)

func sql_parse_test(t *testing.T) {
	sqlStatement := "SELECT id, name FROM users WHERE id = 1"

	stmt, err := sqlparser.Parse(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}

	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		// 处理SELECT语句
		for _, expr := range stmt.SelectExprs {
			fmt.Printf("Column: %s\n", sqlparser.String(expr))
		}

		tableExprs := stmt.From
		for _, tableExpr := range tableExprs {
			fmt.Printf("Table: %s\n", sqlparser.String(tableExpr))
		}

		// 处理WHERE条件
		if stmt.Where != nil {
			fmt.Printf("Where: %s\n", sqlparser.String(stmt.Where))
		}
	default:
		fmt.Println("Unsupported SQL statement")
	}
}
