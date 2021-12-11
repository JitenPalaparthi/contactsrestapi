package main

import (
	"fmt"
	"os"

	"flag"

	"github.com/gin-gonic/gin"

	"contacts/database"
	"contacts/interfaces"
)

var (
	DB_CONNECTION string
	PORT          string
	IdbConnecter  interfaces.DbConnecter
	Database      database.Database
)

// go mod tidy
func main() {

	DB_CONNECTION = os.Getenv("DB_CONNECTION")
	if DB_CONNECTION == "" {
		flag.StringVar(&DB_CONNECTION, "dns", `jiten:jite123@tcp(127.0.0.1:3306)/contactsdb?charset=utf8mb4&parseTime=True&loc=Local`, "pass --dns=connection string to the database")
	}
	PORT = os.Getenv("PORT")

	if PORT == "" {
		flag.StringVar(&PORT, "port", ":50090", "--port=:50080(example port)")

	}

	flag.Parse()

	IdbConnecter = &Database

	db, err := IdbConnecter.GetConnection(DB_CONNECTION)
	if err != nil {
		panic(err)
	}
	fmt.Println(db)

	router := gin.Default()

	//type HandlerFunc func(*Context)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "Hello World")

	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"health": "ok",
		})

	})

	router.Run(PORT)

}
