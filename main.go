package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"contacts/database"
	"contacts/interfaces"
	"contacts/models"
	"flag"

	"github.com/golang/glog"

	"github.com/gin-gonic/gin"
)

var (
	DB_CONNECTION string
	PORT          string
	IdbConnecter  interfaces.DbConnecter
	Database      database.Database
	IContact      interfaces.IContact
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
		flag.StringVar(&DB_CONNECTION, "dns", `jiten:jite123@tcp(127.0.0.1:3306)/contactsdb?charset=utf8mb4&parseTime=True&loc=Local`, "pass --dns=connection string to the database")
	}
	PORT = os.Getenv("PORT")

	if PORT == "" {
		flag.StringVar(&PORT, "port", ":50090", "--port=:50080(example port)")

	}

	flag.Parse()
	glog.Flush()

	IdbConnecter = &Database // New struct

	db, err := IdbConnecter.GetConnection(DB_CONNECTION)
	if err != nil {
		panic(err)
	}

	IContact := &database.ContactDB{DBClient: db}

	/*contact := &models.Contact{Name: "JIten", Email: "JitenP@outlook.com", Mobile: "9618558500"}
	err = contact.Validate()
	if err != nil {
		fmt.Println(err)
	} else {
		id, err := IContact.Create(contact)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Contact Id :", id)
		}
	}*/

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

	router.POST("/v1/person", func(c *gin.Context) {
		var buf []byte
		//	n, err := c.Request.Body.Read(buf)
		buf, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		contact := &models.Contact{}
		err = json.Unmarshal(buf, contact)
		if err != nil {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)

			c.Abort()
			return
		}
		err = contact.Validate()
		if err != nil {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			c.Abort()
			return
		}
		contact.Status = "inactive"
		contact.LastModified = time.Now().UTC().String()

		id, err := IContact.Create(contact)
		if err != nil {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)

			c.Abort()
			return
		}
		c.JSON(201, gin.H{
			"status":  "suceess",
			"message": id,
		})
		glog.Info("Success-->", id)

		c.Abort()
		return
	})

	router.Run(PORT)

}
