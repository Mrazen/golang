package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
)

func main() {
	z, _ := zip.OpenReader("task.zip")
	defer z.Close()

	for _, f := range z.File {

		r, _ := f.Open()

		if rows, _ := csv.NewReader(r).ReadAll(); len(rows) == 10 && len(rows[4]) == 10 {
			fmt.Println(f.Name, rows[4][2])
			fmt.Println(rows)
		}

		r.Close()
	}
}
