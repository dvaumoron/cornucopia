// Generated from ../test.crn - do not edit.

package model

import "context"

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

	query := "select o.login, o.firstname, o.lastname, o.email from users as o where o.login = $1;"
	var LoginTemp string
	var FirstnameTemp string
	var LastnameTemp string
	var EmailTemp string
	err := pool.QueryRowContext(ctx, query, Login).Scan(&LoginTemp, &FirstnameTemp, &LastnameTemp, &EmailTemp)
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

	query := "insert into users(login, firstname, lastname, email) values($1, $2, $3, $4);"
	result, err := pool.ExecContext(ctx, query, Login, Firstname, Lastname, Email)
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func updateUser(pool ExecerContext, ctx context.Context, Login string, Firstname string, Lastname string, Email string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := "update users set firstname = $2, lastname = $3, email = $4 where login = $1;"
	result, err := pool.ExecContext(ctx, query, Login, Firstname, Lastname, Email)
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func deleteUser(pool ExecerContext, ctx context.Context, Login string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	query := "delete from users where login = $1;"
	result, err := pool.ExecContext(ctx, query, Login)
	if err != nil {
		return int64(0), err
	}
	return result.RowsAffected()
}

func (o User) GetMessages(pool QueryerContext, ctx context.Context) ([]Message, error) {
	return getMessagesByUserLogin(pool, ctx, o.Login)
}
