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

type Timelog struct {
	Base
}

func (p Timelog) Get(c *gin.Context) {
	id := c.Query("id")
	log.Printf("the id is %s", id)

	client := models.Client{Path: config.Get().DB.File}
	err := client.Open()
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
	defer client.DB.Close()

	data := models.UserTimeLog{
		UserID: id,
	}
	log.Printf("%+v", data)
	log.Println("thats the data sttruct")
	timeline, err := data.GetTimeline(client.DB)

	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, timeline)

}

func (p Timelog) Post(c *gin.Context) {
	id := c.Query("id")
	client := models.Client{Path: config.Get().DB.File}
	err := client.Open()
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
	defer client.DB.Close()

	data := models.UserTimeLog{
		UserID: id,
	}
	err = data.LoginNow(client.DB)
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
	c.JSON(http.StatusOK, messages.Success)
}
