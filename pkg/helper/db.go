package helper

import (
	"database/sql"
	"log"
)

func StringToNullString(s string) sql.NullString {
	if s != "" {
		return sql.NullString{
			String: s,
			Valid:  true,
		}
	}
	return sql.NullString{
		Valid: false,
	}
}

func WithTransaction(db *sql.DB, fn func(tx *sql.Tx) error) (err error) {
	tx, err := db.Begin()
	if err != nil {
		log.Println("ERROR TX:", err)
		return
	}
	defer func ()  {
		log.Println("DEFER ERR: ", err)
		if err != nil {
			if errRollback := tx.Rollback(); errRollback != nil {
				log.Println("ERROR ROLLBACK: ", errRollback)
			}
		} else {
			if errCommit := tx.Commit(); errCommit != nil {
				err = errCommit
			}
		}
	}()

	err = fn(tx)
	return
}