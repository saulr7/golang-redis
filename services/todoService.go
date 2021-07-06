package services

import (
	"encoding/json"
	"go-redis/db"
	m "go-redis/models"

	"github.com/google/uuid"
)

func AddTodo(t m.Todo) (m.Todo, error) {

	rdb := db.PoolRDB()

	t.ID = uuid.NewString()

	tJson, err := json.Marshal(t)

	if err != nil {
		return t, err
	}

	err = rdb.LPush(ctx, todoHash, tJson, 0).Err()

	if err != nil {
		return t, err
	}

	return t, err

}

func GetTodos() ([]m.Todo, error) {

	rdb := db.PoolRDB()

	var ts []m.Todo
	val := rdb.LRange(ctx, todoHash, 0, 10).Val()

	for _, v := range val {
		var t m.Todo
		json.Unmarshal([]byte(v), &t)
		ts = append(ts, t)
	}

	return ts, nil

}
