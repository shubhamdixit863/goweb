package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

type LikesModelInterface interface {
	Insert(postid int, userid int) (int, error)
	GetTotalLikesByPostId(postid int) (int, error)
	CheckIfUserHasLiked(userid int, postid int) (bool, error)
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
    VALUES(?, ?, ?)`

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

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	return count, nil

}

func (m *LikesModel) CheckIfUserHasLiked(userid int, postid int) (bool, error) {
	log.Println(userid, postid)
	stmt := `SELECT id FROM likes
    WHERE  userid = ? AND postid =? LIMIT 1`

	row := m.DB.QueryRow(stmt, userid, postid)

	s := &Likes{}

	err := row.Scan(&s.ID)
	fmt.Println(err)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, ErrNoRecord
		} else {
			return false, err
		}
	}

	return true, nil

}
