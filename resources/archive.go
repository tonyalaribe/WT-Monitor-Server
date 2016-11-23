package resources

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tonyalaribe/monitor-server/config"
	"github.com/tonyalaribe/monitor-server/logger"
	"github.com/tonyalaribe/monitor-server/messages"
	"github.com/tonyalaribe/monitor-server/models"
)

type Archive struct {
	Base
}

func (p Archive) Get(c *gin.Context) {

	preDay, err := url.QueryUnescape(c.Query("day"))
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	day := strings.Replace(preDay, " ", "+", -1)
	log.Println(day)
	client := models.Client{Path: config.Get().DB.File}
	err = client.Open()
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}
	defer client.DB.Close()

	data := models.Archive{}

	date, err := time.Parse(time.RFC3339, day)
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	archiveURL, err := data.GetForDate(client.DB, date)
	if err != nil {
		logger.Error(c, err)
		c.JSON(http.StatusInternalServerError, messages.ErrInternalServer)
		return
	}

	c.File(archiveURL)
}
