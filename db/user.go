package db

import mydb "github.com/u8x9/godisk/db/mysql"

// UserSignUp 用户注册
func UserSignUp(username, password string) (int64, error) {
	stmt, err := mydb.DBConn().Prepare(`INSERT INTO tbl_user (username, password) VALUES(?,?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(username, password)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UserLogin(username, password string) (ok bool, err error) {
	stmt, err := mydb.DBConn().Prepare("SELECT COUNT(*) as C FROM tbl_user WHERE username = ? AND password=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(username, password)
	r := struct {
		C int64
	}{}
	if err := row.Scan(&r.C); err != nil {
		return false, err
	}
	return r.C == 1, nil
}

