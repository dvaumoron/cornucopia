// Generated from ../test.crn - do not edit.

package model

import (
	"context"
	"database/sql"
)

type User struct {
	Login     string
	Firstname string
	Lastname  string
	Email     string
}

func MakeUser(login string, firstname string, lastname string, email string) User {
	return User{
		Email:     email,
		Firstname: firstname,
		Lastname:  lastname,
		Login:     login,
	}
}

func (o User) Create(pool ExecerContext, ctx context.Context) error {
	_, err := createUser(pool, ctx, o.Login, o.Firstname, o.Lastname, o.Email)
	return err
}

func ReadUser(pool RowQueryerContext, ctx context.Context, Login string) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var LoginTemp string
	var FirstnameTemp string
	var LastnameTemp string
	var EmailTemp string
	err := pool.QueryRowContext(ctx, "select o.login, o.firstname, o.lastname, o.email from users as o where o.login = :Login;", sql.Named("Login", Login)).Scan(&LoginTemp, &FirstnameTemp, &LastnameTemp, &EmailTemp)
	return MakeUser(LoginTemp, FirstnameTemp, LastnameTemp, EmailTemp), err
}

func (o User) Update(pool ExecerContext, ctx context.Context) error {
	_, err := updateUser(pool, ctx, o.Login, o.Firstname, o.Lastname, o.Email)
	return err
}

func (o User) Delete(pool ExecerContext, ctx context.Context) error {
	_, err := deleteUser(pool, ctx, o.Login)
	return err
}

func createUser(pool ExecerContext, ctx context.Context, Login string, Firstname string, Lastname string, Email string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "insert into users(login, firstname, lastname, email) values(:Login, :Firstname, :Lastname, :Email);", sql.Named("Login", Login), sql.Named("Firstname", Firstname), sql.Named("Lastname", Lastname), sql.Named("Email", Email))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func updateUser(pool ExecerContext, ctx context.Context, Login string, Firstname string, Lastname string, Email string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "update users set firstname = :Firstname, lastname = :Lastname, email = :Email where login = :Login;", sql.Named("Login", Login), sql.Named("Firstname", Firstname), sql.Named("Lastname", Lastname), sql.Named("Email", Email))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func deleteUser(pool ExecerContext, ctx context.Context, Login string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "delete from users where login = :Login;", sql.Named("Login", Login))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func (o User) GetMessages(pool QueryerContext, ctx context.Context) ([]Message, error) {
	return getMessagesByUserLogin(pool, ctx, o.Login)
}
