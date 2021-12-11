package handlers

import (
	"contacts/interfaces"
	"contacts/models"
	"encoding/json"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/golang/glog"

	"github.com/gin-gonic/gin"
)

type ContactHandler struct {
	IContact interfaces.IContact
}

func (ch *ContactHandler) CreatePerson() func(c *gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IContact == nil {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ContactHandler or IContact is nil")

			c.Abort()
			return
		}
		var buf []byte
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

		id, err := ch.IContact.Create(contact)
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

	}
}

func (ch *ContactHandler) DeletePerson() func(c *gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IContact == nil {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ContactHandler or IContact is nil")

			c.Abort()
			return
		}

		id, ok := c.Params.Get("id")
		if !ok {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": "invalid id",
			})
			glog.Errorln("Could not fetch param id")

			c.Abort()
			return
		}
		_id, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)

			c.Abort()
			return
		}

		result, err := ch.IContact.Delete(_id)

		if err != nil {
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)

			c.Abort()
			return
		}
		var i = result.(int64)
		if i == 1 {
			c.JSON(200, gin.H{
				"status":  "suceess",
				"message": result,
			})
			glog.Info("Success-->", result)
		} else {
			c.JSON(204, gin.H{
				"status":  "no content",
				"message": result,
			})
			glog.Info("Success-->", result)
		}

		c.Abort()

	}
}
