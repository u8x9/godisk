package db

import mydb "github.com/u8x9/godisk/db/mysql"

func SetUserToken(username, token string) (int64, error) {
	tx, err := mydb.DBConn().Begin()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	stmt, err := tx.Prepare("DELETE FROM tbl_user_token WHERE username = ?")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmt.Close()
	if _, err := stmt.Exec(username); err != nil {
		tx.Rollback()
		return 0, err
	}
	stmtInsert, err := tx.Prepare("INSERT INTO tbl_user_token (username, token) VALUES (?,?)")
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	defer stmtInsert.Close()
	result, err := stmtInsert.Exec(username, token)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}
	return result.LastInsertId()
}
