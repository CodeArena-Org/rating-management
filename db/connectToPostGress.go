package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectToPostGress() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	POSTGRESS_URL := os.Getenv("POSTGRESS_URL")
	log.Println("Postgress URL: ", POSTGRESS_URL)
	db, err := sql.Open("postgres", POSTGRESS_URL)
	if err != nil {
		log.Fatalf("unable to connect to db: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("unable to ping db: %v", err)
	}
	log.Println("Connected to PostGress")
	setUpTables(db)
	DB = db
	return err
}

func setUpTables(db *sql.DB) {
	// create a table with userid string , rating number,  and username string
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (userid VARCHAR(50) PRIMARY KEY, rating INT, username VARCHAR(50))")
	if err != nil {
		log.Fatalf("unable to create table: %v", err)
	}

	// check if table is empty then add data
	count, err := db.Query("SELECT COUNT(*) FROM users")
	if err != nil {
		log.Fatalf("unable to check if table is empty: %v", err)
	}
	defer count.Close()
	var c int
	for count.Next() {
		err = count.Scan(&c)
		if err != nil {
			log.Fatalf("unable to check if table is empty: %v", err)
		}
	}
	if c > 0 {
		log.Println("Table already has data")
	} else {

		// insert two dummy users
		_, err = db.Exec("INSERT INTO users (userid, rating, username) VALUES ('1', 1000, 'user1')")
		if err != nil {
			log.Fatalf("unable to insert data: %v", err)
		}
		_, err = db.Exec("INSERT INTO users (userid, rating, username) VALUES ('2', 1000, 'user2')")
		if err != nil {
			log.Fatalf("unable to insert data: %v", err)
		}
		log.Println("Table created and two dummy users inserted")
	}

	// create a problem table with problem id and problem name and rating range(min, max)
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS problems (problemid VARCHAR(50) PRIMARY KEY, problemname VARCHAR(50), ratingrange INT[])")
	if err != nil {
		log.Fatalf("unable to create table: %v", err)
	}
	// add bunch of problems
	count, err = db.Query("SELECT COUNT(*) FROM problems")
	if err != nil {
		log.Fatalf("unable to check if table is empty: %v", err)
	}
	defer count.Close()
	for count.Next() {
		err = count.Scan(&c)
		if err != nil {
			log.Fatalf("unable to check if table is empty: %v", err)
		}
	}
	if c > 0 {
		log.Println("Problem table already has data")
	} else {
		insertProblems(db)
	}

	// create a table problemSolved with userid and problemid
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS problemsolved (userid VARCHAR(50), problemid VARCHAR(50), PRIMARY KEY(userid, problemid))")
	if err != nil {
		log.Fatalf("unable to create table: %v", err)
	}
	// clear the table
	_, err = db.Exec("DELETE FROM problemsolved")
	if err != nil {
		log.Fatalf("unable to clear table: %v", err)
	}
}

func insertProblems(db *sql.DB) {
	// Initialize random number generator with a seed
	rand.Seed(time.Now().UnixNano())

	// Loop to insert 100 problems
	for i := 0; i < 100; i++ {
		// Generate random rating range
		minRating := rand.Intn(2001)                       // Random number between 0 and 2000
		maxRating := rand.Intn(2001-minRating) + minRating // Random number between minRating and 2000
		ratingRange := fmt.Sprintf("{%d, %d}", minRating, maxRating)

		// Generate problem name (optional)
		problemName := fmt.Sprintf("problem%d", i+1)

		// Insert data into the database
		_, err := db.Exec("INSERT INTO problems (problemid, problemname, ratingrange) VALUES ($1, $2, $3)", fmt.Sprintf("%d", i+1), problemName, ratingRange)
		if err != nil {
			log.Fatalf("unable to insert data: %v", err)
		}
	}

	fmt.Println("Data inserted successfully!")
}
