package resources

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/logger"
	"github.com/tonyalaribe/monitor-server/messages"
	"github.com/tonyalaribe/monitor-server/models"
)

type Users struct {
	Base
}

func (p Users) Get(c *gin.Context) {
	client := models.Client{Path: config.Get().DB.File}
	err := client.Open()
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
	defer client.DB.Close()

	data := models.User{}

	users, err := data.GetAll(client.DB)
	log.Println(users)
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, users)

}

func (p Users) Post(c *gin.Context) {
	log.Print("hhh")
	id := c.Query("id")
	log.Print(id)
	client := models.Client{Path: config.Get().DB.File}
	err := client.Open()
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
	defer client.DB.Close()

	data := models.User{
		UserName: id,
	}

	err = data.Create(client.DB)

	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, messages.Success)

}
