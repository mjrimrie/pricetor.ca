package datalayer

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"os"
)

const databaseName = "priceator"

var Db *pgxpool.Pool

func createUserTable(conn *pgxpool.Pool) error {
	createStatement := `
CREATE TABLE IF NOT EXISTS priceator_user (
id SERIAL PRIMARY KEY,
email VARCHAR UNIQUE,
firstname VARCHAR,
lastname VARCHAR)`
	_, err := conn.Exec(context.Background(), createStatement)
	if err != nil {
		log.Error().Err(err).Msg("failure executing 'user' table create statement")
		return err
	}
	return nil
}
func dropPriceatorUserTable(conn *pgxpool.Pool) error {
	dropStatement := `DROP TABLE priceator_user`
	_, err := conn.Exec(context.Background(), dropStatement)
	if err != nil {
		log.Error().Err(err).Msg("failure executing 'priceator_user' table drop statement")
		return err
	}
	return nil
}
func createMLSWatchTable(conn *pgxpool.Pool) error {
	createStatement := `
CREATE TABLE IF NOT EXISTS mls_watch (
id SERIAL PRIMARY KEY,
listingId VARCHAR NOT NULL,
mlsNumber VARCHAR NOT NULL,
priceator_user INTEGER REFERENCES priceator_user(id))`
	_, err := conn.Exec(context.Background(), createStatement)
	if err != nil {
		log.Error().Err(err).Msg("failure executing 'mls_watch' table create statement")
		return err
	}
	return nil
}
func dropMLSWatchTable(conn *pgxpool.Pool) error {
	dropStatement := `DROP TABLE mls_watch`
	_, err := conn.Exec(context.Background(), dropStatement)
	if err != nil {
		log.Error().Err(err).Msg("failure executing 'mls_watch' table drop statement")
		return err
	}
	return nil
}
func dropDatabase(psqlInfo *string) error {
	db, err := sql.Open("postgres", *psqlInfo)
	if err != nil {
		log.Error().Err(err).Msgf("Could not open database server connection to create %s database",
			databaseName)
		return err
	}
	dropDatabaseStatement := `DROP DATABASE priceator`
	_, err = db.ExecContext(context.Background(), dropDatabaseStatement)
	if err != nil {
		log.Error().Err(err).Msg("Could not create database 'priceator'")
		return err
	}
	return nil
}
func createDatabase(psqlInfo *string) error {
	db, err := sql.Open("postgres", *psqlInfo)
	if err != nil {
		log.Error().Err(err).Msgf("Could not open database server connection to create %s database",
			databaseName)
		return err
	}
	checkForExistingDatabaseStatement := fmt.Sprintf(
		"SELECT EXISTS(SELECT datname FROM pg_database WHERE datname = '%s')",
		databaseName)
	var exists bool
	err = db.QueryRow(checkForExistingDatabaseStatement).Scan(&exists)
	if err != nil {
		log.Error().Err(err).Msgf("Could not check for existing %s database", databaseName)
		return err
	}
	if exists {
		log.Info().Msgf("database '%s' already exists", databaseName)
		return nil
	}
	createDatabaseStatement := `CREATE DATABASE priceator`

	_, err = db.ExecContext(context.Background(), createDatabaseStatement)
	if err != nil {
		log.Error().Err(err).Msg("Could not create database 'priceator'")
		return err
	}
	return nil
}
func Initialize(psqlInfo *string) error {
	err := createDatabase(psqlInfo)
	if err != nil {
		log.Error().Err(err).Msg("could not create database during initializer")
		return err
	}
	psqlDatabaseInfo := fmt.Sprintf(*psqlInfo + " database=priceator")
	err = Connect(psqlDatabaseInfo)
	if err != err {
		log.Fatal().Err(err).Msg("could not initialize datalayer")
		os.Exit(1)
	}

	err = createUserTable(Db)
	if err != nil {
		os.Exit(1)
	}
	err = createMLSWatchTable(Db)
	if err != nil {
		os.Exit(1)
	}
	return nil
}
func Destroy(psqlInfo *string) error{
	fmt.Sprintf(*psqlInfo + " database=priceator")
	err := dropDatabase(psqlInfo)
	if err != nil{
		log.Error().Err(err).Msg("could not drop database")
	}
	return nil
}
func Connect(database_url string) error {
	poolConfig, err := pgxpool.ParseConfig(database_url)
	if err != nil {
		log.Error().Err(err).Msg("unable to parse database config")
		return err
	}
	Db, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Error().Err(err).Msg("unable to connect to database")
		return err
	}
	return nil
}
