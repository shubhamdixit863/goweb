package models

import (
	"database/sql"
	"errors"
	"time"
)

type PostModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (*Post, error)
	Latest() ([]*Post, error)
	LikeQuery(data string) ([]*Post, error)
}

type Post struct {
	ID      int
	Title   string
	Content string
	Created string
	Expires string
}

type PostModel struct {
	DB *sql.DB
}

func (m *PostModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO Posts (title, content, expires,created,updated)
    VALUES(?, ?, ?, ?)`

	result, err := m.DB.Exec(stmt, title, content, expires, time.Now().Format("2006-01-02 15:04:05 Monday"), time.Now().Format("2006-01-02 15:04:05 Monday"))
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *PostModel) Get(id int) (*Post, error) {
	stmt := `SELECT id, title, content, created, expires FROM Posts
    WHERE  id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &Post{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *PostModel) Latest() ([]*Post, error) {
	stmt := `SELECT id, title, content, created, expires FROM Posts
    ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	Posts := []*Post{}

	for rows.Next() {
		s := &Post{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		Posts = append(Posts, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return Posts, nil
}

func (m *PostModel) LikeQuery(data string) ([]*Post, error) {

	stmt := `SELECT id, title, content, created, expires FROM Posts WHERE title LIKE ?
    ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt, data+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	Posts := []*Post{}

	for rows.Next() {
		s := &Post{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		Posts = append(Posts, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return Posts, nil
}
