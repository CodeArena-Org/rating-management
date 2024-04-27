package db

import (
	"fmt"

	_ "github.com/lib/pq"
)

func UpdateRating(userId string, score int32) {
	// increment score
	_, err := DB.Exec("UPDATE users SET rating = rating + $1 WHERE userid = $2", score, userId)
	if err != nil {
		fmt.Println("Error updating rating: ", err)
	}
	// print updated data
	rows, err := DB.Query("SELECT * FROM users WHERE userid = $1", userId)
	if err != nil {
		fmt.Println("Error fetching updated rating: ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var userid string
		var rating int
		var username string
		err = rows.Scan(&userid, &rating, &username)
		if err != nil {
			fmt.Println("Error fetching updated rating: ", err)
		}
		fmt.Println("Updated rating: ", rating)
	}

}
