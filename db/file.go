package db

import (
	"database/sql"

	mydb "github.com/u8x9/godisk/db/mysql"
)

// OnFileUploadFinished 文件上传保存meta
func OnFileUploadFinished(filehash, filename string, filesize int64, fileaddr string) (id int64, err error) {
	stmt, err := mydb.DBConn().Prepare("INSERT IGNORE INTO tbl_file (file_sha1, file_name, file_size, file_addr, `status`) VALUES (?, ?, ?, ?, 1)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

type TableFile struct {
	FileHash string
	Filename sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare("SELECT file_sha1, file_name, file_size, file_addr FROM tbl_file WHERE file_sha1=? AND status=? LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(filehash, 1)
	var ft = new(TableFile)
	if err := row.Scan(&ft.FileHash, &ft.Filename, &ft.FileSize, &ft.FileAddr); err != nil {
		return nil, err
	}
	return ft, nil
}
