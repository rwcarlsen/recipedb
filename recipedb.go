package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/mxk/go-sqlite/sqlite3"
)

var fname = flag.String("csv", "", "recipe csv file path")
var dbname = flag.String("db", "recipes.sqlite", "recipe sqlite db file path")

const createtblinfo = `
CREATE TABLE IF NOT EXISTS info (

`

func main() {
	flag.Parse()
	db, err := sql.Open("sqlite3", *dbname)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err := db.Exec()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(*fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	r := csv.NewReader(f)

	col := 0
	ncols := 1
	nInfoCols := 2 // number of columns before fuel recipe data start
	mode := ""
	for col < ncols {
		for {
			record, err := r.Read()
			if len(record) > 0 {
				if ncols == 1 {
					ncols = len(record) - nInfoCols
				}
				if s := record[0]; s != "" {
					mode = strings.ToLower(strings.TrimSpace(s))
				}

				switch mode {
				case "basic info":
				case "feed material":
				case "irradiation parameters":
				case "input mass fractions":
				case "output mass fractions":
				case "full isotopic data for output composition":
				}
				fmt.Println(record[0])
			}

			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
		}
		col++
	}
}
