package gomysql

import (
	"database/sql"
	"encoding/csv"
	"io"
	"strings"
)

func QueryVersion(db *sql.DB) (string, error) {
	res, err := db.Query("SELECT VERSION()")
	if err != nil {
		return "", err
	}
	defer res.Close()

	var ver string
	for res.Next() {
		err := res.Scan(&ver)
		if err != nil {
			return "", err
		}
		break
	}
	return strings.TrimSpace(ver), nil
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
	err := csvw.WriteAll(vars.Rows())
	if err != nil {
		return err
	}
	csvw.Flush()
	return csvw.Error()
}

func GetInt(db *sql.DB, sqlQuery string) (int, error) {
	var v int
	res, err := db.Query(sqlQuery)
	if err != nil {
		return v, err
	}
	defer res.Close()

	for res.Next() {
		err := res.Scan(&v)
		if err != nil {
			return v, err
		}
		break
	}
	return v, nil
}
