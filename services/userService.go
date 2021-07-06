package services

import (
	"context"
	"encoding/json"
	"go-redis/db"
	m "go-redis/models"

	"github.com/google/uuid"
)

var (
	ctx = context.Background()
)

func AddUser(u m.User) (m.User, error) {

	rdb := db.PoolRDB()

	u.ID = uuid.NewString()

	uJson, err := json.Marshal(u)

	if err != nil {
		return u, err
	}

	err = rdb.Set(ctx, u.ID, uJson, 0).Err()

	if err != nil {
		return u, err
	}

	return u, nil
}

func GetUser(ID string) (m.User, error) {

	rdb := db.PoolRDB()

	val, err := rdb.Get(ctx, ID).Result()

	var u m.User

	if err != nil {
		return u, err
	}

	err = json.Unmarshal([]byte(val), &u)

	if err != nil {
		return u, err
	}

	return u, err
}

func DeleteUser(ID string) error {

	rdb := db.PoolRDB()

	return rdb.Del(ctx, ID).Err()

}
