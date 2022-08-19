package models

import (
	"database/sql"
	"time"
)

type LikesModelInterface interface {
	Insert(postid int, userid int) (int, error)
	GetTotalLikesByPostId(postid int) (int, error)
}

type Likes struct {
	ID      int
	PostId  int
	UserId  int
	Created string
}

type LikesModel struct {
	DB *sql.DB
}

func (m *LikesModel) Insert(postid int, userid int) (int, error) {
	stmt := `INSERT INTO likes (postid, userid, created)
    VALUES(?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, postid, userid, time.Now().Format("2006-01-02 15:04:05 Monday"))
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *LikesModel) GetTotalLikesByPostId(postid int) (int, error) {
	var count int
	stmt := `SELECT  COUNT(*) FROM likes
    WHERE  postId = ?`

	rows, err := m.DB.Query(stmt, postid)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	err = rows.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil

}
