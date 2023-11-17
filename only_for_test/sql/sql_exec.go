/**
 * @Author: xueyanghan
 * @File: main.go
 * @Version: 1.0.0
 * @Description: desc.
 * @Date: 2023/11/17 10:12
 */

package sql

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func sql_exec() {
	// 连接到数据库
	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/dbname")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 要执行的SQL语句
	sqlStatement := "SELECT id, name FROM users WHERE id = 1"

	parts := strings.Fields(sqlStatement)
	switch strings.ToUpper(parts[0]) {
	case "SELECT":
		// 执行SELECT语句
		rows, err := db.Query(sqlStatement)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		// 处理结果
		for rows.Next() {
			var id int
			var name string
			err := rows.Scan(&id, &name)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("ID: %d, Name: %s\n", id, name)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
	case "INSERT", "UPDATE", "DELETE":
		// 执行非查询语句
		result, err := db.Exec(sqlStatement)
		if err != nil {
			log.Fatal(err)
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Rows affected: %d\n", rowsAffected)
	default:
		fmt.Println("Unsupported SQL operation")
	}
}
