package main

import (
	"database/sql"
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"flag"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", DSN())
	if err != nil {
		panic(err)
	}

	vars, err := QueryVariables(db)
	if err != nil {
		panic(err)
	}

	err = vars.WriteCSV(os.Stdout)
	if err != nil {
		panic(err)
	}
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

func QueryVariables(db *sql.DB) (Variables, error) {
	vars := Variables{}
	res, err := db.Query("SHOW VARIABLES")
	if err != nil {
		return vars, err
	}
	defer res.Close()

	for res.Next() {
		var v Variable
		err := res.Scan(&v.Name, &v.Value)
		if err != nil {
			return vars, err
		}
		vars = append(vars, v)
	}
	return vars, nil
}

type Variables []Variable

type Variable struct {
	Name  string
	Value string
}

func (vars Variables) Rows() [][]string {
	rows := [][]string{}
	for _, vi := range vars {
		rows = append(rows, []string{vi.Name, vi.Value})
	}
	return rows
}

func (vars Variables) WriteCSV(w io.Writer) error {
	csvw := csv.NewWriter(w)
	csvw.WriteAll(vars.Rows())
	err := csvw.Error()
	if err != nil {
		return err
	}
	csvw.Flush()
	return csvw.Error()
}
