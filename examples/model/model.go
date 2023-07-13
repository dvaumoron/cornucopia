// Generated from ../test.crn - do not edit.

package model

import (
	"context"
	"database/sql"
	"time"
)

const timeout = time.Duration(int64(5)) * time.Second

type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

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

func SaveUser(pool ExecerContext, ctx context.Context, Name string, Email string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "insert into users(name, email) values(:Name, :Email)", sql.Named("Name", Name), sql.Named("Email", Email))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func GetUserByName(pool RowQueryerContext, ctx context.Context, name string) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var Name string
	var Email string
	err := pool.QueryRowContext(ctx, "select u.name, u.email from users as u where u.name = :name;", sql.Named("name", name)).Scan(&Name, &Email)
	return MakeUser(Name, Email), err
}
