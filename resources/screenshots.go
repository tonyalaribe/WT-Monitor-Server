package resources

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/logger"
	"github.com/tonyalaribe/monitor-server/messages"
	"github.com/tonyalaribe/monitor-server/models"
)

type Screenshots struct {
	Base
}

func (p Screenshots) Get(c *gin.Context) {
	id := c.Query("id")
	client := models.Client{Path: config.Get().DB.File}
	err := client.Open()
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
	defer client.DB.Close()

	data := models.Screenshot{
		UserID: id,
	}

	screenshots, err := data.GetTodaysScreenshots(client.DB, id)
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	c.JSON(http.StatusOK, screenshots)

}

func (p Screenshots) Post(c *gin.Context) {
	id := c.Request.FormValue("id")

	_, file, err := c.Request.FormFile("image")

	if err != nil {
		log.Println(err)
	}

	ff, err := file.Open()
	if err != nil {
		log.Println(err.Error())
	}
	buf, err := ioutil.ReadAll(ff)
	if err != nil {
		log.Println(err)
	}

	err = models.CreateScreenshot(id, buf)
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
	c.JSON(http.StatusOK, messages.Success)
}
