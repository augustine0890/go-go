package data

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

// format the date
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 1, 2021 at 5:34pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 1, 2020 at 5:34pm")
}

// Get all threads in the database and returns it
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

// Create a new thread
func (user *User) CreateThread(topic string) (conv Thread, err error) {
	statement := "INSERT INTO threads (uuid, topic, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, topic, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(uuid.New().String(), topic, user.Id, time.Now()).Scan(
		&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt,
	)
	return
}

// Get a thread by the UUID
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = Db.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = $1", uuid).Scan(
		&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt,
	)
	return
}
