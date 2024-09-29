package pipe_test

import (
	"database/sql"
	"fmt"
	"iter"
	"log"
	"slices"

	"github.com/lufia/pipe"
)

func userNames(rows *sql.Rows) iter.Seq2[string, error] {
	return func(yield func(string, error) bool) {
		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				if !yield("", err) {
					break
				}
				continue
			}
			if !yield(name, nil) {
				break
			}
		}
	}
}

func Example_databaseSql() {
	log.SetFlags(0)
	rdb, err := sql.Open("postgres", "")
	if err != nil {
		log.Fatal(err)
	}
	db := pipe.Value(rdb)
	rows := pipe.TryFrom(db, func(d *sql.DB) (*sql.Rows, error) {
		return d.Query("select name from users")
	}).Defer(func(rows *sql.Rows) { // TODO: err
		rows.Close()
	})
	nameSeq := pipe.From(rows, userNames)
	names, err := pipe.From(nameSeq, slices.Collect).Eval()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(names)
}
