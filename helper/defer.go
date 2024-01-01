package helper

import "database/sql"

func Defer(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollBack := tx.Rollback()
		PanicError(errRollBack)
	} else {
		errCommit := tx.Commit()
		PanicError(errCommit)
	}
}