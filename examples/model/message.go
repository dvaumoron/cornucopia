// Generated from ../test.crn - do not edit.

package model

import "context"

type Message struct {
	Id        int64
	UserLogin string
	Content   string
}

func MakeMessage(id int64, userLogin string, content string) Message {
	return Message{
		Content:   content,
		Id:        id,
		UserLogin: userLogin,
	}
}

func (o Message) Create(pool ExecerContext, ctx context.Context) error {
	_, err := createMessage(pool, ctx, o.UserLogin, o.Content)
	return err
}

func ReadMessage(pool RowQueryerContext, ctx context.Context, id int64) (Message, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := "select o.id, o.user_login, o.content from messages as o where o.id = $1;"
	var idTemp int64
	var userLoginTemp string
	var contentTemp string
	err := pool.QueryRowContext(ctx, query, id).Scan(&idTemp, &userLoginTemp, &contentTemp)
	return MakeMessage(idTemp, userLoginTemp, contentTemp), err
}

func (o Message) Update(pool ExecerContext, ctx context.Context) error {
	_, err := updateMessage(pool, ctx, o.Id, o.UserLogin, o.Content)
	return err
}

func (o Message) Delete(pool ExecerContext, ctx context.Context) error {
	_, err := deleteMessage(pool, ctx, o.Id)
	return err
}

func createMessage(pool ExecerContext, ctx context.Context, userLogin string, content string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := "insert into messages(user_login, content) values($1, $2);"
	result, err := pool.ExecContext(ctx, query, userLogin, content)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func updateMessage(pool ExecerContext, ctx context.Context, id int64, userLogin string, content string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := "update messages set user_login = $2, content = $3 where id = $1;"
	result, err := pool.ExecContext(ctx, query, id, userLogin, content)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func deleteMessage(pool ExecerContext, ctx context.Context, id int64) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := "delete from messages where id = $1;"
	result, err := pool.ExecContext(ctx, query, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func getMessagesByUserLogin(pool QueryerContext, ctx context.Context, login string) ([]Message, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := "select m.id, m.user_login, m.content from messages as m where m.user_login = :login;"
	var idTemp int64
	var userLoginTemp string
	var contentTemp string
	rows, err := pool.QueryContext(ctx, query, login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []Message{}
	for rows.Next() {
		err := rows.Scan(&idTemp, &userLoginTemp, &contentTemp)
		if err != nil {
			return nil, err
		}
		results = append(results, MakeMessage(idTemp, userLoginTemp, contentTemp))
	}
	return results, nil
}

func (o Message) GetUser(pool RowQueryerContext, ctx context.Context) (User, error) {
	return ReadUser(pool, ctx, o.UserLogin)
}
