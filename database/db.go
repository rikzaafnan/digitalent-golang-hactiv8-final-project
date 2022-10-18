package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	username = "root"
	password = "root"
	dbDriver = "postgres"
	dbPort   = "5432"
	dbName   = "golang_test"
	host     = "localhost"
	db       *sqlx.DB
	err      error
)

func createRequiredTables() {
	UserTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username varchar(255) NOT NULL UNIQUE,
		email varchar(255) NOT NULL UNIQUE,
		password varchar(255) NOT NULL,
		age int NOT NULL,
		profile_image_url varchar(255) NOT NULL,
		created_at timestamptz NOT NULL DEFAULT (now()),
		updated_at timestamptz  NULL DEFAULT NULL
	);`

	PhotoTable := `
	CREATE TABLE IF NOT EXISTS photos (
		id SERIAL PRIMARY KEY,
		title varchar(255) NOT NULL ,
		caption varchar(255) NOT NULL ,
		photo_url varchar(255) NOT NULL,
		user_id int NOT NULL,
		created_at timestamptz NOT NULL DEFAULT (now()),
		updated_at timestamptz  NULL DEFAULT NULL
	);`

	CommentTable := `
	CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
		user_id int NOT NULL,
		photo_id int NOT NULL,
		message varchar(255) NULL ,
		created_at timestamptz NOT NULL DEFAULT (now()),
		updated_at timestamptz  NULL DEFAULT NULL
	);`

	SocialMediaTable := `
	CREATE TABLE IF NOT EXISTS social_medias (
		id SERIAL PRIMARY KEY,
		name varchar(255) NOT NULL,
		sical_media_url varchar(255) NOT NULL,
		user_id int NOT NULL,
		created_at timestamptz NOT NULL DEFAULT (now()),
		updated_at timestamptz  NULL DEFAULT NULL
	);`

	createTableQueries := fmt.Sprintf("%s  %s %s %s", UserTable, PhotoTable, CommentTable, SocialMediaTable)
	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Fatal("error while creating movies table =>", err.Error())
	}
}

func InitializeDB() {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, dbPort, dbName)

	db, err = sqlx.Connect(dbDriver, dsn)

	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("error while tyring to ping the database connection", err.Error())
	}

	fmt.Println("successfully connected to my database")
	createRequiredTables()
}

func GetDB() *sqlx.DB {
	return db
}
