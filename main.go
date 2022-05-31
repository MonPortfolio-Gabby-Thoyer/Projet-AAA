package main

import (
	"fmt"
	forum "forum/Server"
)

func main() {
	forum.HandleFunc()
	db := forum.InitDatabase("AAAforum.db")
	defer db.Close()

	// fmt.Println("-- Creation --")
	// fmt.Println("-- Créer --")

	// insertIntoUsers(db, "Mathieu", "m.m@gmail.com", "abcde")
	// insertIntoUsers(db, "Thomas", "t.t@gmail.com", "fghij")
	// insertIntoUsers(db, "Lucas", "l.l@gmail.com", "klmno")

	// fmt.Println("-- Sélection --")

	forum.SelectAllFromTable(db, "users")
	// user := selectUserById(db, 2)
	// fmt.Println(user)

	fmt.Println(forum.SelectUserNameWithPattern(db, "as"))
	// fmt.Println(test)
}
