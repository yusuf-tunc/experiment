package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// create_server()
	struct_example()
}

func create_server() {
		router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
	})

	router.GET("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Login"})
	})

	router.Run(":8080")
}


type Entity interface  {
	update_name(name string)
}


type User struct {
	id string;
	name string;
	username string;
	password string;
	role string;
}

func (user *User) update_name(name string) {
	user.name = name
}

type Item struct {
	id string;
	name string;
	price float64;
}

func (item *Item) update_name(name string) {
	item.name = name
}

func changeNameOf(target Entity, name string) {
	target.update_name(name)
}

func struct_example() {
	user := &User{"user_id", "user_name", "username", "password", "admin"}
	user.update_name("NEW USER NAME")

	item := &Item{"item_id", "item_name", 10.0}
	item.update_name("NEW ITEM NAME")

	fmt.Println("BEFORE:",user.name, "-", item.name)

	changeNameOf(user, "NEW USER NAME 2")
	changeNameOf(item, "NEW ITEM NAME 2")

	fmt.Println("AFTER:", user.name, "-", item.name)
}

