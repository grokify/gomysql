package gomysql

import (
	"database/sql"
	"strings"
)

func QueryVersion(db *sql.DB) (string, error) {
	res, err := db.Query("SELECT VERSION()")
	if err != nil {
		return "", err
	}
	defer res.Close()

	var ver string
	res.Next()
	err = res.Scan(&ver)
	if err != nil {
		return "", err
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

func GetInt(db *sql.DB, sqlQuery string) (int, error) {
	var v int
	res, err := db.Query(sqlQuery)
	if err != nil {
		return v, err
	}
	defer res.Close()

	res.Next()
	err = res.Scan(&v)
	if err != nil {
		return v, err
	}

	return v, nil
}
