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

type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

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

func ReadUser(pool RowQueryerContext, ctx context.Context, login string) (User, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var Login string
	var Firstname string
	var Lastname string
	var Email string
	err := pool.QueryRowContext(ctx, "select o.login, o.firstname, o.lastname, o.email from users as o where o.login = :login;", sql.Named("login", login)).Scan(&Login, &Firstname, &Lastname, &Email)
	return MakeUser(Login, Firstname, Lastname, Email), err
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

func deleteUser(pool ExecerContext, ctx context.Context, login string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "delete from users where login = :login;", sql.Named("login", login))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

type Message struct {
	Id        int64
	UserLogin string
	Content   string
}

func MakeMessage(id int64, userlogin string, content string) Message {
	return Message{
		Content:   content,
		Id:        id,
		UserLogin: userlogin,
	}
}

func (o Message) Create(pool ExecerContext, ctx context.Context) error {
	_, err := createMessage(pool, ctx, o.Id, o.UserLogin, o.Content)
	return err
}

func ReadMessage(pool RowQueryerContext, ctx context.Context, id int64) (Message, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var Id int64
	var UserLogin string
	var Content string
	err := pool.QueryRowContext(ctx, "select o.id, o.userlogin, o.content from messages as o where o.id = :id;", sql.Named("id", id)).Scan(&Id, &UserLogin, &Content)
	return MakeMessage(Id, UserLogin, Content), err
}

func (o Message) Update(pool ExecerContext, ctx context.Context) error {
	_, err := updateMessage(pool, ctx, o.Id, o.UserLogin, o.Content)
	return err
}

func (o Message) Delete(pool ExecerContext, ctx context.Context) error {
	_, err := deleteMessage(pool, ctx, o.Id)
	return err
}

func createMessage(pool ExecerContext, ctx context.Context, Id int64, UserLogin string, Content string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "insert into messages(id, userlogin, content) values(:Id, :UserLogin, :Content);", sql.Named("Id", Id), sql.Named("UserLogin", UserLogin), sql.Named("Content", Content))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func updateMessage(pool ExecerContext, ctx context.Context, Id int64, UserLogin string, Content string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "update messages set userlogin = :UserLogin, content = :Content where id = :Id;", sql.Named("Id", Id), sql.Named("UserLogin", UserLogin), sql.Named("Content", Content))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func deleteMessage(pool ExecerContext, ctx context.Context, id int64) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "delete from messages where id = :id;", sql.Named("id", id))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func getMessagesByUserLogin(pool QueryerContext, ctx context.Context, login string) ([]Message, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var Id int64
	var UserLogin string
	var Content string
	rows, err := pool.QueryContext(ctx, "select o.id, o.userLogin, o.content from messages as o where o.userLogin = :login;", sql.Named("login", login))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []Message{}
	for rows.Next() {
		err := rows.Scan(&Id, &UserLogin, &Content)
		if err != nil {
			return nil, err
		}
		res = append(res, MakeMessage(Id, UserLogin, Content))
	}
	return res, nil
}

func (o User) GetMessages(pool QueryerContext, ctx context.Context) ([]Message, error) {
	return getMessagesByUserLogin(pool, ctx, o.Login)
}
