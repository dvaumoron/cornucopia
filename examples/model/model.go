package model

import (
	"context"
	"database/sql"
	"time"
)

type RowQueryerContext interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type User struct {
	Name  string
	Email string
}

func MakeUser(Name string, Email string) User {
	return User{
		Email: Email,
		Name:  Name,
	}
}

func QueryUser(pool RowQueryerContext, ctx context.Context, name string) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(int64(5))*time.Second)
	defer cancel()

	var Name string
	var Email string
	err := pool.QueryRowContext(ctx, "select u.name, u.email from user as u where u.name = :name;", sql.Named("name", name)).Scan(&Name, &Email)
	return MakeUser(Name, Email), err
}
