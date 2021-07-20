package data

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id       int
	Uuid     string
	Name     string
	Email    string
	Password string
	CreateAt time.Time
}

type Session struct {
	Id      int
	Uuid    string
	Email   string
	UserId  int
	CreatAt time.Time
}

// create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	statement := "INSERT INTO sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	// use QueryRow to return a row and scan the returned id into the Session struct
	err = stmt.QueryRow(uuid.New().String(), user.Email, user.Id, time.Now()).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatAt)
	return
}

// check if session is valid in the database
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatAt)

	if err != nil {
		valid = false
		return
	}
	if session.Id != 0 {
		valid = true
	}
	return
}
