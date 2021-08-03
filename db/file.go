package db

import (
	"BaiDuPan/db/mysql"
	"database/sql"
	"fmt"
)

func OnFileUploadFinished(filehash string, filename string, filesize int64, fileadddr string) bool {
	stmt, err := mysql.DBCoon().Prepare(
		"insert ignore into tbl_file(`file_sha1`,`file_name`,`file_size`,`file_addr`,`status`)values(?,?,?,?,1)")
	if err != nil {
		fmt.Print("Failed to prepare statement,err" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(filehash, filename, filesize, fileadddr)
	if err != nil {
		fmt.Print(err.Error())
		return false
	}
	//判断是否有新的记录插入进去，RowsAffected表示行影响
	if rf, err := ret.RowsAffected(); nil == err {
		if rf <= 0 {
			fmt.Print("File with hash:%s has been uploaded before", filehash)
		}
		return true
	}
	return false
}

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

//从mysql获取文件元信息
func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mysql.DBCoon().Prepare("select file_sha1,file_addr,file_name,file_size from tbl_file where file_sha1 = ? and status =1")
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	defer stmt.Close()

	tfile := &TableFile{}
	err = stmt.QueryRow(filehash).Scan(&tfile.FileHash, &tfile.FileAddr, &tfile.FileName, &tfile.FileSize)
	if err != nil {
		fmt.Printf(err.Error())
		return nil, err
	}
	return tfile, err

}
