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

func MakeUser(name string, email string) User {
	return User{
		Email: email,
		Name:  name,
	}
}

func (o User) Create(pool ExecerContext, ctx context.Context) error {
	_, err := createUser(pool, ctx, o.Name, o.Email)
	return err
}

func ReadUser(pool RowQueryerContext, ctx context.Context, name string) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var Name string
	var Email string
	err := pool.QueryRowContext(ctx, "select o.name, o.email from users as o where o.name = :name;", sql.Named("name", name)).Scan(&Name, &Email)
	return MakeUser(Name, Email), err
}

func (o User) Update(pool ExecerContext, ctx context.Context) error {
	_, err := updateUser(pool, ctx, o.Name, o.Email)
	return err
}

func (o User) Delete(pool ExecerContext, ctx context.Context) error {
	_, err := deleteUser(pool, ctx, o.Name)
	return err
}

func createUser(pool ExecerContext, ctx context.Context, Name string, Email string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "insert into users(name, email) values(:Name, :Email);", sql.Named("Name", Name), sql.Named("Email", Email))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func updateUser(pool ExecerContext, ctx context.Context, Name string, Email string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "update users set email = :Email where name = :Name;", sql.Named("Name", Name), sql.Named("Email", Email))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func deleteUser(pool ExecerContext, ctx context.Context, name string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "delete from users where name = :name;", sql.Named("name", name))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}
