package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
)

const sqlite3DatabaseFilePath = "./db/iot.sqlite3"

type database struct {
	db                   *sql.DB
	tx                   *sql.Tx
	insertStmt           *sql.Stmt
	getGoInTimestampStmt *sql.Stmt
	deleteStmt           *sql.Stmt
}

func SetupDatabase() *database {
	os.Remove(sqlite3DatabaseFilePath)

	db, err := sql.Open("sqlite3", sqlite3DatabaseFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	insertStmt, err := tx.Prepare("INSERT INTO current_users(id, identifiedBy) VALUES(?, ?);")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	getGoInTimestampStmt, err := tx.Prepare("SELECT goInTimestamp FROM current_users WHERE id = ?")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot gen get go in timestamp stmt")
	}

	deleteStmt, err := tx.Prepare("DELETE FROM current_users WHERE id = ?")
	if err != nil {
		log.Fatal().Err(err).Msg("Cannot gen delete stmt")
	}

	return &database{
		db:                   db,
		tx:                   tx,
		insertStmt:           insertStmt,
		getGoInTimestampStmt: getGoInTimestampStmt,
		deleteStmt:           deleteStmt,
	}
}

func (d database) commit() {
	err := d.tx.Commit()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func (d database) InsertNewCurrentUser(id string, identifiedBy string) {
	_, err := d.insertStmt.Exec(id, identifiedBy)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	d.commit()
}

func (d database) GetGoInTimestampOfUser(id string) string {
	var timestamp string
	err := d.getGoInTimestampStmt.QueryRow(id).Scan(timestamp)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	return timestamp
}

func (d database) DeleteUser(id string) {
	_, err := d.deleteStmt.Exec(id)
	if err != nil {
		log.Error().Err(err).Msg("")
	}

	d.commit()
}

func (d database) Close() {
	d.insertStmt.Close()
	d.getGoInTimestampStmt.Close()
	d.deleteStmt.Close()
	d.db.Close()
}
