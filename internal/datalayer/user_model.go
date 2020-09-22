package datalayer

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
)

type User struct {
	id        int64
	email     string
	firstname string
	lastname  string
}

func getUserByEmail(email string) (*User, error) {
	var user User
	selectStatement := `
		SELECT id, email, firstname, lastname FROM priceator_user WHERE email=$1 
	`
	err := Db.QueryRow(context.Background(),
		selectStatement,
		email).Scan(&user.id, &user.email, &user.firstname, &user.lastname)
	switch err {
	case nil:
		return &user, nil
	case pgx.ErrNoRows:
	default:
		log.Debug().Err(err).Msgf("could not retrieve a user with email %s", email)
		return nil, nil
	}
	return &user, nil
}
func getUserById(id int) (*User, error) {
	var user = new(User)
	selectStatement := `
		SELECT id, email, firstname, lastname FROM priceator_user WHERE id=$1`
	err := Db.QueryRow(context.Background(),
		selectStatement,
		id).Scan(&user)
	switch err {
	case nil:
		return user, nil
	case pgx.ErrNoRows:
	default:
		return nil, nil
	}
	return nil, nil
}
func (user *User) delete() error {
	deleteStatement := `DELETE FROM priceator_user WHERE email=$1`
	cmd, err := Db.Exec(context.Background(),
		deleteStatement,
		user.email)
	if err != nil {
		log.Error().Err(err).Msgf("unable to delete user %s from database", user.email)
		return err
	}
	if cmd.Delete() == true {
		log.Debug().Msgf("deleted user %s from database", user.email)
	}
	return nil
}
func (user *User) save() error {
	saveStatement := `INSERT INTO priceator_user(email, firstname, lastname) VALUES ($1, $2, $3)`
	cmd, err := Db.Exec(context.Background(),
		saveStatement,
		user.email,
		user.firstname,
		user.lastname)
	if err != nil {
		log.Error().Err(err).Msg("unable to save user to database")
		return err
	}
	if cmd.Insert() == true {
		log.Debug().Msgf("saved user %s to database", user.email)
	}
	return nil
}
