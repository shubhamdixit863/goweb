package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type CommentModelInterface interface {
	Insert(postid int, userid int, content string) (int, error)
	GetCommentsByPostId(postid int) ([]*Comment, error)
}

type Comment struct {
	ID      int
	PostId  int
	UserId  int
	Content string
	Created string
}

type CommentModel struct {
	DB *sql.DB
}

func (m *CommentModel) Insert(postid int, userid int, content string) (int, error) {
	stmt := `INSERT INTO Comments (postid, userid,content, created)
    VALUES(?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, postid, userid, content, time.Now().Format("2006-01-02 15:04:05 Monday"))
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *CommentModel) GetCommentsByPostId(postid int) ([]*Comment, error) {
	log.Println(postid)
	stmt := `SELECT  content, created FROM comments
    WHERE  postId = ?`

	rows, err := m.DB.Query(stmt, postid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	Comments := []*Comment{}

	for rows.Next() {
		s := &Comment{}

		err = rows.Scan(&s.Content, &s.Created)
		if err != nil {
			return nil, err
		}

		Comments = append(Comments, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(Comments)

	return Comments, nil

}
