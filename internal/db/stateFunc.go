package db

import (
	"blog-processor/global"
	"blog-processor/internal/model"
	"database/sql"
	"errors"
)

func Exist(blog *model.Blog) (bool, error) {
	var flag bool
	err := global.LiteDB.QueryRow(QUERY_EXIST, blog.Md5Path).Scan(&flag)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	return flag, nil
	//_, err := global.LiteDB.Query(QUERY_EXIST, blog.Md5Path)
	//if err != nil {
	//	if errors.Is(err, sql.ErrNoRows) {
	//		return false, nil
	//	}
	//	return false, err
	//}
	//return true, nil
}

func Insert(blog *model.Blog) error {
	res, err := global.Insert.Exec(
		blog.Md5Path,
		blog.BlogHeader.Date,
		blog.BlogHeader.Date,
		blog.Hash)
	if err != nil {
		return err
	}
	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func Update(blog *model.Blog) error {
	res, err := global.Update.Exec(blog.Hash, blog.BlogHeader.LastMod, blog.Md5Path)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func Delete(blog *model.Blog) error {
	res, err := global.Delete.Exec(blog.Md5Path)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func DeleteByMd5(md5 string) error {
	res, err := global.Delete.Exec(md5)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func QueryMeta(blog *model.Blog) (*model.BlogMeta, error) {
	meta := &model.BlogMeta{}
	rows, err := global.LiteDB.Query(QUERY_META, blog.Md5Path)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
		if count > 1 {
			return nil, errors.New("more than one meta found")
		}
		err = rows.Scan(&meta.CreatedAt, &meta.UpdatedAt, &meta.Hash)
		if err != nil {
			return nil, err
		}
	}
	meta.MD5Path = blog.Md5Path
	return meta, nil
}

func QueryAll(blog *model.Blog) (*model.BlogMeta, error) {
	meta := &model.BlogMeta{}
	rows, err := global.LiteDB.Query(QUERY_META_ALL, blog.Md5Path)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
		if count > 1 {
			return nil, errors.New("more than one meta found")
		}
		err = rows.Scan(&meta.Id, &meta.MD5Path, &meta.CreatedAt, &meta.UpdatedAt, &meta.Hash)
		if err != nil {
			return nil, err
		}
	}
	return meta, nil
}

func QueryAllMd5() ([]string, int, error) {
	rows, err := global.LiteDB.Query(QUERY_ALL_MD5)
	if err != nil {
		return nil, -1, err
	}
	defer rows.Close()
	var md5s []string
	count := 0
	for rows.Next() {
		count++
		var oneMd5 string
		err = rows.Scan(&oneMd5)
		if err != nil {
			return nil, -1, err
		}
		md5s = append(md5s, oneMd5)
	}
	return md5s, count, nil
}
