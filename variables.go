package gomysql

import (
	"encoding/csv"
	"io"
)

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
