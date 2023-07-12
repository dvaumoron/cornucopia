package model

import (
	"context"
	"database/sql"
	"time"
)

type RowQueryerContext interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

func MakeUser(name string, email string) User {
	return User{
		email: email,
		name:  name,
	}
}
func QueryUser(pool RowQueryerContext, ctx context.Context, id int64) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, int64(5)*time.Second)
	defer cancel()

	var name string
	var email string
	err := pool.QueryRowContext(ctx, "select u.name, u.email from user as u where u.id = :id;", sql.Named("id", id)).Scan(&name, &email)
	return MakeUser(name, email), err
}
