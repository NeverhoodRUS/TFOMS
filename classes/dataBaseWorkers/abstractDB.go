package databaseworkers

import (
	"database/sql"
	"errors"
	"fmt"
	"tfoms_server/static/exceptions"
	"time"

	"github.com/beego/beego"
	_ "github.com/lib/pq"
)

var (
	host     string = ""
	port     int
	user     string = ""
	password string = ""
	dbname   string = ""

	max_conn  int = 40
	idle_conn int = 10

	db *sql.DB // Глобальное соединение sql

	DB_CONN_ERROR error
)

func init() {
	host = beego.AppConfig.String("db_host")
	port, _ = beego.AppConfig.Int("db_port")
	user = beego.AppConfig.String("db_user")
	password = beego.AppConfig.String("db_password")
	dbname = beego.AppConfig.String("db_name")

	max_conn = beego.AppConfig.DefaultInt("db_max_conn", 50)
	idle_conn = beego.AppConfig.DefaultInt("db_idle_conn", 10)

	DB_CONN_ERROR = errors.New(exceptions.DB_connection_error)
}

func GetPostgresSession() *sql.DB {
	if db == nil {
		psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`, host, port, user, password, dbname)

		dbOpen, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			return nil
		}

		dbOpen.SetConnMaxLifetime(30 * time.Minute)
		dbOpen.SetMaxOpenConns(max_conn)
		dbOpen.SetMaxIdleConns(idle_conn)

		//		err = db.Ping()
		//		if err != nil {
		//			return nil
		//		}

		db = dbOpen
	}

	return db
}

func GetAllRowsAsMap(table string) (*sql.Rows, error) {
	rows, err := db.Query("select * from " + table)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// func rowsToMap(rows *sql.Rows) (*[]map[string]interface{}, error) {
// 	columns, _ := rows.Columns()
// 	columnLength := len(columns)
// 	cache := make([]interface{}, columnLength) //temporarily store each row of data
// 	for index := range cache {                 // initialize a pointer for each column
// 		var a interface{}
// 		cache[index] = &a
// 	}
// 	var list []map[string]interface{} //returned slice
// 	for rows.Next() {
// 		_ = rows.Scan(cache...)

// 		item := make(map[string]interface{})
// 		for i, data := range cache {
// 			item[columns[i]] = *data.(*interface{})
// 		}
// 		list = append(list, item)
// 	}
// 	_ = rows.Close()
// 	return &list, nil
// }
