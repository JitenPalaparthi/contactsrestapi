package main

import (
	"fmt"
	"os"

	"contacts/database"
	"contacts/database/mysql"
	"contacts/database/postgres"
	"contacts/filedb"
	"contacts/handlers"
	"contacts/interfaces"
	"flag"

	"github.com/golang/glog"

	"github.com/gin-gonic/gin"
)

var (
	DB_CONNECTION    string
	PORT             string
	MODE             string
	IdbConnecter     interfaces.DbConnecter
	Database         mysql.Database
	PostgresDatabase postgres.Database
	IContact         interfaces.IContact
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	//flag.Parse()
}

// go mod tidy
func main() {

	DB_CONNECTION = os.Getenv("DB_CONNECTION")
	if DB_CONNECTION == "" {
		flag.StringVar(&DB_CONNECTION, "dsn", `jiten:jite123@tcp(127.0.0.1:3306)/contactsdb?charset=utf8mb4&parseTime=True&loc=Local`, "pass --dns=connection string to the database")
	}
	PORT = os.Getenv("PORT")

	if PORT == "" {
		flag.StringVar(&PORT, "port", ":50090", "--port=:50080(example port)")

	}

	MODE = os.Getenv("MODE")

	if MODE == "" {
		flag.StringVar(&MODE, "mode", "filedb", "--mode=mysql|filedb|postgres")

	}

	flag.Parse()
	glog.Flush()

	var IContact interfaces.IContact

	switch MODE {
	case "mysql":
		IdbConnecter = &Database // New struct

		db, err := IdbConnecter.GetConnection(DB_CONNECTION)
		if err != nil {
			panic(err)
		}
		IContact = &database.ContactDB{DBClient: db}
	case "postgres":
		//"host=localhost user=jiten password=jiten123 dbname=contactsdb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		// docker run -d -p 5432:5432--name mypostgres -e POSTGRES_PASSWORD=jiten123 -e POSTGRES_USER=jiten -e POSTGRES_DB=contactsdb postgress
		IdbConnecter = &PostgresDatabase // New struct

		db, err := IdbConnecter.GetConnection(DB_CONNECTION)
		if err != nil {
			panic(err)
		}

		IContact = &database.ContactDB{DBClient: db}

	case "filedb":
		IContact = &filedb.ContactFileDB{}
	default:
		panic("No db mode is provided")
	}

	contactHandler := &handlers.ContactHandler{IContact: IContact}

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

	router.POST("/v1/person", contactHandler.CreatePerson())
	router.DELETE("/v1/person/:id", contactHandler.DeletePerson())
	router.Run(PORT)

}
