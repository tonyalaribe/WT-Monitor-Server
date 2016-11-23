package models

import (
	"testing"
	"time"

	"github.com/smartystreets/goconvey/convey"
	"github.com/tonyalaribe/monitor-server/config"
)

func TestArchive(t *testing.T) {
	convey.Convey("Should Print Screenshot", t, func() {
		c := Client{Path: config.Get().DB.TestFile}
		err := c.Open()
		convey.So(err, convey.ShouldBeNil)
		defer c.DB.Close()
		archive := Archive{
			UserID: "admin",
		}

		//now := time.Now().AddDate(0, 0, -1)
		now := time.Now().AddDate(0, 0, 0)
		err = archive.Create(c.DB, time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()))
		convey.So(err, convey.ShouldBeNil)

		err = DoArchiveForYesterday()
		convey.So(err, convey.ShouldBeNil)
	})
}
