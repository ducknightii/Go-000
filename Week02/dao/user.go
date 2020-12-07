package dao

import (
	"database/sql"

	"github.com/pkg/errors"
)

func Age(name string) (age int, err error) {
	stmt, err := DB.Prepare("select age from users where name = ? limit 1")
	if err != nil {
		err = errors.Wrap(err, "prepare failed")
		return
	}
	err = stmt.QueryRow(name).Scan(&age)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		err = RecordNotFond
		return
	}
	if err != nil {
		err = errors.Wrap(err, "age find err")
	}
	return
}
