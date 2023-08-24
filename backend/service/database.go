package service

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type JapaneseFamily int64
type JapaneseDifficulty int64

const (
	Romaji JapaneseFamily = iota
	Katakana
	Hiragana
)

const (
	Easy JapaneseDifficulty = iota
	Normal
	Advanced
)

func (jf JapaneseFamily) String() string {
	switch jf {
	case Katakana:
		return "Katakana"
	case Hiragana:
		return "Hiragana"
	case Romaji:
		return "Romaji"
	}
	return "undefined"
}

func (jd JapaneseDifficulty) String() string {
	switch jd {
	case Easy:
		return "Easy"
	case Normal:
		return "Normal"
	case Advanced:
		return "Advanced"
	}
	return "undefined"
}

type RomajiSchema struct {
	Id         int                `sql:"id"`
	Value      string             `sql:"value"`
	Difficulty JapaneseDifficulty `sql:"difficulty"`
}

type HiKaScheme struct {
	Id       int
	Value    string
	RomajiId string
	family   JapaneseFamily
}

func connectDb() *sql.DB {
	cfg := mysql.Config{
		User:   os.Getenv("mysql_username"),
		Passwd: os.Getenv("mysql_password"),
		Net:    "tcp",
		Addr:   os.Getenv("mysql_host"),
		DBName: os.Getenv("mysql_db"),
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	if pingErr := db.Ping(); pingErr != nil {
		log.Fatalln(pingErr)
	}
	return db
}
