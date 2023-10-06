// Generated from ../test.crn - do not edit.

package model

import (
	"context"
	"database/sql"
)

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

func ReadMessage(pool RowQueryerContext, ctx context.Context, Id int64) (Message, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var IdTemp int64
	var UserLoginTemp string
	var ContentTemp string
	err := pool.QueryRowContext(ctx, "select o.id, o.user_login, o.content from messages as o where o.id = :Id;", sql.Named("Id", Id)).Scan(&IdTemp, &UserLoginTemp, &ContentTemp)
	return MakeMessage(IdTemp, UserLoginTemp, ContentTemp), err
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

	result, err := pool.ExecContext(ctx, "insert into messages(id, user_login, content) values(:Id, :UserLogin, :Content);", sql.Named("Id", Id), sql.Named("UserLogin", UserLogin), sql.Named("Content", Content))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func updateMessage(pool ExecerContext, ctx context.Context, Id int64, UserLogin string, Content string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "update messages set user_login = :UserLogin, content = :Content where id = :Id;", sql.Named("Id", Id), sql.Named("UserLogin", UserLogin), sql.Named("Content", Content))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func deleteMessage(pool ExecerContext, ctx context.Context, Id int64) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	result, err := pool.ExecContext(ctx, "delete from messages where id = :Id;", sql.Named("Id", Id))
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func getMessagesByUserLogin(pool QueryerContext, ctx context.Context, login string) ([]Message, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	var IdTemp int64
	var UserLoginTemp string
	var ContentTemp string
	rows, err := pool.QueryContext(ctx, "select o.id, o.user_login, o.content from messages as o where o.user_login = :login;", sql.Named("login", login))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []Message{}
	for rows.Next() {
		err := rows.Scan(&IdTemp, &UserLoginTemp, &ContentTemp)
		if err != nil {
			return nil, err
		}
		results = append(results, MakeMessage(IdTemp, UserLoginTemp, ContentTemp))
	}
	return results, nil
}

func (o Message) GetUser(pool RowQueryerContext, ctx context.Context) (User, error) {
	return ReadUser(pool, ctx, o.UserLogin)
}
