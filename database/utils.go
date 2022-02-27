package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func InsertOrUpdatePlayer(player PlayerInfo) error {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return err
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT OR REPLACE INTO playerinfo(ip, words_guessed, losses) values (?, ?, ?)")
	_, err = stmt.Exec(player.IP, player.WordsGuessed, player.Losses)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return nil
}

func ReadPlayerInfo(ip string) (PlayerInfo, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	res := PlayerInfo{}
	res.IP = ip
	if err != nil {
		return res, err
	}
	defer db.Close()
	stmt, err := db.Prepare("SELECT * FROM playerinfo WHERE ip=?")
	if err != nil {
		return res, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(ip).Scan(&res.IP, &res.WordsGuessed, &res.Losses)
	if err != nil {
		return res, err
	}
	fmt.Println(res)
	return res, nil
}
