package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/grokify/gomysql"
)

func main() {
	db, err := sql.Open("mysql", DSN())
	if err != nil {
		panic(err)
	}

	vars, err := gomysql.QueryVariables(db)
	if err != nil {
		panic(err)
	}

	err = vars.WriteCSV(os.Stdout)
	if err != nil {
		panic(err)
	}
	ver, err := gomysql.QueryVersion(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(ver)

}

func DSN() string {
	var hostFlag = flag.String("h", "", "Connect to host.")
	var portFlag = flag.Int("P", 3306, "Port number to use for connection.")
	var userFlag = flag.String("u", "", "User for login if not current user.")
	var passwordFlag = flag.String("p", "", "Password to use when connecting to server")

	flag.Parse()

	cfg := mysql.Config{
		User:                 *userFlag,
		Passwd:               *passwordFlag,
		Net:                  "tcp",
		Addr:                 *hostFlag + ":" + strconv.Itoa(*portFlag),
		DBName:               "sys",
		AllowNativePasswords: true,
	}
	return cfg.FormatDSN()
}
